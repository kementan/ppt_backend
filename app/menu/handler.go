package menu

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type MenuHandler struct {
	Usecase MenuUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase MenuUsecase, rdb *redis.Client) {
	handler := &MenuHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("menu-export", util.AuthMiddleware(handler.rdb), handler.Export)
	v1.GET("menu-table", handler.GetTable)
	v1.PUT("menu-update", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("menu-id", util.AuthMiddleware(handler.rdb), handler.GetByID)
	v1.POST("menu-create", util.AuthMiddleware(handler.rdb), handler.Create)
	v1.DELETE("menu-delete", util.AuthMiddleware(handler.rdb), handler.Delete)
}

func (handler *MenuHandler) Export(c *gin.Context) {
	handler.Usecase.Export(c)
}

func (handler *MenuHandler) GetTable(c *gin.Context) {
	handler.Usecase.GetTable(c)
}

func (handler *MenuHandler) GetByID(c *gin.Context) {
	handler.Usecase.GetByID(c)
}

func (handler *MenuHandler) Create(c *gin.Context) {
	handler.Usecase.Create(c)
}

func (handler *MenuHandler) Update(c *gin.Context) {
	handler.Usecase.Update(c)
}

func (handler *MenuHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
