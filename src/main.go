package main

import (
	"log"
	"net"
	"time"

	"github.com/ezio1119/fishapp-user/middleware"
	_userGrpcDeliver "github.com/ezio1119/fishapp-user/user/delivery/grpc"
	_userRepo "github.com/ezio1119/fishapp-user/user/repository"
	_userUcase "github.com/ezio1119/fishapp-user/user/usecase"
	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

type env struct {
	DbPass     string `required:"true" split_words:"true"`
	DbDbms     string `required:"true" split_words:"true"`
	DbUser     string `required:"true" split_words:"true"`
	DbName     string `required:"true" split_words:"true"`
	DbPort     string `required:"true" split_words:"true"`
	DbHost     string `required:"true" split_words:"true"`
	DbConnOpt  string `required:"true" split_words:"true"`
	Timeout    int64  `required:"true"`
	ListenPort string `required:"true" split_words:"true"`
	Debug      bool   `required:"true"`
}

func main() {
	var env env
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatal(err)
	}
	CONNECT := env.DbUser + ":" + env.DbPass + "@(" + env.DbHost + ":" + env.DbPort + ")/" + env.DbName + "?" + env.DbConnOpt
	dbConn, err := gorm.Open(env.DbDbms, CONNECT)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	timeoutContext := time.Duration(env.Timeout) * time.Second
	userRepo := _userRepo.NewMysqlUserRepository(dbConn)
	userUcase := _userUcase.NewUserUsecase(userRepo, timeoutContext)
	middL := middleware.InitMiddleware()

	gserver := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middL.LoggerInterceptor(env.Debug),
			middL.AuthInterceptor(),
			grpc_validator.UnaryServerInterceptor(),
			middL.RecoveryInterceptor(),
		)),
	)
	_userGrpcDeliver.NewUserServerGrpc(gserver, userUcase)

	list, err := net.Listen("tcp", ":"+env.ListenPort)
	if err != nil {
		log.Fatal(err)
	}

	err = gserver.Serve(list)
	if err != nil {
		log.Fatal(err)
	}
}
