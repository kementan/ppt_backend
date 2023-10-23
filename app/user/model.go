package user

import (
	"database/sql"
	"time"

	"github.com/gigaflex-co/ppt_backend/util"
)

var table = "ppt_users"
var subtable = "ppt_roles"

type (
	DataWithPagination struct {
		Row        []UserResponse          `json:"row"`
		Pagination util.PaginationResponse `json:"pagination"`
	}

	ConfigData struct {
		Name  string
		Value string
	}

	DukcapilPayload struct {
		NIK          string `json:"nik_val"`
		NamaLengkap  string `json:"NAMA_LGKP"`
		JenisKelamin string `json:"JENIS_KLMIN"`
		TempatLahir  string `json:"TMPT_LHR"`
		TanggalLahir string `json:"TGL_LHR"`
		TRESHOLD     int    `json:"TRESHOLD"`
		UserID       string `json:"user_id"`
		Password     string `json:"password"`
		IPUser       string `json:"ip_user"`
	}

	DukcapilDataRequest struct {
		NIK          string `json:"nik" binding:"required,min=16,max=16"`
		NamaLengkap  string `json:"nama_lengkap" binding:"required"`
		JenisKelamin string `json:"jenis_kelamin"  binding:"required"`
		TempatLahir  string `json:"tempat_lahir" binding:"required"`
	}

	DukcapilContent struct {
		NOKK           string `json:"NO_KK"`
		NamaLengkap    string `json:"NAMA_LGKP"`
		TempatLahir    string `json:"TMPT_LHR"`
		TanggalLahir   string `json:"TGL_LHR"`
		PropName       string `json:"PROP_NAME"`
		KabName        string `json:"KAB_NAME"`
		KecName        string `json:"KEC_NAME"`
		KelName        string `json:"KEL_NAME"`
		NoRT           string `json:"NO_RT"`
		NoRW           string `json:"NO_RW"`
		Alamat         string `json:"ALAMAT"`
		NamaIbu        string `json:"NAMA_LGKP_IBU"`
		StatusKawin    string `json:"STATUS_KAWIN"`
		JenisPekerjaan string `json:"JENIS_PKRJN"`
		JenisKelamin   string `json:"JENIS_KLMIN"`
		NoProp         string `json:"NO_PROP"`
		NoKab          string `json:"NO_KAB"`
		NoKec          string `json:"NO_KEC"`
		NoKel          string `json:"NO_KEL"`
		NIK            string `json:"NIK"`
	}

	DukcapilResponse struct {
		Content          []DukcapilContent `json:"content"`
		LastPage         bool              `json:"lastPage"`
		NumberOfElements int               `json:"numberOfElements"`
		Sort             string            `json:"sort"`
		TotalElements    int               `json:"totalElements"`
		FirstPage        bool              `json:"firstPage"`
		Number           int               `json:"number"`
		Size             int               `json:"size"`
	}

	SMTPData struct {
		SMTPServer   string
		SMTPPort     int
		SMTPEmail    string
		SMTPPassword string
	}

	UserLoginRequest struct {
		Username string `json:"username" binding:"required,min=8,max=150"`
		Password string `json:"password" binding:"required,min=8,max=150"`
		IsGoogle bool   `json:"is_google"`
	}

	UserLoginResponse struct {
		NIK            sql.NullString
		Name           sql.NullString
		Username       string
		Email          string
		Address        sql.NullString
		ProvinceID     sql.NullString
		RegencyID      sql.NullString
		SubdistrictID  sql.NullString
		UrbanvillageID sql.NullString
		Password       string
		GoogleID       sql.NullString
	}

	UserRegisterRequest struct {
		RoleID   string `json:"role_id"`
		NIK      string `json:"nik"`
		Name     string `json:"name" binding:"required,min=3,max=100"`
		Username string `json:"username" binding:"required,min=8,max=150"`
		Email    string `json:"email" binding:"required,min=8,max=150"`
		Password string `json:"password"`
		IsGoogle bool   `json:"is_google"`
	}

	CompletionRequest struct {
		BioData     UserBioRequest     `json:"bio_data"`
		AddressData UserAddressRequest `json:"address_data"`
		Section     string             `json:"section" binding:"required"`
		Email       string             `json:"email" binding:"required"`
	}

	UserBioRequest struct {
		RoleID      string   `json:"role_id_val" binding:"required"`
		SubSectorID []string `json:"sub_sector_id" binding:"required"`
		Name        string   `json:"name_val" binding:"required,min=3,max=100"`
		NIK         string   `json:"nik_val" binding:"required,min=16,max=16"`
		POB         string   `json:"pob_val" binding:"required"`
		DOB         string   `json:"dob_val" binding:"required"`
		Phone       string   `json:"phone_val" binding:"required"`
		Gender      string   `json:"gender_val" binding:"required"`
	}

	UserAddressRequest struct {
		ProvID    string  `json:"prov_id_val"`
		RegID     string  `json:"reg_id_val"`
		SubID     string  `json:"sub_id_val"`
		UrbID     string  `json:"urb_id_val"`
		Address   string  `json:"address_val"`
		Latitude  float64 `json:"latitude_val"`
		Longitude float64 `json:"longitude_val"`
	}

	UserLandRequest struct {
		ProvID    string `json:"prov_id_val"`
		RegID     string `json:"reg_id_val"`
		SubID     string `json:"sub_id_val"`
		UrbID     string `json:"urb_id_val"`
		Name      string `json:"name_val"`
		Address   string `json:"address_val"`
		Latitude  string `json:"latitude_val"`
		Longitude string `json:"longitude_val"`
	}

	UserRecaptchaRequest struct {
		Response string `json:"gresponse"`
	}

	UserCreateRequest struct {
		RoleID         string `json:"role_id" binding:"required,min=1,max=200"`
		ProvinceID     string `json:"province_id" binding:"required,min=1,max=200"`
		RegencyID      string `json:"regency_id" binding:"required,min=1,max=200"`
		SubdistrictID  string `json:"subdistrict_id" binding:"required,min=1,max=200"`
		UrbanvillageID string `json:"urbanvillage_id" binding:"required,min=1,max=200"`
		Name           string `json:"name" binding:"required,min=3,max=100"`
		Username       string `json:"username" binding:"required,min=8,max=150"`
		Email          string `json:"email" binding:"required,min=8,max=150"`
		Phone          string `json:"phone" binding:"required,min=9,max=15"`
		Password       string `json:"password" binding:"required,min=8,max=150"`
		NIK            string `json:"nik" binding:"required,min=16,max=16"`
		NIP            string `json:"nip"`
		DOB            string `json:"dob" binding:"required,min=8,max=150"`
		POB            string `json:"pob" binding:"required,min=8,max=150"`
		Gender         string `json:"gender" binding:"required"`
		ImgUser        string `json:"img_user"`
		ImgID          string `json:"img_id"`
		Address        string `json:"address" binding:"required,min=8,max=150"`
	}

	UserUpdateRequest struct {
		ID              string `json:"id" binding:"required,min=1"`
		RoleID          string `json:"role_id" binding:"required,min=1,max=200"`
		Name            string `json:"name" binding:"required,min=3,max=100"`
		Username        string `json:"username" binding:"required,min=8,max=150"`
		Email           string `json:"email" binding:"required,min=8,max=150"`
		NIK             string `json:"nik" binding:"required,min=16,max=16"`
		Password        string `json:"password"`
		IsActive        *bool  `json:"is_active"`
		IsComplete      *bool  `json:"is_complete"`
		IsVerified      *bool  `json:"is_verified"`
		CurrentPassword string `json:"current_password"`
	}

	UserDeleteRequest struct {
		ID string `json:"id" binding:"required,min=1"`
	}

	UserCreate struct {
		RoleID         int       `json:"role_id"`
		ProvinceID     string    `json:"province_id"`
		RegencyID      string    `json:"regency_id"`
		SubdistrictID  string    `json:"subdistrict_id"`
		UrbanvillageID string    `json:"urbanvillage_id"`
		Name           string    `json:"name"`
		Username       string    `json:"username"`
		Email          string    `json:"email"`
		Phone          string    `json:"phone"`
		Password       string    `json:"password"`
		GoogleID       string    `json:"google_id"`
		NIK            string    `json:"nik"`
		NIP            string    `json:"nip"`
		DOB            string    `json:"dob"`
		POB            string    `json:"pob"`
		ImgUser        string    `json:"img_user"`
		ImgID          string    `json:"img_id"`
		Address        string    `json:"address"`
		IsActive       bool      `json:"is_active"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	UserUpdate struct {
		RoleID         sql.NullString  `json:"role_id"`
		ProvinceID     sql.NullString  `json:"province_id"`
		RegencyID      sql.NullString  `json:"regency_id"`
		SubdistrictID  sql.NullString  `json:"subdistrict_id"`
		UrbanvillageID sql.NullString  `json:"urbanvillage_id"`
		Name           sql.NullString  `json:"name"`
		Username       sql.NullString  `json:"username"`
		Email          sql.NullString  `json:"email"`
		Phone          sql.NullString  `json:"phone"`
		Password       sql.NullString  `json:"password"`
		NIK            sql.NullString  `json:"nik"`
		NIP            sql.NullString  `json:"nip"`
		DOB            sql.NullString  `json:"dob"`
		POB            sql.NullString  `json:"pob"`
		Gender         sql.NullString  `json:"gender"`
		ImgUser        sql.NullString  `json:"img_user"`
		ImgID          sql.NullString  `json:"img_id"`
		Address        sql.NullString  `json:"address"`
		IsActive       *bool           `json:"is_active"`
		IsComplete     *bool           `json:"is_complete"`
		IsVerified     *bool           `json:"is_verified"`
		Latitude       sql.NullFloat64 `json:"latitude"`
		Longitude      sql.NullFloat64 `json:"longitude"`
		CreatedAt      sql.NullString  `json:"created_at"`
		UpdatedAt      sql.NullString  `json:"updated_at"`
	}

	UserResponse struct {
		HashedID       string    `json:"id"`
		RoleID         string    `json:"role_id"`
		RoleName       string    `json:"role_name"`
		ProvinceID     string    `json:"province_id"`
		RegencyID      string    `json:"regency_id"`
		SubdistrictID  string    `json:"subdistrict_id"`
		UrbanvillageID string    `json:"urbanvillage_id"`
		Name           string    `json:"name"`
		Username       string    `json:"username"`
		Password       string    `json:"password"`
		Email          string    `json:"email"`
		Phone          string    `json:"phone"`
		NIK            string    `json:"nik"`
		NIP            string    `json:"nip"`
		DOB            string    `json:"dob"`
		POB            string    `json:"pob"`
		ImgUser        string    `json:"img_user"`
		ImgID          string    `json:"img_id"`
		Address        string    `json:"address"`
		GoogleID       string    `json:"google_id"`
		IsActive       bool      `json:"is_active"`
		IsComplete     bool      `json:"is_complete"`
		IsVerified     bool      `json:"is_verified"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	SearchResponse struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
		} `json:"hits"`
	}

	UserIndex struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	EmailVerificationCode struct {
		VerificationCode string `json:"verification_code" binding:"required,min=1,max=255"`
	}

	IsRegistered struct {
		Email string `json:"email" binding:"required,min=8,max=150"`
	}
)
