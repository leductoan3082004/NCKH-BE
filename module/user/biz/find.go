package userbiz

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nckh-BE/appCommon"
	usermodel "nckh-BE/module/user/model"
)

type userFindStore interface {
	Find(ctx context.Context, conditions interface{}, fields ...string) (*usermodel.User, error)
}

type userFindBiz struct {
	store  userFindStore
	logger logger.Logger
}

func NewFindBiz(store userFindStore) *userFindBiz {
	return &userFindBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("UserFindBiz"),
	}
}

func (biz *userFindBiz) Find(ctx context.Context, id string) (*usermodel.User, error) {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	data, err := biz.store.Find(ctx, bson.M{
		"_id": userId,
	})
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return nil, appCommon.ErrEntityNotFound(usermodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
	}
	return data, nil
}
