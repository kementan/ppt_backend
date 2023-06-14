package service

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	var req ServiceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := ServiceCreate{
		Name:      req.Name,
		CreatedAt: time.Now(),
	}

	data, err := uc.repo.Create(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (uc *useCase) Read(ctx *gin.Context) {
	data, err := uc.repo.Read(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (uc *useCase) Update(ctx *gin.Context) {
	var req ServiceUpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, data)
}

func (uc *useCase) Delete(ctx *gin.Context) {
	var req ServiceDeleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := uc.repo.Delete(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	msg := "data successfully deleted"

	ctx.JSON(http.StatusOK, msg)
}
