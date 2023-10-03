package report_category

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ReportCategoryHandler struct {
	Usecase ReportCategoryUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase ReportCategoryUsecase, rdb *redis.Client) {
	handler := &ReportCategoryHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("report-category-table", util.AuthMiddleware(handler.rdb), handler.GetTable)
	v1.PUT("report-category-update", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("report-category-id", util.AuthMiddleware(handler.rdb), handler.GetByID)
	v1.POST("report-category-create", util.AuthMiddleware(handler.rdb), handler.Create)
	v1.DELETE("report-category-delete", util.AuthMiddleware(handler.rdb), handler.Delete)
}

func (handler *ReportCategoryHandler) GetTable(c *gin.Context) {
	handler.Usecase.GetTable(c)
}

func (handler *ReportCategoryHandler) GetByID(c *gin.Context) {
	handler.Usecase.GetByID(c)
}

func (handler *ReportCategoryHandler) Create(c *gin.Context) {
	handler.Usecase.Create(c)
}

func (handler *ReportCategoryHandler) Update(c *gin.Context) {
	handler.Usecase.Update(c)
}

func (handler *ReportCategoryHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
