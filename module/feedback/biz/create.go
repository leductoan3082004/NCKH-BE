package feedbackbiz

import (
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"nckh-BE/appCommon"
	feedbackmodel "nckh-BE/module/feedback/model"
)

type feedbackCreateStore interface {
	Create(ctx context.Context, data *feedbackmodel.Feedback) error
}

type feedbackCreateBiz struct {
	store  feedbackCreateStore
	logger logger.Logger
}

func NewFeedbackCreateBiz(store feedbackCreateStore) *feedbackCreateBiz {
	return &feedbackCreateBiz{store: store, logger: logger.GetCurrent().GetLogger("FeedbackCreateBiz")}
}

func (biz *feedbackCreateBiz) CreateFeedback(ctx context.Context, data *feedbackmodel.FeedbackCreate) error {
	createData := &feedbackmodel.Feedback{
		Content: data.Content,
		Topic:   data.Topic,
		Email:   data.Email,
		Phone:   data.Phone,
		Name:    data.Name,
	}

	if err := biz.store.Create(ctx, createData); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotCreateEntity(feedbackmodel.EntityName, err)
	}

	return nil
}
