package user

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UseCase UserUseCase
}

func NewHandler(router *gin.Engine, usecase UserUseCase) {
	handler := &UserHandler{
		UseCase: usecase,
	}

	router.POST("login", handler.Login)
	router.POST("register", handler.Register)

	router.GET("dummy_user", handler.Dummy)
	router.GET("user", handler.Read)
	router.PUT("user", handler.Update)
	router.DELETE("user", handler.Delete)
}

func (handler *UserHandler) Login(ctx *gin.Context) {
	handler.UseCase.Login(ctx)
}

func (handler *UserHandler) Register(ctx *gin.Context) {
	handler.UseCase.Create(ctx)
}

func (handler *UserHandler) Dummy(ctx *gin.Context) {
	handler.UseCase.Dummy(ctx)
}

func (handler *UserHandler) Read(ctx *gin.Context) {
	handler.UseCase.Read(ctx)
}

func (handler *UserHandler) Update(ctx *gin.Context) {
	handler.UseCase.Update(ctx)
}

func (handler *UserHandler) Delete(ctx *gin.Context) {
	handler.UseCase.Delete(ctx)
}
