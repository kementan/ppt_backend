package role

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	RoleUseCase interface {
		Create(ctx *gin.Context)
		Read(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	useCase struct {
		repo RoleRepository
	}
)

func NewUseCase(repo RoleRepository) RoleUseCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) Create(ctx *gin.Context) {
	var req RoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := RoleCreate{
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
	var req RoleUpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := RoleUpdate{
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
	var req RoleDeleteRequest
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
