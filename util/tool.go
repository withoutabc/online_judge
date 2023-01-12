package util

import (
	"github.com/gin-gonic/gin"
)

func IfExist(c *gin.Context, a string) {
	if a == "" {
		RespParamErr(c)
		return
	}
	return
}
