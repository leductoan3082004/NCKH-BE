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

func List(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging appCommon.Paging
		var filter postmodel.PostList
		if err := c.ShouldBind(&paging); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&filter); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		store := poststorage.NewMgDBStore(db)
		biz := postbiz.NewPostListBiz(store)
		res, err := biz.ListDataWithCondition(c.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, paging, nil))
	}
}
