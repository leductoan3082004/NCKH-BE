package userrpctransport

import (
	"context"
	goservice "github.com/lequocbinh04/go-sdk"
	"github.com/lequocbinh04/go-sdk/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
	"nckh-BE/appCommon"
	usermodel "nckh-BE/module/user/model"
	userstorage "nckh-BE/module/user/storage"
	userproto "nckh-BE/proto/user"
)

type userFindByIdStore interface {
	GetUser(ctx context.Context, userId string) (*usermodel.User, error)
}
type userFindByIdBiz struct {
	store  userFindByIdStore
	logger logger.Logger
}

func newUserFindBiz(store userFindByIdStore) *userFindByIdBiz {
	return &userFindByIdBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("UserFindByIdBizGRPC"),
	}
}

func (s *userFindByIdBiz) GetUser(ctx context.Context, request *userproto.UserRequest) (*userproto.UserResponse, error) {
	user, err := s.store.GetUser(ctx, request.Id)
	if err != nil {
		s.logger.WithSrc().Error(err)
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

func GetUserByIdServer(sc goservice.ServiceContext) userproto.UserServiceServer {
	db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
	dbStore := userstorage.NewMgDBStore(db)
	userBiz := newUserFindBiz(dbStore)
	return userBiz
}
