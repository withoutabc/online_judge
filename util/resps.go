package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespTemplate struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

func RespOK(c *gin.Context) {
	c.JSON(http.StatusOK, RespTemplate{
		Status: 200,
		Info:   "success",
	})
}

var ParamError = RespTemplate{
	Status: 400,
	Info:   "params error",
}

func RespParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError)
}

var InternalErr = RespTemplate{
	Status: 500,
	Info:   "internal error",
}

func RespInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalErr)
}

type NormSuccess struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   any    `json:"data"`
}

func RespNormSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, NormSuccess{Status: 200,
		Info: "success",
		Data: data,
	})
}
