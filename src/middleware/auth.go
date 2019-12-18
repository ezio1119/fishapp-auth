package middleware

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	update  string = "/user_grpc.UserService/Update"
	refresh string = "/user_grpc.UserService/RefreshIDToken"
	delete  string = "/user_grpc.UserService/Delete"
)

func (m *GoMiddleware) authFunc(ctx context.Context) (context.Context, error) {
	t, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	userID, err := m.parseToken(t)
	if err != nil {
		return nil, err
	}
	// grpc_ctxtags.Extract(ctx).Set("auth.sub", userID)
	newCtx := context.WithValue(ctx, "userID", userID)
	return newCtx, nil
}

func (m *GoMiddleware) AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var err error
		method := info.FullMethod

		if method == update || method == refresh || method == delete {
			ctx, err = m.authFunc(ctx)
		}
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		return handler(ctx, req)
	}
}
