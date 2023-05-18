package middleware

//
//import (
//	"net/http"
//	"online_judge/service"
//	"strings"
//
//	"github.com/gin-gonic/gin"
//)
//
//// JWTAuthMiddleware 基于JWT的认证中间件
//func JWTAuthMiddleware() func(c *gin.Context) {
//	return func(c *gin.Context) {
//		//Token放在Header的Authorization中，并使用Bearer开头
//		authHeader := c.Request.Header.Get("Authorization")
//		if authHeader == "" {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"status": 2003,
//				"info":   "请求头中auth为空",
//			})
//			c.Abort()
//			return
//		}
//		//按空格分割
//		parts := strings.SplitN(authHeader, " ", 2)
//		if !(len(parts) == 2 && parts[0] == "Bearer") {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"status": 2004,
//				"info":   "请求头中auth格式有误",
//			})
//			c.Abort()
//			return
//		}
//		// parts[1]是获取到的tokenString，使用解析JWT的函数来解析
//		mc, err := service.ParseToken(parts[1])
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"status": 2005,
//				"info":   "无效的Token",
//			})
//			c.Abort()
//			return
//		}
//		//这步全文都没有Get过
//		c.Set("uid", mc.Uid)
//		c.Next()
//	}
//}
