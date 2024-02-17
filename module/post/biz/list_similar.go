package postbiz

import (
	"context"
	"fmt"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nckh-BE/appCommon"
	postmodel "nckh-BE/module/post/model"
)

type listSimilarPostStore interface {
	ListRandom(ctx context.Context, condition bson.M, paging *appCommon.Paging, moreInfo ...string) ([]postmodel.SimplePost, error)
	Find(ctx context.Context, conditions interface{}, fields ...string) (*postmodel.Post, error)
}

type listSimilarPostBiz struct {
	store  listSimilarPostStore
	logger logger.Logger
}

func NewListSimilarPostBiz(store listSimilarPostStore) *listSimilarPostBiz {
	return &listSimilarPostBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("ListSimilarPostBiz"),
	}
}

func (biz *listSimilarPostBiz) ListSimilarPost(ctx context.Context, filter *postmodel.PostSimilarList) ([]postmodel.SimplePost, error) {
	postId, err := primitive.ObjectIDFromHex(filter.PostId)
	if err != nil {
		return []postmodel.SimplePost{}, appCommon.ErrInvalidRequest(err)
	}
	post, err := biz.store.Find(ctx, bson.M{
		"_id": postId,
	})
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return []postmodel.SimplePost{}, appCommon.ErrEntityNotFound(postmodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return []postmodel.SimplePost{}, appCommon.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	if filter.Type != "tag" && filter.Type != "category" {
		return []postmodel.SimplePost{}, appCommon.ErrInvalidRequest(fmt.Errorf("type is invalid"))
	}

	conditions := bson.M{}
	if filter.Type == "tag" {
		conditions["tag"] = bson.M{"$in": post.Tag}
	} else {
		conditions["category"] = bson.M{"$in": post.Category}
	}

	res, err := biz.store.ListRandom(ctx, conditions, &appCommon.Paging{
		Page:  1,
		Limit: filter.Limit,
	})
	if err != nil {
		return []postmodel.SimplePost{}, appCommon.ErrCannotListEntity(postmodel.EntityName, err)
	}

	return res, nil
}
