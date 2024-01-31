package middleware

import (
	"github.com/gin-gonic/gin"
	"nckh-BE/appCommon"
	usermodel "nckh-BE/module/user/model"
)

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		if user.IsAdmin == false {
			panic(appCommon.ErrNoPermission(nil))
		}
		c.Next()
	}
}
