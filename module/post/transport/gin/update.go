package postgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"nckh-BE/appCommon"
	postbiz "nckh-BE/module/post/biz"
	postmodel "nckh-BE/module/post/model"
	poststorage "nckh-BE/module/post/storage"
	"net/http"
)

func Update(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data postmodel.PostUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		store := poststorage.NewMgDBStore(db)
		biz := postbiz.NewUpdateBiz(store)
		if err := biz.Update(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
