package postbiz

import (
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nckh-BE/appCommon"
	postmodel "nckh-BE/module/post/model"
)

type deleteStore interface {
	Delete(ctx context.Context, conditions interface{}) error
}

type DeleteBiz struct {
	store  deleteStore
	logger logger.Logger
}

func NewDeleteBiz(store deleteStore) *DeleteBiz {
	return &DeleteBiz{store: store, logger: logger.GetCurrent().GetLogger("PostDeleteBiz")}
}
func (biz *DeleteBiz) Delete(ctx context.Context, filter *postmodel.PostDelete) error {
	postId, err := primitive.ObjectIDFromHex(filter.PostId)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	conditions := bson.M{"_id": postId}
	if err := biz.store.Delete(ctx, conditions); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(postmodel.EntityName, err)
	}
	return nil
}
