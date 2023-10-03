package commodity

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type CommodityHandler struct {
	Usecase CommodityUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase CommodityUsecase, rdb *redis.Client) {
	handler := &CommodityHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("commodity-table", util.AuthMiddleware(handler.rdb), handler.GetTable)
	v1.PUT("commodity-update", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("commodity-id", util.AuthMiddleware(handler.rdb), handler.GetByID)
	v1.POST("commodity-create", util.AuthMiddleware(handler.rdb), handler.Create)
	v1.DELETE("commodity-delete", util.AuthMiddleware(handler.rdb), handler.Delete)
}

func (handler *CommodityHandler) GetTable(c *gin.Context) {
	handler.Usecase.GetTable(c)
}

func (handler *CommodityHandler) GetByID(c *gin.Context) {
	handler.Usecase.GetByID(c)
}

func (handler *CommodityHandler) Create(c *gin.Context) {
	handler.Usecase.Create(c)
}

func (handler *CommodityHandler) Update(c *gin.Context) {
	handler.Usecase.Update(c)
}

func (handler *CommodityHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
