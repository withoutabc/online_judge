package main

import (
	"time"
)

func main() {
	ticker := time.NewTicker(time.Minute * 5) // 每5秒执行一次
	defer ticker.Stop()                       // 结束ticker
	for {
		select {
		case <-ticker.C:
			UpdateRanking()
		}
	}
}

// UpdateRanking 从redis中读取数据，更新排名
func UpdateRanking() {

}
