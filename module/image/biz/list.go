package imagebiz

import (
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson"
	"nckh-BE/appCommon"
	imagemodel "nckh-BE/module/image/model"
)

type listStore interface {
	ListDataWithCondition(ctx context.Context, condition bson.M, paging *appCommon.Paging, moreInfo ...string) ([]imagemodel.Image, error)
}
type listBiz struct {
	store  listStore
	logger logger.Logger
}

func NewListBiz(store listStore) *listBiz {
	return &listBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("ImageListBiz"),
	}
}

func (biz *listBiz) ListDataWithCondition(ctx context.Context, paging *appCommon.Paging, moreInfo ...string) ([]imagemodel.Image, error) {
	paging.Fulfill()
	result, err := biz.store.ListDataWithCondition(ctx, nil, paging, moreInfo...)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return []imagemodel.Image{}, appCommon.ErrCannotListEntity(imagemodel.EntityName, err)
	}
	return result, nil
}
