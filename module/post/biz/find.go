package postbiz

import (
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nckh-BE/appCommon"
	postmodel "nckh-BE/module/post/model"
)

type findPostStore interface {
	Find(ctx context.Context, conditions interface{}, moreInfo ...string) (*postmodel.Post, error)
}

type FindPostBiz struct {
	store  findPostStore
	logger logger.Logger
}

func NewFindPostBiz(store findPostStore) *FindPostBiz {
	return &FindPostBiz{store: store, logger: logger.GetCurrent().GetLogger("FindPostBiz")}
}

func (biz *FindPostBiz) Find(ctx context.Context, filter *postmodel.PostFind) (*postmodel.Post, error) {
	postId, err := primitive.ObjectIDFromHex(filter.PostId)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	post, err := biz.store.Find(ctx, map[string]interface{}{"_id": postId})
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return nil, appCommon.ErrEntityNotFound(postmodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(postmodel.EntityName, err)
	}
	return post, nil
}
