package controller

import "github.com/gin-gonic/gin"

type HandlerInterface interface {
	Connection(c *gin.Context)
}
