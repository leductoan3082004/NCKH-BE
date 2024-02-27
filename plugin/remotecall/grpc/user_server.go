package userplugin

import "flag"

type userGrpcServer struct {
	url    string
	prefix string
	name   string
}

func NewUserGrpcServer(name string, prefix string) *userGrpcServer {
	return &userGrpcServer{
		prefix: prefix,
		name:   name,
	}
}

func (t *userGrpcServer) GetPrefix() string {
	return t.prefix
}

func (t *userGrpcServer) Get() interface{} {
	//TODO implement me
	panic("implement me")
}

func (t *userGrpcServer) Name() string {
	return t.name
}

func (t *userGrpcServer) InitFlags() {
	flag.StringVar(&t.url, t.prefix+"-url", "0.0.0.0:50051", "User gRPC url")
}

func (t *userGrpcServer) Configure() error {
	//TODO implement me
	panic("implement me")
}

func (t *userGrpcServer) Run() error {
	return t.Configure()
}

func (t *userGrpcServer) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}
