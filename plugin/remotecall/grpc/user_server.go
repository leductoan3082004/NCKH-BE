package userplugin

import (
	"flag"
	"fmt"
	"github.com/lequocbinh04/go-sdk/logger"
	"google.golang.org/grpc"
	userproto "nckh-BE/proto/user"
	"net"
)

type userGrpcServer struct {
	port     string
	prefix   string
	name     string
	server   *grpc.Server
	logger   logger.Logger
	handlers []userproto.UserServiceServer
}

func NewUserGrpcServer(name string, prefix string, handlers ...userproto.UserServiceServer) *userGrpcServer {
	return &userGrpcServer{
		prefix:   prefix,
		name:     name,
		handlers: handlers,
	}
}

func (t *userGrpcServer) GetPrefix() string {
	return t.prefix
}

func (t *userGrpcServer) Get() interface{} {
	return t
}

func (t *userGrpcServer) Name() string {
	return t.name
}

func (t *userGrpcServer) InitFlags() {
	flag.StringVar(&t.port, t.prefix+"-port", "50051", "User gRPC server serving port")
}

func (t *userGrpcServer) Configure() error {
	t.logger = logger.GetCurrent().GetLogger(t.name)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", t.port))

	if err != nil {
		t.logger.Fatalln(err)
	}

	t.logger.Infoln(fmt.Sprintf("GRPC Server is listening on %s ...\n", t.port))
	t.server = grpc.NewServer()

	for i := range t.handlers {
		userproto.RegisterUserServiceServer(t.server, t.handlers[i])
	}

	go func() {
		if err := t.server.Serve(lis); err != nil {
			t.logger.Fatalln(err)
		}
	}()

	return nil
}

func (t *userGrpcServer) Run() error {
	return t.Configure()
}

func (t *userGrpcServer) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}
