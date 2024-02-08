package imagebiz

import (
	"context"
	"fmt"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson"
	"nckh-BE/appCommon"
	imagemodel "nckh-BE/module/image/model"
	"time"
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

func (biz *listBiz) ListDataWithCondition(ctx context.Context, filter *imagemodel.ImageList, paging *appCommon.Paging, moreInfo ...string) ([]imagemodel.Image, error) {
	filter.FulFill()
	paging.Fulfill()

	timeFrom := time.Unix(*filter.TimeFrom, 0)
	timeTo := time.Unix(*filter.TimeTo, 0)

	fmt.Println(timeFrom, timeTo)

	conditions := bson.M{
		"created_at": bson.M{
			"$gte": timeFrom,
			"$lte": timeTo,
		},
	}
	result, err := biz.store.ListDataWithCondition(ctx, conditions, paging, moreInfo...)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return []imagemodel.Image{}, appCommon.ErrCannotListEntity(imagemodel.EntityName, err)
	}
	return result, nil
}
