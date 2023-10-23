package user

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gigaflex-co/ppt_backend/config"
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
)

type (
	UserUsecase interface {
		InitCreate(c *gin.Context)
		RemoveSession(c *gin.Context)
		CleanUser(c *gin.Context)
		GetTable(c *gin.Context)
		GetByID(c *gin.Context)
		IsComplete(c *gin.Context)
		IsVerified(c *gin.Context)
		Login(c *gin.Context)
		VerifyReCaptcha(c *gin.Context)
		VerifyEmail(c *gin.Context)
		Logout(c *gin.Context)
		Refresh(c *gin.Context)
		DoCompletion(c *gin.Context)
		GetCompletion(c *gin.Context)
		Register(c *gin.Context)
		Read(c *gin.Context)
		Update(c *gin.Context)
		IsRegistered(c *gin.Context)
		IsNullPassword(c *gin.Context)
		Delete(c *gin.Context)
	}

	usecase struct {
		repo UserRepository
		rdb  *redis.Client
		cfg  config.Config
	}
)

func NewUsecase(repo UserRepository, rdb *redis.Client, config config.Config) UserUsecase {
	return &usecase{
		repo: repo,
		rdb:  rdb,
		cfg:  config,
	}
}

func (uc *usecase) InitCreate(c *gin.Context) {
	ok, err := uc.repo.InitCreate(c)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, ok)
}

func (uc *usecase) RemoveSession(c *gin.Context) {
	if err := uc.rdb.FlushDB(c.Request.Context()).Err(); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "successfully remove all sessions")
}

func (uc *usecase) CleanUser(c *gin.Context) {
	ok, err := uc.repo.CleanUser(c)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, ok)
}

func (uc *usecase) GetTable(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	sortBy := c.DefaultQuery("sort_by", "id")
	sortOrder := c.DefaultQuery("sort_order", "desc")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	arg := util.DataFilter{
		Search:    search,
		SortBy:    sortBy,
		SortOrder: sortOrder,
		Page:      pageInt,
		PageSize:  pageSizeInt,
	}

	totalRecords, err := uc.repo.CountRecords(c, arg)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSizeInt)))

	data, err := uc.repo.GetTable(c, arg)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	pagination := util.PaginationResponse{
		CurrentPage:  pageInt,
		PageSize:     pageSizeInt,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	}

	response := DataWithPagination{
		Row:        data,
		Pagination: pagination,
	}

	util.JOK(c, http.StatusOK, response)
}

