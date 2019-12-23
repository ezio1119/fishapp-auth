package main

import (
	"log"
	"net"
	"time"

	"github.com/ezio1119/fishapp-user/conf"
	"github.com/ezio1119/fishapp-user/middleware"
	_userGrpcDeliver "github.com/ezio1119/fishapp-user/user/delivery/grpc"
	_userRepo "github.com/ezio1119/fishapp-user/user/repository"
	_userUcase "github.com/ezio1119/fishapp-user/user/usecase"
	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func main() {
	conf.Readconf()
	CONNECT := conf.C.Db.User + ":" + conf.C.Db.Pass + "@(" + conf.C.Db.Host + ":" + conf.C.Db.Port + ")/" + conf.C.Db.Name + "?" + conf.C.Db.ConnOpt
	dbConn, err := gorm.Open(conf.C.Db.Dbms, CONNECT)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	timeoutContext := time.Duration(conf.C.Sv.Timeout) * time.Second
	userRepo := _userRepo.NewMysqlUserRepository(dbConn)
	userUcase := _userUcase.NewUserUsecase(userRepo, timeoutContext)
	middL := middleware.InitMiddleware()

	gserver := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middL.LoggerInterceptor(),
			grpc_validator.UnaryServerInterceptor(),
			middL.RecoveryInterceptor(),
		)),
	)
	_userGrpcDeliver.NewUserServerGrpc(gserver, userUcase)

	list, err := net.Listen("tcp", ":"+conf.C.Sv.Port)
	if err != nil {
		log.Fatal(err)
	}

	err = gserver.Serve(list)
	if err != nil {
		log.Fatal(err)
	}
}
