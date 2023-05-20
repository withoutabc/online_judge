package middleware

import (
	"github.com/gin-gonic/gin"
	"online_judge/util"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			util.NormErr(c, util.BlankAuthErrCode)
			c.Abort()
			return
		}
		//按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			util.NormErr(c, util.WrongAuthFormatErrCode)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，使用解析JWT的函数来解析
		_, err := util.ParseToken(parts[1])
		if err != nil {
			util.NormErr(c, util.InValidTokenErrCode)
			c.Abort()
			return
		}
		c.Next()
	}
}
