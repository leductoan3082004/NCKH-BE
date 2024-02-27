package userrpcclient

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nckh-BE/appCommon"
	usermodel "nckh-BE/module/user/model"
	userproto "nckh-BE/proto/user"
)

func (s *rpcStore) GetUser(ctx context.Context, userId string) (*usermodel.User, error) {
	res, err := s.store.GetUser(ctx, &userproto.UserRequest{Id: userId})
	if err != nil {
		return nil, appCommon.ErrInternal(err)
	}

	id, _ := primitive.ObjectIDFromHex(res.GetId())
	return &usermodel.User{
		MgDBModel: appCommon.MgDBModel{
			Id:        id,
			Status:    int(res.GetStatus()),
			CreatedAt: res.CreatedAt.AsTime(),
			UpdatedAt: res.UpdatedAt.AsTime(),
		},
		Name:    res.GetName(),
		IsAdmin: res.GetIsAdmin(),
	}, nil
}
