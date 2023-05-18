package main

import "fmt"

func main() {
	var str string
	var i int
	fmt.Scan(&str)
	s := []rune(str) //汉字转换为字节
	for i = 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			break //第一次不同时终止i的递增
		}
	}
	if i <= len(s)/2-1 {
		fmt.Println(false)
		return
	}
	fmt.Println(true)
}
