package dao

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	A *gormadapter.Adapter
	E *casbin.Enforcer
)

func InitAdapter() {
	a, err := gormadapter.NewAdapter("mysql", "debian-sys-maint:ZF0kfsp5uMD2lVo7@tcp(127.0.0.1:3306)/online_judge", true)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	A = a
}

func InitEnforcer() *casbin.Enforcer {
	e, err := casbin.NewEnforcer("./model.conf", A)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	E = e
	return e
}
