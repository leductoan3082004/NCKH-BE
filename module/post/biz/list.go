package postbiz

import (
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson"
	"nckh-BE/appCommon"
	postmodel "nckh-BE/module/post/model"
)

type PostListStore interface {
	ListDataWithCondition(ctx context.Context, condition bson.M, paging *appCommon.Paging, moreInfo ...string) ([]postmodel.SimplePost, error)
}

type PostListBiz struct {
	store  PostListStore
	logger logger.Logger
}

func NewPostListBiz(store PostListStore) *PostListBiz {
	return &PostListBiz{store: store, logger: logger.GetCurrent().GetLogger("PostListBiz")}
}

func (biz *PostListBiz) ListDataWithCondition(ctx context.Context, paging *appCommon.Paging, filter *postmodel.PostList) ([]postmodel.SimplePost, error) {
	paging.Fulfill()

	condition := bson.M{}

	if filter.Tag != nil {
		condition["tag"] = *filter.Tag
	}
	if filter.Category != nil {
		condition["category"] = *filter.Category
	}
	if filter.Content != nil {
		condition["$text"] = bson.M{
			"$search": *filter.Content,
		}
	}
	res, err := biz.store.ListDataWithCondition(ctx, condition, paging)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return []postmodel.SimplePost{}, appCommon.ErrCannotListEntity(postmodel.EntityName, err)
	}

	if res == nil {
		return []postmodel.SimplePost{}, nil
	}
	if len(res) > 0 {
		paging.NextCursor = res[len(res)-1].Id.Hex()
	}
	return res, nil
}
