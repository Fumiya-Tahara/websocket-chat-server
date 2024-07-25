package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer(r *gin.Engine, handler HandlerInterface) {
	r.GET("/connection", handler.Connection)
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	err := r.Run()
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
