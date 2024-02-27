package userrpctransport

import (
	"context"
	"github.com/lequocbinh04/go-sdk/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
	usermodel "nckh-BE/module/user/model"
	userproto "nckh-BE/proto/user"
)

type userFindByIdStore interface {
	GetUser(ctx context.Context, userId string) (*usermodel.User, error)
}
type userFindByIdBiz struct {
	store  userFindByIdStore
	logger logger.Logger
}

func NewUserFindByIdBiz(store userFindByIdStore) *userFindByIdBiz {
	return &userFindByIdBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("UserFindByIdBizGRPC"),
	}
}

func (s *userFindByIdBiz) GetUser(ctx context.Context, request *userproto.UserRequest) (*userproto.UserResponse, error) {
	user, err := s.store.GetUser(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &userproto.UserResponse{
		Id:        user.Id.Hex(),
		Name:      user.Name,
		IsAdmin:   user.IsAdmin,
		Status:    int32(user.Status),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}
