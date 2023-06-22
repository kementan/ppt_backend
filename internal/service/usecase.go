package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/xsysproject/ppt_backend/helper"
)

type (
	ServiceUseCase interface {
		Create(ctx *gin.Context)
		Read(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	useCase struct {
		repo ServiceRepository
	}
)

func NewUseCase(repo ServiceRepository) ServiceUseCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) Create(ctx *gin.Context) {
	var req ServiceCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.JERR(ctx, http.StatusInternalServerError, err)
		return
	}

	arg := ServiceCreate{
		Name: req.Name,
	}

	data, err := uc.repo.Create(ctx, arg)
	if err != nil {
		helper.JERR(ctx, http.StatusInternalServerError, err)
		return
	}

	helper.JOK(ctx, http.StatusOK, data)
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
	var req ServiceUpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.JERR(ctx, http.StatusInternalServerError, err)
		return
	}

	arg := ServiceUpdate{
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
	var req ServiceDeleteRequest

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
