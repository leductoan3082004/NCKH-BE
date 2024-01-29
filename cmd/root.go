package cmd

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"github.com/spf13/cobra"
	"net/http"
)

func newService() goservice.Service {

	service := goservice.New(
		goservice.WithName("mindzone"),
		goservice.WithVersion("1.0.0"),
	)

	if err := service.Init(); err != nil {
		panic(err)
	}
	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start a mindzone service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.GET("/health", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
				})
			})
		})

		if err := service.Start(func() {}); err != nil {
			panic(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
