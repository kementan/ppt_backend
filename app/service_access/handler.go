package service_access

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ServiceAccessHandler struct {
	Usecase ServiceAccessUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase ServiceAccessUsecase, rdb *redis.Client) {
	handler := &ServiceAccessHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("service-access", util.AuthMiddleware(handler.rdb), handler.Read)
	v1.PUT("service-access", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("service-access", util.AuthMiddleware(handler.rdb), handler.Create)
	v1.DELETE("service-access", util.AuthMiddleware(handler.rdb), handler.Delete)
}

func (handler *ServiceAccessHandler) Create(c *gin.Context) {
	handler.Usecase.Create(c)
}

func (handler *ServiceAccessHandler) Read(c *gin.Context) {
	handler.Usecase.Read(c)
}

func (handler *ServiceAccessHandler) Update(c *gin.Context) {
	handler.Usecase.Update(c)
}

func (handler *ServiceAccessHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
