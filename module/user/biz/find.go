package userbiz

import (
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"nckh-BE/appCommon"
	usermodel "nckh-BE/module/user/model"
)

type userFindStore interface {
	GetUser(ctx context.Context, userId string) (*usermodel.User, error)
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
	data, err := biz.store.GetUser(ctx, id)
	if err != nil {
		if err == appCommon.ErrRecordNotFound {
			return nil, appCommon.ErrEntityNotFound(usermodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
	}
	return data, nil
}
