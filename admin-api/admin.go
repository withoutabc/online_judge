package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_judge/dao"
	"online_judge/util"
)

func main() {
	r := gin.Default()
	dao.InitAdapter()
	e := dao.InitEnforcer()
	r.POST("/admin/:user_id", func(c *gin.Context) {
		userId := c.Param("user_id")
		b, err := e.AddGroupingPolicy(userId, "admin_group")
		if err != nil {
			util.RespInternalErr(c)
			return
		}
		if b != true {
			c.JSON(http.StatusBadRequest, gin.H{"info": "已经是管理员"})
			return
		}
		util.RespOK(c)
	})
	r.Run(":2334")
}
