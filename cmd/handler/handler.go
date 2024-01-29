package handler

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"nckh-BE/middleware"
	postgin "nckh-BE/module/post/transport/gin"
	usergin "nckh-BE/module/user/transport/gin"
)

func MainRoute(router *gin.Engine, sc goservice.ServiceContext) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "pong",
		})
	})
	router.Use(middleware.Recover())
	router.POST("v1/register", usergin.Register(sc))
	router.POST("v1/login", usergin.Login(sc))

	authedRoutes := router.Group("v1")
	{
		authedRoutes.POST("/", postgin.Create(sc))
	}
}
