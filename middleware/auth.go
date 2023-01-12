package middleware

import (
	"github.com/gin-gonic/gin"
	"online_judge/util"
)

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