func (uc *usecase) GetByID(c *gin.Context) {
	type ID struct {
		ID string `json:"id" binding:"required,min=1"`
	}

	var req ID

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	id, _ := util.Decrypt(req.ID, "f")

	data, err := uc.repo.GetDataBy(c, "u.id", id)
	if err != nil {
		if err.Error() == "record not found" {
			util.JERR(c, http.StatusNotFound, errors.New("record not found"))
		} else {
			util.JERR(c, http.StatusInternalServerError, err)
		}
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) Login(c *gin.Context) {
	var req UserLoginRequest

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	user, err := uc.repo.PassByUEmail(c, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.JERR(c, http.StatusNotFound, errors.New("invalid username/password"))
			return
		}
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if req.IsGoogle {
		err = util.CheckPassword(req.Password, user.GoogleID.String)
	} else {
		err = util.CheckPassword(req.Password, user.Password)
	}

	if err != nil {
		util.JERR(c, http.StatusUnauthorized, errors.New("invalid username/password"))
		return
	}

	token, err := util.CreateToken(c, uc.rdb, user.Email)
	if err != nil {
		if err.Error() == "isloggedin" {
			util.JERR(c, http.StatusBadRequest, errors.New("user currently active"))
			return
		}
		util.JERR(c, http.StatusInternalServerError, errors.New("failed to generate new token"))
		return
	}

	data := map[string]string{
		"name":     "",
		"username": user.Username,
		"email":    user.Email,
		"token":    token,
		"prov_id":  "",
		"reg_id":   "",
		"sub_id":   "",
		"urb_id":   "",
		"address":  "",
		"nik":      "",
	}

	if user.Name.Valid {
		data["name"] = user.Name.String
	}

	if user.ProvinceID.Valid {
		data["prov_id"] = user.ProvinceID.String
	}

	if user.RegencyID.Valid {
		data["reg_id"] = user.RegencyID.String
	}

	if user.SubdistrictID.Valid {
		data["sub_id"] = user.SubdistrictID.String
	}

	if user.UrbanvillageID.Valid {
		data["urb_id"] = user.UrbanvillageID.String
	}

	if user.Address.Valid {
		data["address"] = user.Address.String
	}

	if user.NIK.Valid {
		data["nik"] = user.NIK.String
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) IsComplete(c *gin.Context) {
	email := c.Param("email")
	userResponse, err := uc.repo.IsComplete(c, email)
	if err != nil && err != sql.ErrNoRows {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	} else if err == sql.ErrNoRows {
		util.JOK(c, http.StatusOK, false)
		return
	}

	if userResponse.IsComplete {
		res := map[string]any{
			"status": true,
			"name":   userResponse.Name,
		}
		util.JOK(c, http.StatusOK, res)
		return
	} else {
		res := map[string]any{
			"status": false,
			"name":   userResponse.Name,
		}
		util.JOK(c, http.StatusOK, res)
		return
	}
}

func (uc *usecase) IsVerified(c *gin.Context) {
	email := c.Param("email")
	userResponse, err := uc.repo.IsVerified(c, email)
	if err != nil && err != sql.ErrNoRows {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	} else if err == sql.ErrNoRows {
		util.JOK(c, http.StatusOK, false)
		return
	}

	if userResponse.IsVerified {
		res := map[string]any{
			"status": true,
			"name":   userResponse.Name,
		}
		util.JOK(c, http.StatusOK, res)
		return
	} else {
		res := map[string]any{
			"status": false,
			"name":   userResponse.Name,
		}
		util.JOK(c, http.StatusOK, res)
		return
	}
}

func (uc *usecase) VerifyReCaptcha(c *gin.Context) {
	var req UserRecaptchaRequest

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err = util.VerifyReCaptcha(c, req.Response); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	util.JOK(c, http.StatusOK, "recaptcha successfully validated")
}

func (uc *usecase) VerifyEmail(c *gin.Context) {
	var req EmailVerificationCode

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	email, err := util.Decrypt(req.VerificationCode, "f")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	data, err := uc.repo.GetDataBy(c, "email", email)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if data.IsVerified {
		util.JOK(c, http.StatusOK, data.IsVerified)
		return
	}

	verified, err := uc.repo.DoVerify(c, email)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
	}

	if !verified.IsVerified {
		util.JOK(c, http.StatusOK, true)
		return
	}

	util.JOK(c, http.StatusOK, false)
}

func (uc *usecase) Refresh(c *gin.Context) {
	token, err := util.Refresh(c)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, gin.H{"token": token})
}

func DukcapilValidation(payload DukcapilPayload) ([]DukcapilContent, error) {
	originalDate, err := time.Parse("2006-01-02", payload.TanggalLahir)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, err
	}

	formattedDate := originalDate.Format("02-01-2006")

	pd := []byte(`{
		"NIK": "` + payload.NIK + `",
		"NAMA_LGKP": "` + payload.NamaLengkap + `",
		"JENIS_KLMIN": "` + payload.JenisKelamin + `",
		"TMPT_LHR": "` + payload.TempatLahir + `",
		"TGL_LHR": "` + formattedDate + `",
		"TRESHOLD": 100,
		"user_id": "26082022160454BADAN_PENYULUHAN_SDM8370",
		"password": "TH854Y",
		"ip_user": "10.160.84.10"
	}`)

	req, err := http.NewRequest("POST", "http://10.1.241.250/api/nik_verifby_elemen.php", bytes.NewBuffer(pd))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response DukcapilResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	content := response.Content

	return content, nil
}

