package service

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ServiceHandler struct {
	Usecase ServiceUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase ServiceUsecase, rdb *redis.Client) {
	handler := &ServiceHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("service-export", util.AuthMiddleware(handler.rdb), handler.Export)
	v1.GET("service-table", util.AuthMiddleware(handler.rdb), handler.GetTable)
	v1.PUT("service-update", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("service-id", util.AuthMiddleware(handler.rdb), handler.GetByID)
	v1.POST("service-create", util.AuthMiddleware(handler.rdb), handler.Create)
	v1.DELETE("service-delete", util.AuthMiddleware(handler.rdb), handler.Delete)
}

func (handler *ServiceHandler) Export(c *gin.Context) {
	handler.Usecase.Export(c)
}

func (handler *ServiceHandler) GetTable(c *gin.Context) {
	handler.Usecase.GetTable(c)
}

func (handler *ServiceHandler) GetByID(c *gin.Context) {
	handler.Usecase.GetByID(c)
}

func (handler *ServiceHandler) Create(c *gin.Context) {
	handler.Usecase.Create(c)
}

func (handler *ServiceHandler) Update(c *gin.Context) {
	handler.Usecase.Update(c)
}

func (handler *ServiceHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
