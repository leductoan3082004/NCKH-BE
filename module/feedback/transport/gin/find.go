package feedbackgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"nckh-BE/appCommon"
	feedbackbiz "nckh-BE/module/feedback/biz"
	feedbackmodel "nckh-BE/module/feedback/model"
	feedbackstorage "nckh-BE/module/feedback/storage"
	"net/http"
)

func Find(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		store := feedbackstorage.NewMgDBStore(db)
		biz := feedbackbiz.NewFeedbackFindBiz(store)
		res, err := biz.Find(c.Request.Context(), &feedbackmodel.FeedbackFind{FeedbackId: id})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
