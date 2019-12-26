package repository

import (
	"github.com/go-redis/redis/v7"
)

type AuthRepository struct {
	Client *redis.Client
}
