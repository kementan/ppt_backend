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
	router.PUT("role", handler.Update)
	router.DELETE("role", handler.Delete)
}

func (handler *RoleHandler) Create(ctx *gin.Context) {
	handler.UseCase.Create(ctx)
}

func (handler *RoleHandler) Read(ctx *gin.Context) {
	handler.UseCase.Read(ctx)
}

func (handler *RoleHandler) Update(ctx *gin.Context) {
	handler.UseCase.Update(ctx)
}

func (handler *RoleHandler) Delete(ctx *gin.Context) {
	handler.UseCase.Delete(ctx)
}
