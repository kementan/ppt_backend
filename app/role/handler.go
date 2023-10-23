package role

import (
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RoleHandler struct {
	Usecase RoleUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase RoleUsecase, rdb *redis.Client) {
	handler := &RoleHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.GET("role-validation", util.AuthMiddleware(handler.rdb), handler.RoleValidation)

	v1.GET("role-table", util.AuthMiddleware(handler.rdb), handler.GetTable)
	v1.GET("role-8sd234kh34-list", util.AuthMiddleware(handler.rdb), handler.GetPublicRole)
	v1.GET("role-2238sdas9d-list", util.AuthMiddleware(handler.rdb), handler.GetAllRole)
	v1.PUT("role-update", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("role-id", util.AuthMiddleware(handler.rdb), handler.GetByID)
	v1.POST("role-create", util.AuthMiddleware(handler.rdb), handler.Create)
	v1.DELETE("role-delete", util.AuthMiddleware(handler.rdb), handler.Delete)
}

func (handler *RoleHandler) RoleValidation(c *gin.Context) {
	handler.Usecase.RoleValidation(c)
}

func (handler *RoleHandler) GetTable(c *gin.Context) {
	handler.Usecase.GetTable(c)
}

func (handler *RoleHandler) GetPublicRole(c *gin.Context) {
	handler.Usecase.GetPublicRole(c)
}

func (handler *RoleHandler) GetAllRole(c *gin.Context) {
	handler.Usecase.GetAllRole(c)
}

func (handler *RoleHandler) GetByID(c *gin.Context) {
	handler.Usecase.GetByID(c)
}

func (handler *RoleHandler) Create(c *gin.Context) {
	handler.Usecase.Create(c)
}

func (handler *RoleHandler) Update(c *gin.Context) {
	handler.Usecase.Update(c)
}

func (handler *RoleHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
