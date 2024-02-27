package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"github.com/lequocbinh04/go-sdk/plugin/aws"
	"github.com/lequocbinh04/go-sdk/plugin/storage/sdkmgo"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"nckh-BE/appCommon"
	"nckh-BE/cmd/handler"
	usercomposer "nckh-BE/grpc/user"
	"nckh-BE/plugin/appredis"
	usergrpcclient "nckh-BE/plugin/remotecall/grpc"
	jwtProvider "nckh-BE/plugin/tokenprovider/jwt"
	userproto "nckh-BE/proto/user"
	"net"
)

func newService() goservice.Service {

	service := goservice.New(
		goservice.WithName("mindzone"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkmgo.NewMongoDB("mongodb", appCommon.DBMain)),
		goservice.WithInitRunnable(appredis.NewRedisDB("redis", appCommon.PluginRedis)),
		goservice.WithInitRunnable(jwtProvider.NewJwtProvider("jwt", appCommon.PluginJWT)),
		goservice.WithInitRunnable(aws.New("aws", appCommon.PluginAWS)),
		goservice.WithInitRunnable(usergrpcclient.NewUserGRPC("user", appCommon.PluginUserClient)),
	)

	if err := service.Init(); err != nil {
		panic(err)
	}
	return service
}

func startGRPCService(sc goservice.ServiceContext) {
	logger := sc.Logger("grpc")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		logger.WithSrc().Fatalln(err)
	}
	logger.Infof("GRPC Server is listening on %d ...\n", 50051)
	s := grpc.NewServer()
	userproto.RegisterUserServiceServer(s, usercomposer.GetUserByIdServer(sc))
	if err := s.Serve(lis); err != nil {
		logger.WithSrc().Fatalln(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start a mindzone service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			handler.MainRoute(engine, service)
		})

		go startGRPCService(service)
		if err := service.Start(func() {}); err != nil {
			panic(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
