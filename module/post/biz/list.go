package postbiz

import (
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson"
	"nckh-BE/appCommon"
	postmodel "nckh-BE/module/post/model"
)

type PostListStore interface {
	ListDataWithCondition(ctx context.Context, condition bson.M, paging *appCommon.Paging, moreInfo ...string) ([]postmodel.Post, error)
}

type PostListBiz struct {
	store  PostListStore
	logger logger.Logger
}

func NewPostListBiz(store PostListStore) *PostListBiz {
	return &PostListBiz{store: store, logger: logger.GetCurrent().GetLogger("PostListBiz")}
}

func (biz *PostListBiz) ListDataWithCondition(ctx context.Context, paging *appCommon.Paging, filter *postmodel.PostList) ([]postmodel.Post, error) {
	paging.Fulfill()

	condition := bson.M{}

	if filter.Tag != nil {
		condition["tag"] = *filter.Tag
	}
	if filter.Category != nil {
		condition["category"] = *filter.Category
	}
	res, err := biz.store.ListDataWithCondition(ctx, condition, paging)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return []postmodel.Post{}, appCommon.ErrCannotListEntity(postmodel.EntityName, err)
	}

	if len(res) > 0 {
		paging.NextCursor = res[len(res)-1].Id.Hex()
	}
	return res, nil
}
