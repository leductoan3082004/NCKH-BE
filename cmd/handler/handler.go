package handler

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
)

func MainRoute(router *gin.Engine, sc goservice.ServiceContext) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "pong",
		})
	})
}
