package usercomposer

import (
	goservice "github.com/lequocbinh04/go-sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"nckh-BE/appCommon"
	userstorage "nckh-BE/module/user/storage"
	userrpctransport "nckh-BE/module/user/transport/rpc"
	userproto "nckh-BE/proto/user"
)

func GetUserByIdServer(sc goservice.ServiceContext) userproto.UserServiceServer {
	db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
	dbStore := userstorage.NewMgDBStore(db)
	userBiz := userrpctransport.NewUserFindByIdBiz(dbStore)
	return userBiz
}
