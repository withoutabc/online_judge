package main

import (
	"online_judge/dao"
	"online_judge/exert/judge"
	"online_judge/redis"
	"time"
)

func main() {
	dao.InitDB()
	redis.InitRedis()
	judge.InitRabbitMq()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			judge.Produce()
		}
	}
}
