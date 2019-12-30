package controllers

import (
	"context"
	"fmt"

	"github.com/ezio1119/fishapp-auth/domain"
	"github.com/ezio1119/fishapp-auth/interfaces/controllers/user_grpc"
	"github.com/ezio1119/fishapp-auth/usecase/interactor"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type UserController struct {
	UserInteractor interactor.UUserInteractor
}

type contextKey string

const UserIDCtxKey contextKey = "userID"

func getUserIDCtx(ctx context.Context) (int64, error) {
	v := ctx.Value(UserIDCtxKey)
	userID, ok := v.(int64)
	if !ok {
		return 0, fmt.Errorf("userID not found")
	}

	return userID, nil
}

func (c *UserController) transformUserRPC(u *domain.User) (*user_grpc.User, error) {
	updatedAt, err := ptypes.TimestampProto(u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.TimestampProto(u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user_grpc.User{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}, nil
}

func (c *UserController) transformTokenPairRPC(tp *domain.TokenPair) *user_grpc.TokenPair {
	return &user_grpc.TokenPair{
		IdToken:      tp.IDToken,
		RefreshToken: tp.RefreshToken,
	}
}

func (c *UserController) Create(ctx context.Context, in *user_grpc.CreateReq) (*user_grpc.UserWithToken, error) {
	user := &domain.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
	tokenPair, err := c.UserInteractor.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	tokenPairRPC := c.transformTokenPairRPC(tokenPair)
	userRPC, err := c.transformUserRPC(user)
	if err != nil {
		return nil, err
	}
	return &user_grpc.UserWithToken{
		User:      userRPC,
		TokenPair: tokenPairRPC,
	}, nil
}

func (c *UserController) GetByID(ctx context.Context, in *user_grpc.ID) (*user_grpc.User, error) {
	user, err := c.UserInteractor.GetByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	userRPC, err := c.transformUserRPC(user)
	if err != nil {
		return nil, err
	}
	return userRPC, nil
}

func (c *UserController) Update(ctx context.Context, in *user_grpc.UpdateReq) (*user_grpc.User, error) {
	user := &domain.User{
		ID:       in.Id,
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
	if err := c.UserInteractor.Update(ctx, user); err != nil {
		return nil, err
	}
	userRPC, err := c.transformUserRPC(user)
	if err != nil {
		return nil, err
	}
	return userRPC, nil
}

func (c *UserController) Delete(ctx context.Context, in *user_grpc.ID) (*wrappers.BoolValue, error) {
	if err := c.UserInteractor.Delete(ctx, in.Id); err != nil {
		return nil, err
	}
	return &wrappers.BoolValue{
		Value: false,
	}, nil
}

func (c *UserController) Login(ctx context.Context, in *user_grpc.LoginReq) (*user_grpc.UserWithToken, error) {
	user, tokenPair, err := c.UserInteractor.Login(ctx, in.Email, in.Password)
	if err != nil {
		return nil, err
	}
	userRPC, err := c.transformUserRPC(user)
	tokenPairRPC := c.transformTokenPairRPC(tokenPair)
	return &user_grpc.UserWithToken{
		User:      userRPC,
		TokenPair: tokenPairRPC,
	}, nil
}

func (c *UserController) AddBlackList(ctx context.Context, in *user_grpc.AddBlackListReq) (*wrappers.BoolValue, error) {
	exp, err := ptypes.Duration(in.Expiration)
	if err != nil {
		return nil, err
	}
	success, err := c.UserInteractor.AddBlackList(ctx, in.Jti, exp)
	if err != nil {
		return nil, err
	}
	return &wrappers.BoolValue{
		Value: success,
	}, nil
}

func (c *UserController) CheckBlackListAndGenToken(ctx context.Context, in *user_grpc.CheckBlackListReq) (*user_grpc.TokenPair, error) {
	tokenPair, err := c.UserInteractor.CheckBlackListAndGenToken(ctx, in.Id, in.Jti)
	if err != nil {
		return nil, err
	}
	return c.transformTokenPairRPC(tokenPair), nil
}
