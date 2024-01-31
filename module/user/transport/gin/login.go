package usergin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"nckh-BE/appCommon"
	userbiz "nckh-BE/module/user/biz"
	usermodel "nckh-BE/module/user/model"
	userstorage "nckh-BE/module/user/storage"
	"nckh-BE/plugin/tokenprovider"
	"net/http"
)

func Login(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserLogin
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		tokenProvider := sc.MustGet(appCommon.PluginJWT).(tokenprovider.Provider)
		store := userstorage.NewMgDBStore(db)
		biz := userbiz.NewUserLoginBiz(store, tokenProvider)
		token, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(token))
	}
}
