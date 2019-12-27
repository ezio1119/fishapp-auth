package repository

import (
	"github.com/go-redis/redis/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BlackListRepository struct {
	Client *redis.Client
}

func (r *BlackListRepository) SAdd(token string) error {
	result, err := r.Client.SAdd("token", token).Result()
	if err != nil {
		return err
	}
	if result == 0 {
		return status.Error(codes.Unknown, "The same token is registered")
	}
	return nil
}

func (r *BlackListRepository) SIsMember(t string) (bool, error) {
	return r.Client.SIsMember("token", t).Result()
}
