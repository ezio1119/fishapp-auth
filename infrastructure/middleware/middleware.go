package middleware

import (
	"google.golang.org/grpc"
)

type middleware struct{}

type Middleware interface {
	RecoveryInterceptor() grpc.UnaryServerInterceptor
	LoggerInterceptor() grpc.UnaryServerInterceptor
	ValidatorInterceptor() grpc.UnaryServerInterceptor
}

func InitMiddleware() Middleware {
	return &middleware{}
}
