package infrastructure

import (
	"github.com/ezio1119/fishapp-profile/controllers/profile_grpc"
	"github.com/ezio1119/fishapp-profile/infrastructure/middleware"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(middLe middleware.Middleware, profileController profile_grpc.ProfileServiceServer) *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middLe.LoggerInterceptor(),
			middLe.ValidatorInterceptor(),
			middLe.RecoveryInterceptor(),
		)),
	)
	profile_grpc.RegisterProfileServiceServer(server, profileController)
	reflection.Register(server)
	return server
}
