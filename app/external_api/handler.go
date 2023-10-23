package external_api

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ExternalApiHandler struct {
	Usecase ExternalApiUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase ExternalApiUsecase, rdb *redis.Client) {
	handler := &ExternalApiHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("api-openweathermap", handler.GetWeather)
}

func (handler *ExternalApiHandler) GetWeather(c *gin.Context) {
	handler.Usecase.GetWeather(c)
}
