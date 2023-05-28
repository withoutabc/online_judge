package main

import (
	"online_judge/dao"
	"online_judge/exert/judge"
	"time"

	"online_judge/redis"
)

func main() {
	dao.InitDB()
	redis.InitRedis()
	judge.InitRabbitMq()

	for {
		time.Sleep(2 * time.Second)
		judge.Produce()
	}
}
