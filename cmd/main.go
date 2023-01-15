package main

import (
	"online_judge/api"
	"online_judge/controller"
	"online_judge/dao"
)

func main() {
	dao.InitDB()
	go controller.Judge()
	api.InitRouter()
}
