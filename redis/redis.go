package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var RDB *redis.Client

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Print(err)
	}
	RDB = rdb
	log.Println("redis连接成功")
}
