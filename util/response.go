package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func JOK(c *gin.Context, status int, data interface{}) {
	response := Response{
		Status: status,
		Data:   data,
	}

	c.JSON(status, response)
}

func JERR(c *gin.Context, status int, err error) {
	response := Response{
		Status:  status,
		Message: errorResponse(err),
	}

	c.JSON(status, response)
}

func ValidateRequest(c *gin.Context) error {
	contentLength := c.Request.ContentLength
	if contentLength == 0 {
		return fmt.Errorf("request body is required")
	}
	return nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func MergeErrors(err1, err2 error) error {
	return fmt.Errorf("%w: %v", err1, err2)
}
