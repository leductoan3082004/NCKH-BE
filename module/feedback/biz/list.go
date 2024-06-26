package feedbackbiz

import (
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson"
	"nckh-BE/appCommon"
	feedbackmodel "nckh-BE/module/feedback/model"
)

type feedBackListStore interface {
	ListDataWithCondition(ctx context.Context, condition bson.M, paging *appCommon.Paging, moreInfo ...string) ([]feedbackmodel.Feedback, error)
}

type feedbackListBiz struct {
	store  feedBackListStore
	logger logger.Logger
}

func NewFeedbackListBiz(store feedBackListStore) *feedbackListBiz {
	return &feedbackListBiz{store: store, logger: logger.GetCurrent().GetLogger("FeedBackListBiz")}
}

func (biz *feedbackListBiz) List(ctx context.Context, paging *appCommon.Paging) ([]feedbackmodel.Feedback, error) {
	paging.Fulfill()
	res, err := biz.store.ListDataWithCondition(ctx, bson.M{}, paging)
	if err != nil {
		biz.logger.Errorln("Failed to list feedback", err)
		return []feedbackmodel.Feedback{}, appCommon.ErrCannotListEntity(feedbackmodel.EntityName, err)
	}

	if len(res) > 0 {
		paging.NextCursor = res[len(res)-1].Id.Hex()
	}
	return res, nil
}
