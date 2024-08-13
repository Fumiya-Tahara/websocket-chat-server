package main

import (
	"github.com/Fumiya-Tahara/websocket-chat-server/internal/controller"
	"github.com/Fumiya-Tahara/websocket-chat-server/internal/controller/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	r := gin.Default()

	clients := make(map[*websocket.Conn]bool)
	broadcast := make(chan []byte)

	wsHandler := ws.NewWebsocketHandler(clients, broadcast)

	handler := controller.NewHandler(wsHandler)

	controller.StartServer(r, handler)
}
