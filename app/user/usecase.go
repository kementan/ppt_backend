package user

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gigaflex-co/ppt_backend/config"
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type (
	UserUsecase interface {
		InitCreate(c *gin.Context)
		RemoveSession(c *gin.Context)
		CleanUser(c *gin.Context)
		CheckData(c *gin.Context)
		Login(c *gin.Context)
		VerifyReCaptcha(c *gin.Context)
		Logout(c *gin.Context)
		Refresh(c *gin.Context)
		Register(c *gin.Context)
		Read(c *gin.Context)
		Update(c *gin.Context)
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

func (uc *usecase) CheckData(c *gin.Context) {
	ok, err := uc.repo.CheckData(c)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, ok)
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

	err = util.CheckPassword(req.Password, user.Password)
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
		"username": user.Username,
		"email":    user.Email,
		"token":    token,
	}

	util.JOK(c, http.StatusOK, data)
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

func (uc *usecase) Refresh(c *gin.Context) {
	token, err := util.Refresh(c)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, gin.H{"token": token})
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

	role_id, err := uc.repo.GetDefaultRole(c)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}
	name, _ := util.Encrypt(req.Name, "f")
	password, _ := util.HashPassword(req.Password)

	arg := UserCreate{
		RoleID:   role_id,
		Name:     name,
		Username: req.Username,
		Email:    req.Email,
		Password: password,
	}

	data, err := uc.repo.Create(c, arg)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, data)
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
	var req UserUpdateRequest

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	arg := UserUpdate{
		Name: sql.NullString{
			String: req.Name,
			Valid:  true,
		},
	}

	data, err := uc.repo.Update(c, req.ID, arg)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
	}

	util.JOK(c, http.StatusOK, data)
}

func (uc *usecase) Delete(c *gin.Context) {
	var req UserDeleteRequest

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	err = uc.repo.Delete(c, req.ID)
	if err != nil {
		util.JERR(c, http.StatusInternalServerError, err)
		return
	}

	util.JOK(c, http.StatusOK, "data successfully deleted")
}
