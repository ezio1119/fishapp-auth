package infrastructure

import (
	"github.com/ezio1119/fishapp-auth/infrastructure/middleware"
	"github.com/ezio1119/fishapp-auth/interfaces/controllers"
	"github.com/ezio1119/fishapp-auth/interfaces/controllers/user_grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(c *controllers.UserController) *grpc.Server {
	middle := middleware.InitMiddleware()
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middle.LoggerInterceptor(),
			grpc_validator.UnaryServerInterceptor(),
			middle.RecoveryInterceptor(),
		)),
	)
	user_grpc.RegisterUserServiceServer(server, c)
	reflection.Register(server)
	return server
}
