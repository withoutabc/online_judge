package middleware

import (
	"github.com/gin-gonic/gin"
	"online_judge/dao"
	"online_judge/util"
)

func AdminAuthority() func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, b := c.Get("user_id")
		if b != true {
			util.RespUnauthorized(c)
			c.Abort()
			return
		}
		dao.InitAdapter()
		e := dao.InitEnforcer()
		b, err := e.Enforce(userId, "admin_data", "read")
		if err != nil || b != true {
			util.RespUnauthorized(c)
			c.Abort()
			return
		}
	}
}
