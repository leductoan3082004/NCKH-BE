package usergrpcclient

import (
	"flag"
	"github.com/lequocbinh04/go-sdk/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	userproto "nckh-BE/proto/user"
)

type userGRPC struct {
	url    string
	name   string
	prefix string
	logger logger.Logger
	client userproto.UserServiceClient
}

func NewUserGRPC(name string, prefix string) *userGRPC {
	return &userGRPC{
		name:   name,
		prefix: prefix,
	}
}

func (r *userGRPC) GetPrefix() string {
	return r.prefix
}

func (r *userGRPC) Get() interface{} {
	return r.client
}

func (r *userGRPC) Name() string {
	return r.name
}

func (r *userGRPC) InitFlags() {
	flag.StringVar(&r.url, r.prefix+"-url", "localhost:50051", "User gRPC url")
}

func (r *userGRPC) Configure() error {
	r.logger = logger.GetCurrent().GetLogger(r.name)
	r.logger.WithSrc().Infof("Connecting to %s", r.url)
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial(r.url, opts)
	if err != nil {
		r.logger.Fatalln(err)
	}
	r.logger.WithSrc().Infof("Connected to %s", r.url)
	r.client = userproto.NewUserServiceClient(clientConn)
	return nil
}

func (r *userGRPC) Run() error {
	return r.Configure()
}

func (r *userGRPC) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}
