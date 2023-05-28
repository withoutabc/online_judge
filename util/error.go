package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"online_judge/logs"
	"time"
)

var ErrorCodeMap = map[int]error{
	BlankAuthErrCode:       BlankAuthErr,
	WrongAuthFormatErrCode: WrongAuthFormatErr,
	InValidTokenErrCode:    InvalidTokenErr,

	UnauthorizedErrCode:  UnauthorizedErr,
	InternalServeErrCode: InternalServeErr,

	IdNotIntegral:           IdNotIntegralErr,
	NoRecordErrCode:         NoRecordErr,
	RepeatedUsernameErrCode: RepeatedUsernameErr,
	WrongPasswordErrCode:    WrongPasswordErr,
	UpdateFailErrCode:       UpdateFailErr,
	RepeatedTitleErrCode:    RepeatedTitleErr,

	BindingQueryErrCode: BindingQueryErr,
	WrongTimeCode:       WrongTimeCodeErr,
	WrongLevelCode:      WrongLevelCodeErr,
}

var (
	BlankAuthErr       = errors.New("请求头中auth为空")
	WrongAuthFormatErr = errors.New("请求头中auth格式有误")
	InvalidTokenErr    = errors.New("无效的Token")

	UnauthorizedErr  = errors.New("unauthorized")
	InternalServeErr = errors.New("internal serve error")

	IdNotIntegralErr    = errors.New("id is not a integral")
	NoRecordErr         = errors.New("no record")
	RepeatedUsernameErr = errors.New("repeated username")
	WrongPasswordErr    = errors.New("wrong password")
	UpdateFailErr       = errors.New("update failed")
	RepeatedTitleErr    = errors.New("repeated title")

	BindingQueryErr   = errors.New("binding error")
	WrongTimeCodeErr  = errors.New("wrong time")
	WrongLevelCodeErr = errors.New("wrong level")
)

const (
	NoErrCode = 0

	UnauthorizedErrCode  = 401
	InternalServeErrCode = 500

	BlankAuthErrCode       = 2003
	WrongAuthFormatErrCode = 2004
	InValidTokenErrCode    = 2005

	IdNotIntegral           = 9999
	NoRecordErrCode         = 10000
	RepeatedUsernameErrCode = 10001
	WrongPasswordErrCode    = 10002
	UpdateFailErrCode       = 10003 //update or delete error
	RepeatedTitleErrCode    = 10004

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

func Log(err error) {
	logs.Log().Error(time.Now().Format("2006-01-02 15:04:05"), zap.Error(err))
}