func (uc *usecase) DoCompletion(c *gin.Context) {
	var r CompletionRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	var payload DukcapilPayload
	var gender string

	if r.BioData.Gender == "l" {
		gender = "Laki-laki"
	} else {
		gender = "Perempuan"
	}

	if r.Section == "bio" {
		payload = DukcapilPayload{
			NIK:          r.BioData.NIK,
			NamaLengkap:  r.BioData.Name,
			JenisKelamin: gender,
			TempatLahir:  r.BioData.POB,
			TanggalLahir: r.BioData.DOB,
			TRESHOLD:     100,
			UserID:       "26082022160454BADAN_PENYULUHAN_SDM8370",
			Password:     "TH854Y",
			IPUser:       "10.160.84.10",
		}

		content, err := DukcapilValidation(payload)
		if err != nil {
			util.JERR(c, http.StatusInternalServerError, err)
			return
		}

		namaLengkapSesuai := true
		tempatLahirSesuai := true
		tanggalLahirSesuai := true
		nikSesuai := true
		jenisKelaminSesuai := true

		for _, item := range content {
			if item.NamaLengkap != "Sesuai (100)" {
				namaLengkapSesuai = false
			}
			if item.TempatLahir != "Sesuai (100)" {
				tempatLahirSesuai = false
			}
			if item.TanggalLahir != "Sesuai" {
				tanggalLahirSesuai = false
			}
			if item.NIK != "Sesuai" {
				nikSesuai = false
			}
			if item.JenisKelamin != "Sesuai" {
				jenisKelaminSesuai = false
			}
		}

		if namaLengkapSesuai && tempatLahirSesuai && tanggalLahirSesuai && nikSesuai && jenisKelaminSesuai {
			name, _ := util.Encrypt(payload.NamaLengkap, "f")
			pob, _ := util.Encrypt(payload.TempatLahir, "f")
			dob, _ := util.Encrypt(payload.TanggalLahir, "f")
			nik, _ := util.Encrypt(r.BioData.NIK, "f")
			role_id, _ := util.Decrypt(r.BioData.RoleID, "f")
			phone, _ := util.Encrypt(r.BioData.Phone, "f")

			dataToUpdate := UserUpdate{
				RoleID: sql.NullString{
					String: role_id,
					Valid:  true,
				},
				Name: sql.NullString{
					String: name,
					Valid:  true,
				},
				POB: sql.NullString{
					String: pob,
					Valid:  true,
				},
				DOB: sql.NullString{
					String: dob,
					Valid:  true,
				},
				NIK: sql.NullString{
					String: nik,
					Valid:  true,
				},
				Gender: sql.NullString{
					String: r.BioData.Gender,
					Valid:  true,
				},
				Phone: sql.NullString{
					String: phone,
					Valid:  true,
				},
			}

			_, err := uc.repo.Update(c, "email", r.Email, dataToUpdate, r.Section)
			if err != nil {
				util.JERR(c, http.StatusInternalServerError, err)
				return
			}
		}
		util.JOK(c, http.StatusOK, content)
		return
	} else if r.Section == "address" {
		dataToUpdate := UserUpdate{
			ProvinceID: sql.NullString{
				String: r.AddressData.ProvID,
				Valid:  true,
			},
			RegencyID: sql.NullString{
				String: r.AddressData.RegID,
				Valid:  true,
			},
			SubdistrictID: sql.NullString{
				String: r.AddressData.SubID,
				Valid:  true,
			},
			UrbanvillageID: sql.NullString{
				String: r.AddressData.UrbID,
				Valid:  true,
			},
			Address: sql.NullString{
				String: r.AddressData.Address,
				Valid:  true,
			},
			Latitude: sql.NullFloat64{
				Float64: r.AddressData.Latitude,
				Valid:   true,
			},
			Longitude: sql.NullFloat64{
				Float64: r.AddressData.Longitude,
				Valid:   true,
			},
		}

		data, err := uc.repo.Update(c, "email", r.Email, dataToUpdate, r.Section)
		if err != nil {
			util.JERR(c, http.StatusInternalServerError, err)
			return
		}
		util.JOK(c, http.StatusOK, data)
		return
	}
}

func (uc *usecase) GetCompletion(c *gin.Context) {
	type User struct {
		Email string `json:"email"`
	}

	var r User

	if err := c.ShouldBindJSON(&r); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	data, err := uc.repo.GetDataBy(c, "email", r.Email)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	email, err := util.RevokeToken(c, uc.rdb)
	if err != nil {
		util.JERR(c, http.StatusUnauthorized, err)
	}

	util.JOK(c, http.StatusOK, gin.H{
		"user":    email,
		"message": "Logout successful",
	})
}

func (uc *usecase) Register(c *gin.Context) {
	var req UserRegisterRequest

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	check := map[string]string{
		"username": req.Username,
		"email":    req.Email,
	}

	for key, val := range check {
		_, err := uc.repo.GetDataBy(c, key, val)
		if err != nil && err != sql.ErrNoRows {
			util.JERR(c, http.StatusConflict, errors.New(key+" sudah digunakan"))
			return
		} else if err == nil {
			util.JERR(c, http.StatusConflict, errors.New(key+" sudah digunakan"))
			return
		}
	}

	var role_id int
	if req.RoleID == "" {
		role_id, err = uc.repo.GetDefaultRole(c)
		if err != nil {
			util.JERR(c, http.StatusInternalServerError, err)
			return
		}
	} else {
		dec_role_id, _ := util.Decrypt(req.RoleID, "f")
		role_id, err = strconv.Atoi(dec_role_id)
		if err != nil {
			util.JERR(c, http.StatusInternalServerError, err)
			return
		}
	}

	name, _ := util.Encrypt(req.Name, "f")
	var password, google_id string

	if req.IsGoogle {
		google_id, _ = util.HashPassword(req.Password)
	} else {
		password, _ = util.HashPassword(req.Password)
	}

	arg := UserCreate{
		RoleID:   role_id,
		Name:     name,
		Username: req.Username,
		Email:    req.Email,
		NIK:      req.NIK,
		Password: password,
		GoogleID: google_id,
	}

	data, err := uc.repo.Create(c, arg)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	go sendActivationEmail(uc, c, uc.cfg.AllowOrigin, req.Email)

	util.JOK(c, http.StatusOK, data)
}

