package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes"

	"github.com/ezio1119/fishapp-user/models"
	"github.com/ezio1119/fishapp-user/user"
	"github.com/ezio1119/fishapp-user/user/delivery/grpc/user_grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	UUsecase user.Usecase
}

func NewUserServerGrpc(gserver *grpc.Server, us user.Usecase) {
	userServer := &server{
		UUsecase: us,
	}
	user_grpc.RegisterUserServiceServer(gserver, userServer)
	reflection.Register(gserver)
}

func (s *server) transformUserRPC(u *models.User) (*user_grpc.User, error) {
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

func (s *server) transformTokenPairRPC(tp *models.TokenPair) *user_grpc.TokenPair {
	return &user_grpc.TokenPair{
		IdToken:      tp.IDToken,
		RefreshToken: tp.RefreshToken,
	}
}

func (s *server) Create(ctx context.Context, in *user_grpc.CreateReq) (*user_grpc.UserWithToken, error) {
	user := &models.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
	tokenPair, err := s.UUsecase.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	tokenPairRPC := s.transformTokenPairRPC(tokenPair)
	userRPC, err := s.transformUserRPC(user)
	if err != nil {
		return nil, err
	}
	return &user_grpc.UserWithToken{
		User:      userRPC,
		TokenPair: tokenPairRPC,
	}, nil
}

func (s *server) GetByID(ctx context.Context, in *user_grpc.ID) (*user_grpc.User, error) {
	user, err := s.UUsecase.GetByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	userRPC, err := s.transformUserRPC(user)
	if err != nil {
		return nil, err
	}
	return userRPC, nil
}

func (s *server) Update(ctx context.Context, in *user_grpc.UpdateReq) (*user_grpc.User, error) {
	user := &models.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
	userID := ctx.Value("userID").(int64)
	user.ID = userID
	if err := s.UUsecase.Update(ctx, user); err != nil {
		return nil, err
	}
	userRPC, err := s.transformUserRPC(user)
	if err != nil {
		return nil, err
	}
	return userRPC, nil
}

func (s *server) Delete(ctx context.Context, in *empty.Empty) (*user_grpc.DeleteRes, error) {
	userID := ctx.Value("userID").(int64)
	if err := s.UUsecase.Delete(ctx, userID); err != nil {
		return nil, err
	}
	return &user_grpc.DeleteRes{
		Deleted: true,
	}, nil
}

func (s *server) RefreshIDToken(ctx context.Context, in *empty.Empty) (*user_grpc.TokenPair, error) {
	userID := ctx.Value("userID").(int64)
	tokenPair, err := s.UUsecase.RefreshIDToken(ctx, userID)
	if err != nil {
		return nil, err
	}
	tokenPairRPC := s.transformTokenPairRPC(tokenPair)
	return tokenPairRPC, nil
}

func (s *server) Login(ctx context.Context, in *user_grpc.LoginReq) (*user_grpc.UserWithToken, error) {
	user, tokenPair, err := s.UUsecase.Login(ctx, in.Email, in.Password)
	if err != nil {
		return nil, err
	}
	userRPC, err := s.transformUserRPC(user)
	tokenPairRPC := s.transformTokenPairRPC(tokenPair)
	return &user_grpc.UserWithToken{
		User:      userRPC,
		TokenPair: tokenPairRPC,
	}, nil
}
