package infrastructure

import (
	"github.com/ezio1119/fishapp-profile/infrastructure/middleware"
	"github.com/ezio1119/fishapp-profile/pb"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(middLe middleware.Middleware, profileController pb.ProfileServiceServer) *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middLe.LoggerInterceptor(),
			middLe.ValidatorInterceptor(),
			middLe.RecoveryInterceptor(),
		)),
	)
	pb.RegisterProfileServiceServer(server, profileController)
	reflection.Register(server)
	return server
}
