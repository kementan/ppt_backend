package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.com/xsysproject/schemes"
)

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Data interface{}) {

	jsonResponse := schemes.SchemeResponses{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func ErrorResponse(ctx *gin.Context, Error interface{}) {
	err := schemes.SchemeErrorResponse{
		StatusCode: http.StatusBadRequest,
		Error:      Error,
	}

	ctx.AbortWithStatusJSON(err.StatusCode, err)
}
