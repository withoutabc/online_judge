package api

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"online_judge/model"
	"online_judge/service"
	"online_judge/util"
	"strconv"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	confPassword := c.PostForm("confirmPassword")
	//判断是否有效输入
	if username == "" || password == "" || confPassword == "" {
		util.RespParamErr(c)
		return
	}
	//检索数据库
	u, err := service.SearchUserByUsername(username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RespInternalErr(c)
		return
	}
	//用户是否存在
	if u.Username != "" {
		util.NormErr(c, 400, "user has existed")
		return
	}
	//两次密码是否一致
	if confPassword != password {
		util.NormErr(c, 400, "different password")
		return
	}
	//生成盐值
	var salt []byte
	salt, err = service.GenerateSalt()
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	//加密
	hashedPassword := service.HashWithSalt(password, salt)
	//用户信息写入数据库
	err = service.CreateUser(model.User{
		Username: username,
		Password: hashedPassword,
		Salt:     salt,
	})
	if err != nil {
		fmt.Printf("create user err:%v", err)
		util.RespInternalErr(c)
		return
	}
	//查找用户
	u, err = service.SearchUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 400, "user don't exist")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	util.RespOK(c)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//有效输入
	if username == " " || password == "" {
		util.RespParamErr(c)
		return
	}
	//查找用户
	u, err := service.SearchUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 400, "user don't exist")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	//对输入密码加密
	hashedPassword := service.HashWithSalt(password, u.Salt)
	//转化密码，对比
	if bytes.Equal(hashedPassword, u.Password) == false {
		util.NormErr(c, 400, "wrong password")
		return
	}
	// 正确则登录成功
	aToken, rToken, _ := service.GenToken(strconv.Itoa(u.Uid))
	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"info":          "login success",
		"uid":           u.Uid,
		"token":         aToken,
		"refresh_token": rToken,
	})
}

func Refresh(c *gin.Context) {
	aToken := c.PostForm("token")
	rToken := c.PostForm("refresh_token")
	if aToken == "" || rToken == "" {
		util.RespParamErr(c)
		return
	}
	newAToken, newRToken, uid, err := service.RefreshToken(aToken, rToken)
	if err != nil {
		fmt.Printf("err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"info":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"uid":           uid,
		"token":         newAToken,
		"refresh_token": newRToken,
	})
}

func ChangePassword(c *gin.Context) {
	//获取参数
	uid := c.Param("uid")
	password := c.PostForm("password")
	newPassword := c.PostForm("newPassword")
	confPassword := c.PostForm("confirmPassword")
	//有效输入
	if password == "" || newPassword == "" || confPassword == "" {
		util.RespParamErr(c)
		return
	}
	//密码是否一致
	if confPassword != newPassword {
		util.NormErr(c, 400, "different password")
		return
	}
	if password == newPassword {
		util.NormErr(c, 400, "same password")
		return
	}
	//检索用户处理
	u, err := service.SearchUserByUid(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 400, "user don't exist")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	//对输入密码加密
	hashedPassword := service.HashWithSalt(password, u.Salt)
	//判断密码
	if bytes.Equal(hashedPassword, u.Password) == false {
		util.NormErr(c, 400, "wrong password")
		return
	}
	//密码正确,生成盐值
	u, err = service.SearchUserByUid(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 400, "user don't exist")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	var salt []byte
	salt, err = service.GenerateSalt()
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	//加密
	hashedPassword = service.HashWithSalt(newPassword, salt)
	//修改密码和盐值
	err = service.ChangePassword(hashedPassword, u.Username, salt)
	if err != nil {
		log.Printf("update password error:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
