package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lequocbinh04/go-sdk/sdkcm"
	"nckh-BE/appCommon"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*sdkcm.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}

				appErr := appCommon.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
