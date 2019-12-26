package controllers

import (
	"context"

	"github.com/ezio1119/fishapp-auth/interfaces/controllers/auth_grpc"
	"github.com/ezio1119/fishapp-auth/usecase/interactor"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type AuthController struct {
	AuthInteractor interactor.UAuthInteractor
}

func (c *AuthController) RefreshIDToken(context.Context, *empty.Empty) (*auth_grpc.TokenPair, error) {
	panic("not implement")
}
func (c *AuthController) CheckBlackList(context.Context, *empty.Empty) (*wrappers.BoolValue, error) {
	panic("not implement")
}
func (c *AuthController) AddBlackList(context.Context, *empty.Empty) (*wrappers.BoolValue, error) {
	panic("not implement")
}
