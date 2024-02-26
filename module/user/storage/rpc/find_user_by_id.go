package userrpc

import (
	"context"
	"nckh-BE/appCommon"
	usermodel "nckh-BE/module/user/model"
	userproto "nckh-BE/proto/user"
)

func (s *rpcStore) GetUser(ctx context.Context, userId string) (*usermodel.User, error) {
	res, err := s.store.GetUser(ctx, &userproto.UserRequest{Id: userId})
	if err != nil {
		return nil, appCommon.ErrInternal(err)
	}

	return &usermodel.User{
		MgDBModel: appCommon.MgDBModel{
			Status:    int(res.GetStatus()),
			CreatedAt: res.CreatedAt.AsTime(),
			UpdatedAt: res.UpdatedAt.AsTime(),
		},
		Name:    res.GetName(),
		IsAdmin: res.GetIsAdmin(),
	}, nil
}
