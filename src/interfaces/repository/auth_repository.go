package repository

import (
	"github.com/go-redis/redis/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthRepository struct {
	Client *redis.Client
}

func (r *AuthRepository) SAdd(token string) error {
	result, err := r.Client.SAdd("token", token).Result()
	if err != nil {
		return err
	}
	if result == 0 {
		return status.Error(codes.Unknown, "The same token is registered")
	}
	return nil
}

func (r *AuthRepository) SIsMember(t string) (bool, error) {
	return r.Client.SIsMember("token", t).Result()
}
