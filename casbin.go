package main

import "online_judge/dao"

func main() {
	dao.InitAdapter()
	e := dao.InitEnforcer()
	// 使user组可以对数据1有写权限
	b, err := e.AddPolicy("user_group", "common_data", "write")

	if err != nil || b != true {
		panic(err)
	}

	b, err = e.AddPolicy("user_group", "common_data", "read")

	if err != nil || b != true {
		panic(err)
	}
	b, err = e.AddPolicy("admin_group", "admin_data", "write")

	if err != nil || b != true {
		panic(err)
	}

	b, err = e.AddPolicy("admin_group", "admin_data", "read")

	if err != nil || b != true {
		panic(err)
	}
}
