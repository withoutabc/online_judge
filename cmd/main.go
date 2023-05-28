package main

import (
	"online_judge/api"
	"online_judge/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
