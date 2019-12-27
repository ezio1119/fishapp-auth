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

func NewGrpcServer(middLe *middleware.Middleware, uc *controllers.UserController) *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middLe.LoggerInterceptor(),
			grpc_validator.UnaryServerInterceptor(),
			middLe.RecoveryInterceptor(),
		)),
	)
	user_grpc.RegisterUserServiceServer(server, uc)
	reflection.Register(server)
	return server
}
