package controller

import "github.com/gin-gonic/gin"

type Handler struct {
	WebsocketHandler WebsocketHandlerInterface
}

func NewHandler(wh WebsocketHandlerInterface) *Handler {
	return &Handler{
		WebsocketHandler: wh,
	}
}

func (h *Handler) Connection(c *gin.Context) {
	h.WebsocketHandler.HandleWebSocket(c)
}

func (h *Handler) BroadcastMessages() {
	h.WebsocketHandler.HandleMessages()
}
