package controller

import "github.com/gin-gonic/gin"

func StartServer(r *gin.Engine, handler HandlerInterface) {
	r.GET("/connection", handler.Connection)
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.Run()
}
