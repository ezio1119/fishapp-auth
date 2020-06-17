package main

import (
	"net"
	"time"

	"github.com/ezio1119/fishapp-profile/conf"
	"github.com/ezio1119/fishapp-profile/infrastructure"
	"github.com/ezio1119/fishapp-profile/infrastructure/middleware"
	"github.com/ezio1119/fishapp-profile/registry"
)

func main() {
	dbConn, err := infrastructure.NewGormConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	t := time.Duration(conf.C.Sv.Timeout) * time.Second
	registry := registry.NewRegistry(dbConn, t)
	profileController := registry.NewProfileController()

	middLe := middleware.InitMiddleware()
	server := infrastructure.NewGrpcServer(middLe, profileController)

	list, err := net.Listen("tcp", ":"+conf.C.Sv.Port)
	if err != nil {
		panic(err)
	}

	if err = server.Serve(list); err != nil {
		panic(err)
	}
}
