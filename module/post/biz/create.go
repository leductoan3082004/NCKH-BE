package postbiz

import (
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"nckh-BE/appCommon"
	postmodel "nckh-BE/module/post/model"
)

type createStore interface {
	Create(ctx context.Context, data *postmodel.Post) error
}

type createBiz struct {
	store  createStore
	logger logger.Logger
}

func NewCreateBiz(store createStore) *createBiz {
	return &createBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("PostCreateBiz"),
	}
}

func (biz *createBiz) Create(ctx context.Context, data *postmodel.PostCreate) (*postmodel.Post, error) {
	createData := &postmodel.Post{
		Title:   data.Title,
		Content: data.Content,
		Author:  data.Author,
	}
	if err := biz.store.Create(ctx, createData); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotCreateEntity(postmodel.EntityName, err)
	}
	return createData, nil
}
