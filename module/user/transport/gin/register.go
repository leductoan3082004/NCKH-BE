package usergin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"nckh-BE/appCommon"
	userbiz "nckh-BE/module/user/biz"
	usermodel "nckh-BE/module/user/model"
	userstorage "nckh-BE/module/user/storage"
	"net/http"
)

func Register(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserRegister
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)

		store := userstorage.NewMgDBStore(db)
		biz := userbiz.NewUserRegisterBiz(store)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
