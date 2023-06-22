package service

import (
	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	UseCase ServiceUseCase
}

func NewHandler(router *gin.Engine, usecase ServiceUseCase) {
	handler := &ServiceHandler{
		UseCase: usecase,
	}

	router.POST("service", handler.Create)
	router.GET("service", handler.Read)
	router.PUT("service", handler.Update)
	router.DELETE("service", handler.Delete)
}

func (handler *ServiceHandler) Create(ctx *gin.Context) {
	handler.UseCase.Create(ctx)
}

func (handler *ServiceHandler) Read(ctx *gin.Context) {
	handler.UseCase.Read(ctx)
}

func (handler *ServiceHandler) Update(ctx *gin.Context) {
	handler.UseCase.Update(ctx)
}

func (handler *ServiceHandler) Delete(ctx *gin.Context) {
	handler.UseCase.Delete(ctx)
}
