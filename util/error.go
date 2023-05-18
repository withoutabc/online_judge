package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ErrorCodeMap = map[int]error{

	UnauthorizedErrCode:  UnauthorizedErr,
	InternalServeErrCode: InternalServeErr,

	IdNotIntegral:           IdNotIntegralErr,
	NoRecordErrCode:         NoRecordErr,
	RepeatedUsernameErrCode: RepeatedUsernameErr,
	WrongPasswordErrCode:    WrongPasswordErr,

	BindingQueryErrCode: BindingQueryErr,
	WrongTimeCode:       WrongTimeCodeErr,
	WrongLevelCode:      WrongLevelCodeErr,
}

var (
	UnauthorizedErr  = errors.New("unauthorized")
	InternalServeErr = errors.New("internal serve error")

	IdNotIntegralErr    = errors.New("id is not a integral")
	NoRecordErr         = errors.New("no record")
	RepeatedUsernameErr = errors.New("repeated username")
	WrongPasswordErr    = errors.New("wrong password")

	BindingQueryErr   = errors.New("binding error")
	WrongTimeCodeErr  = errors.New("wrong time")
	WrongLevelCodeErr = errors.New("wrong level")
)

const (
	NoErrCode = 0

	UnauthorizedErrCode  = 401
	InternalServeErrCode = 500

	IdNotIntegral           = 9999
	NoRecordErrCode         = 10000
	RepeatedUsernameErrCode = 10001
	WrongPasswordErrCode    = 10002

	BindingQueryErrCode = 40000
	WrongTimeCode       = 40001
	WrongLevelCode      = 40002
	//Unknown
)

func NormErr(c *gin.Context, errCode int) {
	c.JSON(http.StatusBadRequest, RespTemplate{
		Status: errCode,
		Info:   ErrorCodeMap[errCode].Error(),
	})
}
