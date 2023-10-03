package user

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type UserHandler struct {
	Usecase UserUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase UserUsecase, rdb *redis.Client) {
	handler := &UserHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("18734asd22343/346sdcc8s9938499sd/init-create", handler.InitCreate)       //
	v1.GET("18734asd22343/823278sdacc829938s/remove-session", handler.RemoveSession) //
	v1.GET("18734asd22343/149bcb78926e31e3db/clean-user", handler.CleanUser)         //
	v1.GET("18734asd22343/18734926e31e7839sd/data-valid", handler.CheckData)         //

	v1.GET("user", util.AuthMiddleware(handler.rdb), handler.Read)
	v1.PUT("user", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("user/login", handler.Login)
	v1.POST("user/register", handler.Register)
	v1.POST("user/verify-recaptcha", handler.VerifyReCaptcha)
	v1.POST("user/logout", util.AuthMiddleware(handler.rdb), handler.Logout)
	v1.POST("user/refresh", util.AuthMiddleware(handler.rdb), handler.Refresh)
	v1.DELETE("user", util.AuthMiddleware(handler.rdb), handler.Delete)
}

func (handler *UserHandler) InitCreate(c *gin.Context) {
	handler.Usecase.InitCreate(c)
}

func (handler *UserHandler) RemoveSession(c *gin.Context) {
	handler.Usecase.RemoveSession(c)
}

func (handler *UserHandler) CleanUser(c *gin.Context) {
	handler.Usecase.CleanUser(c)
}

func (handler *UserHandler) CheckData(c *gin.Context) {
	handler.Usecase.CheckData(c)
}

func (handler *UserHandler) Login(c *gin.Context) {
	handler.Usecase.Login(c)
}

func (handler *UserHandler) VerifyReCaptcha(c *gin.Context) {
	handler.Usecase.VerifyReCaptcha(c)
}

func (handler *UserHandler) Register(c *gin.Context) {
	handler.Usecase.Register(c)
}

func (handler *UserHandler) Refresh(c *gin.Context) {
	handler.Usecase.Refresh(c)
}

func (handler *UserHandler) Logout(c *gin.Context) {
	handler.Usecase.Logout(c)
}

func (handler *UserHandler) Read(c *gin.Context) {
	handler.Usecase.Read(c)
}

func (handler *UserHandler) Update(c *gin.Context) {
	handler.Usecase.Update(c)
}

func (handler *UserHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
