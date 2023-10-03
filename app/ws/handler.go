package ws

import (
	"log"
	"net/http"

	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func NewHandler(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("get-notification", wsHandler)
	v1.GET("get-kafka", kfHandler)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:4200"
	},
}

func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Could not upgrade WebSocket connection: %v", err)
		return
	}
	defer conn.Close()

	go func() {
		for {
			message := <-util.WsChannel
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Printf("Error writing WebSocket message: %v", err)
				return
			}
		}
	}()
	select {}
}

func kfHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Could not upgrade WebSocket connection: %v", err)
		return
	}
	defer conn.Close()

	util.ConsumeKafkaMessage("notifications-topic")

	go func() {
		for {
			message := <-util.KfChannel
			log.Printf("message: %v", message)
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Printf("Error writing WebSocket message: %v", err)
				return
			}
		}
	}()
	select {}
}
