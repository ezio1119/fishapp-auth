package repository

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

type BlackListRepository struct {
	Client *redis.Client
}

func (r *BlackListRepository) SetNX(token string, exp time.Duration) (bool, error) {
	fmt.Println(exp)
	return r.Client.SetNX(token, "", exp).Result()
}

func (r *BlackListRepository) Exists(t string) (int64, error) {
	return r.Client.Exists(t).Result()
}
