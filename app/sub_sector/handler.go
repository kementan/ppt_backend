package sub_sector

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type SubSectorHandler struct {
	Usecase SubSectorUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase SubSectorUsecase, rdb *redis.Client) {
	handler := &SubSectorHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("sub-sector-table", util.AuthMiddleware(handler.rdb), handler.GetTable)
	v1.GET("sub-sector-list", util.AuthMiddleware(handler.rdb), handler.GetList)
	v1.PUT("sub-sector-update", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("sub-sector-id", util.AuthMiddleware(handler.rdb), handler.GetByID)
	v1.POST("sub-sector-create", util.AuthMiddleware(handler.rdb), handler.Create)
	v1.DELETE("sub-sector-delete", util.AuthMiddleware(handler.rdb), handler.Delete)
}

func (handler *SubSectorHandler) GetTable(c *gin.Context) {
	handler.Usecase.GetTable(c)
}

func (handler *SubSectorHandler) GetList(c *gin.Context) {
	handler.Usecase.GetList(c)
}

func (handler *SubSectorHandler) GetByID(c *gin.Context) {
	handler.Usecase.GetByID(c)
}

func (handler *SubSectorHandler) Create(c *gin.Context) {
	handler.Usecase.Create(c)
}

func (handler *SubSectorHandler) Update(c *gin.Context) {
	handler.Usecase.Update(c)
}

func (handler *SubSectorHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
