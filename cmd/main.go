package main

import (
	"online_judge/api"
	"online_judge/dao"
	"online_judge/service"
	"time"
)

func main() {
	dao.InitDB()
	go func() {
		ticker := time.NewTicker(time.Minute * 2)
		for range ticker.C {
			service.Judge()
		}
	}()
	api.InitRouter()
}
