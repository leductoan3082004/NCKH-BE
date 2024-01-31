package postbiz

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nckh-BE/appCommon"
	postmodel "nckh-BE/module/post/model"
)

type updateStore interface {
	Update(ctx context.Context, conditions interface{}, data interface{}) error
}

type updateBiz struct {
	store  updateStore
	logger logger.Logger
}

func NewUpdateBiz(store updateStore) *updateBiz {
	return &updateBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("UpdatePostBiz"),
	}
}

func (biz *updateBiz) Update(ctx context.Context, data *postmodel.PostUpdate) error {
	postId, err := primitive.ObjectIDFromHex(data.PostId)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	update := bson.M{}
	if data.Author != nil {
		update["author"] = *data.Author
	}
	if data.Title != nil {
		update["title"] = *data.Title
	}
	if data.Content != nil {
		update["content"] = *data.Content
	}
	if data.ImageUrl != nil {
		update["image_url"] = *data.ImageUrl
	}
	if data.Tag != nil {
		update["tag"] = *data.Tag
	}
	if data.Category != nil {
		update["category"] = *data.Category
	}

	condition := bson.M{
		"_id": postId,
	}

	if err := biz.store.Update(ctx, condition, update); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(postmodel.EntityName, err)
	}

	return nil
}
