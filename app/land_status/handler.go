package land_status

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type LandStatusHandler struct {
	Usecase LandStatusUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase LandStatusUsecase, rdb *redis.Client) {
	handler := &LandStatusHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("land-status-table", util.AuthMiddleware(handler.rdb), handler.GetTable)
	v1.PUT("land-status-update", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("land-status-id", util.AuthMiddleware(handler.rdb), handler.GetByID)
	v1.POST("land-status-create", util.AuthMiddleware(handler.rdb), handler.Create)
	v1.DELETE("land-status-delete", util.AuthMiddleware(handler.rdb), handler.Delete)
}

func (handler *LandStatusHandler) GetTable(c *gin.Context) {
	handler.Usecase.GetTable(c)
}

func (handler *LandStatusHandler) GetByID(c *gin.Context) {
	handler.Usecase.GetByID(c)
}

func (handler *LandStatusHandler) Create(c *gin.Context) {
	handler.Usecase.Create(c)
}

func (handler *LandStatusHandler) Update(c *gin.Context) {
	handler.Usecase.Update(c)
}

func (handler *LandStatusHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
