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

func ListRandom(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data postmodel.PostSimilarList
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		store := poststorage.NewMgDBStore(db)
		biz := postbiz.NewListSimilarPostBiz(store)
		res, err := biz.ListSimilarPost(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
