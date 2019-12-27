package main

import (
	"log"
	"net"
	"time"

	"github.com/ezio1119/fishapp-auth/conf"
	"github.com/ezio1119/fishapp-auth/infrastructure"
	"github.com/ezio1119/fishapp-auth/infrastructure/middleware"
	"github.com/ezio1119/fishapp-auth/registry"
)

func main() {
	dbConn := infrastructure.NewGormConn()
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	t := time.Duration(conf.C.Sv.Timeout) * time.Second
	redisClient := infrastructure.NewRedisClient()
	userController := registry.NewUserController(t, dbConn, redisClient)

	middLe := middleware.InitMiddleware()
	server := infrastructure.NewGrpcServer(middLe, userController)
	list, err := net.Listen("tcp", ":"+conf.C.Sv.Port)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Serve(list)
	if err != nil {
		log.Fatal(err)
	}
}
