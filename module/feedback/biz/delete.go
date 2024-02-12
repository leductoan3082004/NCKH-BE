package feedbackbiz

import (
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nckh-BE/appCommon"
	feedbackmodel "nckh-BE/module/feedback/model"
)

type feedbackDeleteStore interface {
	DeleteMany(ctx context.Context, conditions interface{}) error
}

type feedbackDeleteBiz struct {
	store  feedbackDeleteStore
	logger logger.Logger
}

func NewFeedbackDeleteBiz(store feedbackDeleteStore) *feedbackDeleteBiz {
	return &feedbackDeleteBiz{store: store, logger: logger.GetCurrent().GetLogger("FeedbackDeleteBiz")}
}

func (biz *feedbackDeleteBiz) Delete(ctx context.Context, filter *feedbackmodel.FeedbackDelete) error {
	ids := make([]primitive.ObjectID, len(filter.FeedbackIds))

	for i, item := range filter.FeedbackIds {
		id, err := primitive.ObjectIDFromHex(item)
		if err != nil {
			return err
		}
		ids[i] = id
	}

	if err := biz.store.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": ids}}); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(feedbackmodel.EntityName, err)
	}

	return nil
}
