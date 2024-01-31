package handler

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"nckh-BE/middleware"
	imagegin "nckh-BE/module/image/transport/gin"
	postgin "nckh-BE/module/post/transport/gin"
	usergin "nckh-BE/module/user/transport/gin"
)

func MainRoute(router *gin.Engine, sc goservice.ServiceContext) {
	router.Use(middleware.AllowCORS())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "pong",
		})
	})
	router.Use(middleware.Recover())
	router.POST("v1/register", usergin.Register(sc))
	router.POST("v1/login", usergin.Login(sc))

	authedRoutes := router.Group("v1", middleware.RequiredAuth(sc))

	post := authedRoutes.Group("/post", middleware.AdminAuthorization())
	{
		post.POST("/", postgin.Create(sc))
		post.PUT("/", postgin.Update(sc))
		post.DELETE("/", postgin.Delete(sc))
	}
	authedRoutes.GET("/post", postgin.List(sc))
	authedRoutes.GET("/post/:id", postgin.Find(sc))
	image := authedRoutes.Group("image", middleware.AdminAuthorization())
	{
		image.POST("/", imagegin.UploadByFile(sc))
		image.DELETE("/", imagegin.Delete(sc))
		image.GET("/", imagegin.List(sc))
	}

}
