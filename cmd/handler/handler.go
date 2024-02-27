package handler

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"nckh-BE/middleware"
	feedbackgin "nckh-BE/module/feedback/transport/gin"
	imagegin "nckh-BE/module/image/transport/gin"
	postgin "nckh-BE/module/post/transport/gin"
	usergin "nckh-BE/module/user/transport/gin"
)

func MainRoute(router *gin.Engine, sc goservice.ServiceContext) {
	router.Use(middleware.AllowCORS())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ping": "pong",
		})
	})
	router.Use(middleware.Recover())
	router.POST("v1/user/register", usergin.Register(sc))
	router.POST("v1/user/login", usergin.Login(sc))

	v1 := router.Group("/v1")
	authedRoutes := router.Group("v1", middleware.RequiredAuth(sc))

	post := authedRoutes.Group("/post", middleware.AdminAuthorization())
	{
		post.POST("/", postgin.Create(sc))
		post.PUT("/", postgin.Update(sc))
		post.DELETE("/", postgin.Delete(sc))
	}
	v1.GET("/post", postgin.List(sc))
	v1.GET("/post/:id", postgin.Find(sc))
	v1.GET("/post/suggestion", postgin.ListRandom(sc))

	image := authedRoutes.Group("image", middleware.AdminAuthorization())
	{
		image.POST("/", imagegin.UploadByFile(sc))
		image.DELETE("/", imagegin.Delete(sc))
		image.GET("/", imagegin.List(sc))
	}

	feedback := v1.Group("/feedback")
	{
		feedback.POST("/", feedbackgin.Create(sc))
		feedback.GET(
			"/",
			middleware.RequiredAuth(sc),
			middleware.AdminAuthorization(),
			feedbackgin.List(sc),
		)
		feedback.DELETE(
			"/",
			middleware.RequiredAuth(sc),
			middleware.AdminAuthorization(),
			feedbackgin.Delete(sc),
		)
		feedback.GET(
			"/:id",
			middleware.RequiredAuth(sc),
			middleware.AdminAuthorization(),
			feedbackgin.Find(sc),
		)
	}
}
