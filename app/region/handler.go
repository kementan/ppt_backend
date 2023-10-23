package region

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RegionHandler struct {
	Usecase RegionUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase RegionUsecase, rdb *redis.Client) {
	handler := &RegionHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("region-list/:level/:parentCode", util.AuthMiddleware(handler.rdb), handler.GetList)
	v1.GET("region-by-kode/:code", util.AuthMiddleware(handler.rdb), handler.GetRegion)
}

func (handler *RegionHandler) GetList(c *gin.Context) {
	handler.Usecase.GetList(c)
}

func (handler *RegionHandler) GetRegion(c *gin.Context) {
	handler.Usecase.GetRegion(c)
}
