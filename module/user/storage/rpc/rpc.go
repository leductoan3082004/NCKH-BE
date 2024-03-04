package userrpcclient

import (
	"nckh-BE/proto/user"
)

type rpcStore struct {
	store userproto.UserServiceClient
}

func NewRpcStore(store userproto.UserServiceClient) *rpcStore {
	return &rpcStore{store: store}
}
