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

	v1.GET("18734asd22343/346sdcc8s9938499sd/init-create", handler.InitCreate)       // create init 1 admin account
	v1.GET("18734asd22343/823278sdacc829938s/remove-session", handler.RemoveSession) // remove all session in redis
	v1.GET("18734asd22343/149bcb78926e31e3db/clean-user", handler.CleanUser)         // delete trash data
	v1.POST("user/is-registered", handler.IsRegistered)
	v1.POST("user/login", handler.Login)
	v1.POST("user/register", handler.Register)
	v1.POST("user/verify-recaptcha", handler.VerifyReCaptcha)
	v1.POST("user/verify-email", handler.VerifyEmail)

	v1.GET("8asd87asd98/7asd8a7sd68as7", util.AuthMiddleware(handler.rdb), handler.IsVerified) // get data by session
	v1.GET("user-is-complete/:email", util.AuthMiddleware(handler.rdb), handler.IsComplete)
	v1.GET("user-is-verified/:email", util.AuthMiddleware(handler.rdb), handler.IsVerified)
	v1.GET("user-list", util.AuthMiddleware(handler.rdb), handler.Read)
	v1.GET("user-table", util.AuthMiddleware(handler.rdb), handler.GetTable)
	v1.PUT("user-update", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("user-id", util.AuthMiddleware(handler.rdb), handler.GetByID)
	v1.POST("user/12389dsa-9982783", util.AuthMiddleware(handler.rdb), handler.IsNullPassword)
	v1.POST("user/logout", util.AuthMiddleware(handler.rdb), handler.Logout)
	v1.POST("user/refresh", util.AuthMiddleware(handler.rdb), handler.Refresh)
	v1.POST("user/do-completion", util.AuthMiddleware(handler.rdb), handler.DoCompletion)
	v1.POST("user/get-completion", util.AuthMiddleware(handler.rdb), handler.GetCompletion)
	v1.DELETE("user-delete", util.AuthMiddleware(handler.rdb), handler.Delete)
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

func (handler *UserHandler) GetTable(c *gin.Context) {
	handler.Usecase.GetTable(c)
}

func (handler *UserHandler) IsComplete(c *gin.Context) {
	handler.Usecase.IsComplete(c)
}

func (handler *UserHandler) IsVerified(c *gin.Context) {
	handler.Usecase.IsVerified(c)
}

func (handler *UserHandler) Login(c *gin.Context) {
	handler.Usecase.Login(c)
}

func (handler *UserHandler) VerifyReCaptcha(c *gin.Context) {
	handler.Usecase.VerifyReCaptcha(c)
}

func (handler *UserHandler) VerifyEmail(c *gin.Context) {
	handler.Usecase.VerifyEmail(c)
}

func (handler *UserHandler) Register(c *gin.Context) {
	handler.Usecase.Register(c)
}

func (handler *UserHandler) Refresh(c *gin.Context) {
	handler.Usecase.Refresh(c)
}

func (handler *UserHandler) DoCompletion(c *gin.Context) {
	handler.Usecase.DoCompletion(c)
}

func (handler *UserHandler) GetCompletion(c *gin.Context) {
	handler.Usecase.GetCompletion(c)
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

func (handler *UserHandler) GetByID(c *gin.Context) {
	handler.Usecase.GetByID(c)
}

func (handler *UserHandler) IsRegistered(c *gin.Context) {
	handler.Usecase.IsRegistered(c)
}

func (handler *UserHandler) IsNullPassword(c *gin.Context) {
	handler.Usecase.IsNullPassword(c)
}

func (handler *UserHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
