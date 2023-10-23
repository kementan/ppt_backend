package module_generator

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ModuleGeneratorHandler struct {
	Usecase ModuleGeneratorUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase ModuleGeneratorUsecase, rdb *redis.Client) {
	handler := &ModuleGeneratorHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("module-generator-table", util.AuthMiddleware(handler.rdb), handler.GetTable)
	v1.GET("module-generator-list", util.AuthMiddleware(handler.rdb), handler.GetList)
	v1.PUT("module-generator-update", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("module-generator-id", util.AuthMiddleware(handler.rdb), handler.GetByID)
	v1.POST("module-generator-create", util.AuthMiddleware(handler.rdb), handler.Create)
	v1.DELETE("module-generator-delete", util.AuthMiddleware(handler.rdb), handler.Delete)
}

func (handler *ModuleGeneratorHandler) GetTable(c *gin.Context) {
	handler.Usecase.GetTable(c)
}

func (handler *ModuleGeneratorHandler) GetList(c *gin.Context) {
	handler.Usecase.GetList(c)
}

func (handler *ModuleGeneratorHandler) GetByID(c *gin.Context) {
	handler.Usecase.GetByID(c)
}

func (handler *ModuleGeneratorHandler) Create(c *gin.Context) {
	handler.Usecase.Create(c)
}

func (handler *ModuleGeneratorHandler) Update(c *gin.Context) {
	handler.Usecase.Update(c)
}

func (handler *ModuleGeneratorHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
