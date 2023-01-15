package main

import (
	"online_judge/api"
	"online_judge/dao"
	"online_judge/service"
)

func main() {
	dao.InitDB()
	go service.Judge()
	api.InitRouter()
}
