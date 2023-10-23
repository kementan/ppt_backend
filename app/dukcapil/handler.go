package dukcapil

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type DukcapilHandler struct {
	rdb *redis.Client
}

func NewHandler(router *gin.Engine, rdb *redis.Client) {
	handler := &DukcapilHandler{
		rdb: rdb,
	}

	v1 := router.Group("/v1")

	v1.POST("id-validation", util.AuthMiddleware(handler.rdb), handler.IdValidation)
	// v1.POST("id-validation", handler.IdValidation)
}

func (handler *DukcapilHandler) IdValidation(c *gin.Context) {
	IdValidation(c)
}
