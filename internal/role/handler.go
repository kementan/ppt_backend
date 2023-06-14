package role

import (
	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	UseCase RoleUseCase
}

func NewHandler(router *gin.Engine, usecase RoleUseCase) {
	handler := &RoleHandler{
		UseCase: usecase,
	}

	router.POST("role", handler.Create)
	router.GET("role", handler.Read)
}

func (handler *RoleHandler) Create(ctx *gin.Context) {
	handler.UseCase.Create(ctx)
}

func (handler *RoleHandler) Read(ctx *gin.Context) {
	handler.UseCase.Read(ctx)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
