package role

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	RoleUseCase interface {
		Create(ctx *gin.Context)
		Read(ctx *gin.Context)
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
