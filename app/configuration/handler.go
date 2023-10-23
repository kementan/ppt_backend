package configuration

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ConfigurationHandler struct {
	Usecase ConfigurationUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase ConfigurationUsecase, rdb *redis.Client) {
	handler := &ConfigurationHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("configuration-hd", handler.GetHD)
	v1.GET("configuration-table", util.AuthMiddleware(handler.rdb), handler.GetTable)
	v1.PUT("configuration-update", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("configuration-id", util.AuthMiddleware(handler.rdb), handler.GetByID)
	v1.POST("configuration-create", util.AuthMiddleware(handler.rdb), handler.Create)
	v1.DELETE("configuration-delete", util.AuthMiddleware(handler.rdb), handler.Delete)
}

func (handler *ConfigurationHandler) GetTable(c *gin.Context) {
	handler.Usecase.GetTable(c)
}

func (handler *ConfigurationHandler) GetHD(c *gin.Context) {
	handler.Usecase.GetHD(c)
}

func (handler *ConfigurationHandler) GetByID(c *gin.Context) {
	handler.Usecase.GetByID(c)
}

func (handler *ConfigurationHandler) Create(c *gin.Context) {
	handler.Usecase.Create(c)
}

func (handler *ConfigurationHandler) Update(c *gin.Context) {
	handler.Usecase.Update(c)
}

func (handler *ConfigurationHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
