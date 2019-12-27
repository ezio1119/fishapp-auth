package infrastructure

import (
	"log"

	"github.com/ezio1119/fishapp-auth/conf"
	"github.com/go-redis/redis/v7"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.C.Kvs.Host + ":" + conf.C.Kvs.Port,
		Password: conf.C.Kvs.Pass,
		DB:       conf.C.Kvs.Db,
		Network:  conf.C.Kvs.Net,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	return client
}
