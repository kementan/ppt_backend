package encdec

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func NewHandler(router *gin.Engine, rdb *redis.Client) {
	v1 := router.Group("/v1")

	v1.GET("enc_dec", util.AuthMiddleware(rdb), EncDec)
	// v1.GET("enc-dec", EncDec)
}

func EncDec(c *gin.Context) {
	var req StringRequest
	var result string

	err := util.ValidateRequest(c)
	if err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.JERR(c, http.StatusBadRequest, err)
		return
	}

	if req.Type == "enc" {
		result, _ = util.Encrypt(req.Value, req.T)
	} else {
		result, _ = util.Decrypt(req.Value, req.T)
	}

	message := map[string]any{
		"type":    "encdec",
		"content": req.Value,
		"time":    time.Now(),
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error encoding message to JSON: %v", err)
		return
	}

	util.WsChannel <- string(jsonMessage)
	util.ProduceKafkaMessage("notifications-topic", string(jsonMessage))

	util.JOK(c, http.StatusOK, result)
}
