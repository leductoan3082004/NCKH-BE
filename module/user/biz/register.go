package userbiz

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/lequocbinh04/go-sdk/logger"
	"nckh-BE/appCommon"
	usermodel "nckh-BE/module/user/model"
)

type userRegisterStore interface {
	Create(ctx context.Context, data *usermodel.User) error
	Find(ctx context.Context, conditions interface{}, fields ...string) (*usermodel.User, error)
}

type userRegisterBiz struct {
	store  userRegisterStore
	logger logger.Logger
}

func NewUserRegisterBiz(store userRegisterStore) *userRegisterBiz {
	return &userRegisterBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("userBizRegister"),
	}
}

func (biz *userRegisterBiz) Register(ctx context.Context, data *usermodel.UserRegister) error {

	_, err := biz.store.Find(ctx, bson.M{
		"username": data.Username,
	})
	if err != nil {
		if err != appCommon.ErrRecordNotFound {
			biz.logger.WithSrc().Errorln(err)
			return appCommon.ErrInternal(err)
		}
	} else {
		return usermodel.ErrUsernameExisted
	}

	salt := appCommon.GenSalt(30)
	password, err := appCommon.HMACEncode(data.Password, salt)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrInternal(err)
	}

	user := &usermodel.User{
		Username: data.Username,
		Password: password,
		Salt:     salt,
		Name:     data.Name,
		IsAdmin:  false,
	}

	if err := biz.store.Create(ctx, user); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
