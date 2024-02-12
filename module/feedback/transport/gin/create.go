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

func Create(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data feedbackmodel.FeedbackCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		store := feedbackstorage.NewMgDBStore(db)
		biz := feedbackbiz.NewFeedbackCreateBiz(store)
		if err := biz.CreateFeedback(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
