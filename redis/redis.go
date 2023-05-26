package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var RDB *redis.Client

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10,              // 设置连接池大小
		MinIdleConns: 1,               // 设置最小空闲连接数
		IdleTimeout:  5 * time.Minute, // 设置连接的空闲超时时间
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	RDB = rdb
	log.Println("redis连接成功")
}
