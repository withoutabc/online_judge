package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"online_judge/model"
	"online_judge/util"
)

var DB *gorm.DB

//func InitDB() {
//	db, err := sql.Open("mysql", "root:224488@tcp(127.0.0.1:3306)/online_judge?charset=utf8mb4&loc=Local&parseTime=true")
//	if err != nil {
//		logs.Fatalf("connect mysql error:%v", err)
//	}
//	DB = db
//	fmt.Println(db.Ping())
//}

//debian-sys-maint:ZF0kfsp5uMD2lVo7

// InitDB gorm连接
func InitDB() {
	dsn := "root:224488@tcp(127.0.0.1:3306)/online_judge?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		util.Log(err)
		panic("failed to connect database")
	}
	DB = db
	log.Println("连接成功")
	AutoMigrate()
}

func AutoMigrate() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Problem{})
	DB.AutoMigrate(&model.Submission{})
	DB.AutoMigrate(&model.Testcase{})
	DB.AutoMigrate(&model.Info{})
}

func GetDB() *gorm.DB {
	return DB
}
