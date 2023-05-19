package service

import (
	"gorm.io/gorm"
	"log"
	"online_judge/dao"
	"online_judge/model"
	"online_judge/util"
	"time"
)

type UserDao interface {
	CreateUser(user *model.User) error
	SearchUserById(uid int64) (user model.User, err error)
	SearchUserByName(username string) (user model.User, err error)
	ChangePwd(uid int64, password string, salt []byte) error
}

func NewUserServiceImpl() *UserDaoImpl {
	return &UserDaoImpl{
		UserDao: dao.NewUserDao(),
	}
}

type UserDaoImpl struct {
	UserDao
}

func (u *UserDaoImpl) Register(user model.User) int {
	//检索数据库
	mysqlUser, err := u.UserDao.SearchUserByName(user.Username)
	if mysqlUser.Username != "" {
		return util.RepeatedUsernameErrCode
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return util.InternalServeErrCode
	}
	//生成盐值
	var salt []byte
	salt, err = util.GenerateSalt()
	if err != nil {
		log.Println(err)
		return util.InternalServeErrCode
	}
	//加密
	hashedPassword := util.HashWithSalt(user.Password, salt)
	//用户信息写入数据库
	user.Password = string(hashedPassword)
	user.Salt = salt
	if err = u.UserDao.CreateUser(&user); err != nil {
		util.Find()
		return util.InternalServeErrCode
	}
	return util.NoErrCode
}

func (u *UserDaoImpl) Login(user model.User) (model.RespLogin, int) {
	mysqlUser, err := u.UserDao.SearchUserByName(user.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Find()
			return model.RespLogin{}, util.NoRecordErrCode
		} else {
			util.Find()
			return model.RespLogin{}, util.InternalServeErrCode
		}
	}
	//if password right
	if string(util.HashWithSalt(user.Password, mysqlUser.Salt)) != mysqlUser.Password {
		util.Find()
		log.Println(mysqlUser.Password)
		log.Println(string(util.HashWithSalt(user.Password, mysqlUser.Salt)))
		return model.RespLogin{}, util.WrongPasswordErrCode
	}
	//generate token
	token, _, err := util.GenToken(mysqlUser.UserId)
	if err != nil {
		util.Find()
		return model.RespLogin{}, util.InternalServeErrCode
	}
	return model.RespLogin{
		UserId:    mysqlUser.UserId,
		LoginTime: time.Now().Format("2006-01-02 15:04:05"),
		Token:     token,
	}, util.NoErrCode
}

func (u *UserDaoImpl) ChangePwd(pwd model.ReqChangePwd) int {
	//检索数据库
	mysqlUser, err := u.UserDao.SearchUserById(pwd.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Find()
			return util.NoRecordErrCode
		} else {
			util.Find()
			return util.InternalServeErrCode
		}
	}
	//if password right
	if string(util.HashWithSalt(pwd.OldPassword, mysqlUser.Salt)) != mysqlUser.Password {
		util.Find()
		return util.WrongPasswordErrCode
	}
	//生成盐值
	var salt []byte
	salt, err = util.GenerateSalt()
	if err != nil {
		return util.InternalServeErrCode
	}
	//加密
	hashedPassword := util.HashWithSalt(pwd.NewPassword, salt)
	//修改密码和盐值
	err = u.UserDao.ChangePwd(pwd.UserId, string(hashedPassword), salt)
	if err != nil {
		return util.InternalServeErrCode
	}
	return util.NoErrCode
}
