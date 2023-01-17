package middleware

import (
	"github.com/gin-gonic/gin"
	"online_judge/util"
)

// Auth 没有用 只是保存一下
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := c.Cookie("uid")
		if err != nil {
			util.RespUnauthorizedErr(c)
			c.Abort()
			return
		}
		if uid == "" {
			util.RespUnauthorizedErr(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
