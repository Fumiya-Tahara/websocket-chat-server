package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Fumiya-Tahara/websocket-chat-server/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
}

func NewWebsocketHandler(clients map[*websocket.Conn]bool, broadcast chan []byte) *WebsocketHandler {
	return &WebsocketHandler{
		clients:   clients,
		broadcast: broadcast,
	}
}

func (wh *WebsocketHandler) HandleWebSocket(ctx *gin.Context) {
	upgrader := websocket.Upgrader{}

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("could not upgrade connection:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not upgrade connection"})

		return
	}
	defer ws.Close()

	wh.clients[ws] = true

	for {
		var message controller.Message

		err := ws.ReadJSON(&message)
		if err != nil {
			log.Println("could not read message:", err)
			delete(wh.clients, ws)

			break
		}

		messageBytes, err := json.Marshal(message)
		if err != nil {
			log.Println("could not marshal JSON:", err)

			continue
		}

		wh.broadcast <- messageBytes
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "connection closed"})
}

func (wh *WebsocketHandler) HandleMessages() {
	for {
		data := <-wh.broadcast

		var message controller.Message

		err := json.Unmarshal(data, &message)
		if err != nil {
			log.Println("could not unmarshal JSON:", err)

			continue
		}

		for client := range wh.clients {
			err := client.WriteJSON(message)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(wh.clients, client)
			}
		}
	}
}
