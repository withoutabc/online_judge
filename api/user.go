package api

import (
	"github.com/gin-gonic/gin"
	"online_judge/model"
	"online_judge/service"
	"online_judge/util"
	"strconv"
)

type UserServiceImpl struct {
	UserService
}

func NewUserApi() *UserServiceImpl {
	return &UserServiceImpl{
		UserService: service.NewUserServiceImpl(),
	}
}

type UserService interface {
	Register(user model.User) int
	Login(user model.User) (model.RespLoginRole, int)
	ChangePwd(pwd model.ReqChangePwd) int
}

func (u *UserServiceImpl) Register(c *gin.Context) {
	//receive
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	code := u.UserService.Register(user)
	switch code {
	case util.InternalServeErrCode:
		util.RespInternalErr(c)
		return
	case util.RepeatedUsernameErrCode:
		util.NormErr(c, util.RepeatedUsernameErrCode)
		return
	}
	//response
	util.RespOK(c)
}

func (u *UserServiceImpl) Login(c *gin.Context) {
	//receive
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	respLogin, code := u.UserService.Login(user)
	switch code {
	case util.InternalServeErrCode:
		util.Find()
		util.RespInternalErr(c)
		return
	case util.NoRecordErrCode:
		util.NormErr(c, util.NoRecordErrCode)
		return
	case util.WrongPasswordErrCode:
		util.NormErr(c, util.WrongPasswordErrCode)
		return
	}
	util.RespNormSuccess(c, respLogin)
}

func (u *UserServiceImpl) ChangePassword(c *gin.Context) {
	//获取参数
	userId := c.Param("user_id")
	IntUserId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		util.NormErr(c, util.IdNotIntegral)
		return
	}
	var pwd model.ReqChangePwd
	if err := c.ShouldBind(&pwd); err != nil {
		util.NormErr(c, util.BindingQueryErrCode)
		return
	}
	pwd.UserId = IntUserId
	code := u.UserService.ChangePwd(pwd)
	switch code {
	case util.InternalServeErrCode:
		util.RespInternalErr(c)
		return
	case util.NoRecordErrCode:
		util.NormErr(c, util.NoRecordErrCode)
		return
	case util.WrongPasswordErrCode:
		util.NormErr(c, util.WrongPasswordErrCode)
		return
	}
	util.RespOK(c)
}

//func Refresh(c *gin.Context) {
//	//refresh_token
//	rToken := c.PostForm("refresh_token")
//	if rToken == "" {
//		util.RespParamErr(c)
//		return
//	}
//	_, err := service.ParseToken(rToken)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status": 2005,
//			"info":   "无效的Token",
//		})
//		return
//	}
//	//生成新的token
//	newAToken, newRToken, err := service.RefreshToken(rToken)
//	if err != nil {
//		fmt.Printf("err:%v", err)
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status": 400,
//			"info":   err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, model.RespToken{
//		Status: 200,
//		Info:   "refresh token success",
//		Data: model.Token{
//			Token:        newAToken,
//			RefreshToken: newRToken,
//		},
//	})
//}