func sendActivationEmail(uc *usecase, c *gin.Context, origin, toEmail string) {
	m := gomail.NewMessage()

	data, err := uc.repo.GetConfig(c)
	if err != nil {
		log.Printf("failed to retrieve config: %v", err)
	}

	smtpServer := data.SMTPServer
	smtpPort := data.SMTPPort
	smtpEmail := data.SMTPEmail
	smtpPassword, _ := util.Decrypt(data.SMTPPassword, "f")

	m.SetHeader("From", smtpEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Verifikasi Email")

	code, err := util.Encrypt(toEmail, "f")
	if err != nil {
		log.Printf("to encrypt email: %v", err)
	}
	activationLink := origin + "/auth/email-verification/" + code
	emailBody := fmt.Sprintf("Silahkan klik link berikut untuk mengaktifkan email akun anda :\n<a href=\"%s\">Aktifkan Email</a>", activationLink)
	m.SetBody("text/html", emailBody)

	d := gomail.NewDialer(smtpServer, smtpPort, smtpEmail, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email: %v", err)
	}
}

func (uc *usecase) Read(c *gin.Context) {
	data, err := uc.repo.Read(c)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) Update(c *gin.Context) {
	var r UserUpdateRequest
	var password string

	if err := c.ShouldBindJSON(&r); err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	role_id, _ := util.Decrypt(r.RoleID, "f")
	name, _ := util.Encrypt(r.Name, "f")
	nik, _ := util.Encrypt(r.NIK, "f")

	if r.Password != "" {
		password, _ = util.HashPassword(r.Password)
	} else {
		if r.CurrentPassword != "" {
			password = r.CurrentPassword
		} else {
			password = ""
		}
	}

	dataToUpdate := UserUpdate{
		RoleID: sql.NullString{
			String: role_id,
			Valid:  true,
		},
		Name: sql.NullString{
			String: name,
			Valid:  true,
		},
		NIK: sql.NullString{
			String: nik,
			Valid:  true,
		},
		Username: sql.NullString{
			String: r.Username,
			Valid:  true,
		},
		Email: sql.NullString{
			String: r.Email,
			Valid:  true,
		},
		Password: sql.NullString{
			String: password,
			Valid:  true,
		},
		IsActive:   r.IsActive,
		IsComplete: r.IsComplete,
		IsVerified: r.IsVerified,
	}

	id, _ := util.Decrypt(r.ID, "f")

	data, err := uc.repo.Update(c, "id", id, dataToUpdate, "bio")
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) IsRegistered(c *gin.Context) {
	var req IsRegistered

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	data, err := uc.repo.GetDataBy(c, "email", req.Email)
	if err != nil && err != sql.ErrNoRows {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if err == sql.ErrNoRows {
		util.JOK(c, http.StatusOK, false)
		return
	}

	if data.Email == "" {
		util.JOK(c, http.StatusOK, false)
		return
	}

	util.JOK(c, http.StatusOK, true)
}

func (uc *usecase) IsNullPassword(c *gin.Context) {
	var req IsRegistered

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	data, err := uc.repo.IsNullPassword(c, "email", req.Email)
	if err != nil && err != sql.ErrNoRows {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	if err == sql.ErrNoRows {
		util.JOK(c, http.StatusOK, false)
		return
	}

	if data.Email == "" {
		util.JOK(c, http.StatusOK, false)
		return
	}

	util.JOK(c, http.StatusOK, true)
}

func (uc *usecase) Delete(c *gin.Context) {
	var req util.TableIDs

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	successIDs, failedIDs, err := uc.repo.Delete(c, req.IDs)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	response := struct {
		SuccessIDs []string `json:"success_ids"`
		FailedIDs  []string `json:"failed_ids"`
	}{
		SuccessIDs: successIDs,
		FailedIDs:  failedIDs,
	}

	util.JOK(c, http.StatusOK, response)
}
