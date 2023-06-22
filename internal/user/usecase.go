package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/xsysproject/ppt_backend/helper"
)

type (
	UserUseCase interface {
		Login(ctx *gin.Context)
		Create(ctx *gin.Context)
		Dummy(ctx *gin.Context)
		Read(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	useCase struct {
		repo UserRepository
	}
)

func NewUseCase(repo UserRepository) UserUseCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) Login(ctx *gin.Context) {
	var req UserLoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.JERR(ctx, http.StatusInternalServerError, err)
		return
	}

	user, err := uc.repo.GetDataBy(ctx, "nik", req.NIK)
	if err != nil {
		if err == sql.ErrNoRows {
			helper.JERR(ctx, http.StatusNotFound, err)
			return
		}
		helper.JERR(ctx, http.StatusInternalServerError, err)
		return
	}

	err = helper.CheckPassword(req.Password, user.Password)
	if err != nil {
		helper.JERR(ctx, http.StatusUnauthorized, err)
		return
	}
}

func (uc *useCase) Create(ctx *gin.Context) {
	var req UserCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.JERR(ctx, http.StatusInternalServerError, err)
		return
	}

	exist, err := uc.repo.CheckESData(ctx, "email", req.Email)
	if exist {
		helper.JERR(ctx, http.StatusConflict, err)
		return
	}

	exist1, err := uc.repo.CheckESData(ctx, "nik", req.NIK)
	if exist1 {
		helper.JERR(ctx, http.StatusConflict, err)
		return
	}

	exist2, err := uc.repo.CheckESData(ctx, "nip", req.NIP)
	if exist2 {
		helper.JERR(ctx, http.StatusConflict, err)
		return
	}

	exist3, err := uc.repo.CheckESData(ctx, "phone", req.Phone)
	if exist3 {
		helper.JERR(ctx, http.StatusConflict, err)
		return
	}

	role_id, _ := helper.Decrypt(req.RoleID)
	email, _ := helper.Encrypt(req.Email)
	nik, _ := helper.Encrypt(req.NIK)
	nip, _ := helper.Encrypt(req.NIP)
	password, _ := helper.HashPassword(req.Password)
	phone, _ := helper.Encrypt(req.Phone)

	arg := UserCreate{
		RoleID:   role_id,
		Name:     req.Name,
		NIK:      nik,
		NIP:      nip,
		Email:    email,
		Password: password,
		Phone:    phone,
	}

	data, err := uc.repo.Create(ctx, arg)
	if err != nil {
		helper.JERR(ctx, http.StatusInternalServerError, err)
		return
	}

	helper.JOK(ctx, http.StatusOK, data)
}

func (uc *useCase) Dummy(ctx *gin.Context) {
	err := uc.repo.Dummy(ctx)
	if err != nil {
		helper.JERR(ctx, http.StatusInternalServerError, err)
		return
	}

	helper.JOK(ctx, http.StatusOK, nil)
}

func (uc *useCase) Read(ctx *gin.Context) {
	data, err := uc.repo.Read(ctx)
	if err != nil {
		helper.JERR(ctx, http.StatusInternalServerError, err)
		return
	}

	helper.JOK(ctx, http.StatusOK, data)
}

func (uc *useCase) Update(ctx *gin.Context) {
	var req UserUpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.JERR(ctx, http.StatusInternalServerError, err)
		return
	}

	arg := UserUpdate{
		Name: sql.NullString{
			String: req.Name,
			Valid:  true,
		},
	}

	data, err := uc.repo.Update(ctx, req.ID, arg)
	if err != nil {
		helper.JERR(ctx, http.StatusInternalServerError, err)
	}

	helper.JOK(ctx, http.StatusOK, data)
}

func (uc *useCase) Delete(ctx *gin.Context) {
	var req UserDeleteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.JERR(ctx, http.StatusInternalServerError, err)
		return
	}

	err := uc.repo.Delete(ctx, req.ID)
	if err != nil {
		helper.JERR(ctx, http.StatusInternalServerError, err)
		return
	}

	helper.JOK(ctx, http.StatusOK, "data successfully deleted")
}
