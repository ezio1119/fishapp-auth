package infrastructure

import (
	"log"

	"github.com/go-redis/redis/v7"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "auth-kvs:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	return client
}
