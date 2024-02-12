package feedbackbiz

import (
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nckh-BE/appCommon"
	feedbackmodel "nckh-BE/module/feedback/model"
)

type feedbackFindStore interface {
	Find(ctx context.Context, conditions interface{}, fields ...string) (*feedbackmodel.Feedback, error)
}
type feedbackFindBiz struct {
	store  feedbackFindStore
	logger logger.Logger
}

func NewFeedbackFindBiz(store feedbackFindStore) *feedbackFindBiz {
	return &feedbackFindBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("FindFeedbackBiz"),
	}
}

func (biz *feedbackFindBiz) Find(ctx context.Context, filter *feedbackmodel.FeedbackFind) (*feedbackmodel.Feedback, error) {
	feedbackId, err := primitive.ObjectIDFromHex(filter.FeedbackId)

	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	data, err := biz.store.Find(ctx, bson.M{
		"_id": feedbackId,
	})

	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return nil, appCommon.ErrEntityNotFound(feedbackmodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(feedbackmodel.EntityName, err)
	}
	return data, nil
}
