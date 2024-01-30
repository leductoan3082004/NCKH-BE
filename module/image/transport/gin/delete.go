package imagegin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"nckh-BE/appCommon"
	imagebiz "nckh-BE/module/image/biz"
	imagemodel "nckh-BE/module/image/model"
	imagestorage "nckh-BE/module/image/storage"
	"net/http"
)

func Delete(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data imagemodel.ImageDelete
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		store := imagestorage.NewMgDBStore(db)
		biz := imagebiz.NewDeleteBiz(store)
		if err := biz.DeleteImage(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
