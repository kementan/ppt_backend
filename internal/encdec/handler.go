package encdec

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/xsysproject/ppt_backend/helper"
)

func NewHandler(router *gin.Engine) {
	router.GET("enc_dec", EncDec)
}

func EncDec(ctx *gin.Context) {
	var req StringRequest
	var result string

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.Type == "enc" {
		result, _ = helper.Encrypt(req.Value)
	} else {
		result, _ = helper.Decrypt(req.Value)
	}

	ctx.JSON(http.StatusOK, result)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
