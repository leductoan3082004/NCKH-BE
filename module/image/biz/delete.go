package imagebiz

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nckh-BE/appCommon"
	imagemodel "nckh-BE/module/image/model"
)

type deleteImageStore interface {
	Delete(ctx context.Context, conditions interface{}) error
}

type deleteImageBiz struct {
	store  deleteImageStore
	logger logger.Logger
}

func NewDeleteBiz(store deleteImageStore) *deleteImageBiz {
	return &deleteImageBiz{store: store, logger: logger.GetCurrent().GetLogger("DeleteImageBiz")}
}

func (biz *deleteImageBiz) DeleteImage(ctx context.Context, filter *imagemodel.ImageDelete) error {
	imageId, err := primitive.ObjectIDFromHex(filter.ImageId)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	if err := biz.store.Delete(ctx, bson.M{
		"_id": imageId,
	}); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(imagemodel.EntityName, err)
	}
	return nil
}
