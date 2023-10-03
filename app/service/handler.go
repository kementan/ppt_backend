package service

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ServiceHandler struct {
	Usecase ServiceUsecase
	rdb     *redis.Client
}

func NewHandler(router *gin.Engine, usecase ServiceUsecase, rdb *redis.Client) {
	handler := &ServiceHandler{
		Usecase: usecase,
		rdb:     rdb,
	}

	v1 := router.Group("/v1")

	v1.POST("service-image", util.AuthMiddleware(handler.rdb), handler.Upload)

	v1.GET("service", util.AuthMiddleware(handler.rdb), handler.Read)
	v1.PUT("service", util.AuthMiddleware(handler.rdb), handler.Update)
	v1.POST("service", util.AuthMiddleware(handler.rdb), handler.Create)
	v1.DELETE("service", util.AuthMiddleware(handler.rdb), handler.Delete)
}

func (handler *ServiceHandler) Upload(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form"})
		return
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get file from form"})
		return
	}
	defer file.Close()

	currentDir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "???"})
		return
	}
	savePath := filepath.Join(currentDir, "ppt_nginx/uploads/")

	if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create upload directory"})
		return
	}

	outFile, err := os.Create(savePath + "uploaded_image.png")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create file on server"})
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to copy file data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
}

func (handler *ServiceHandler) Create(c *gin.Context) {
	handler.Usecase.Create(c)
}

func (handler *ServiceHandler) Read(c *gin.Context) {
	handler.Usecase.Read(c)
}

func (handler *ServiceHandler) Update(c *gin.Context) {
	handler.Usecase.Update(c)
}

func (handler *ServiceHandler) Delete(c *gin.Context) {
	handler.Usecase.Delete(c)
}
