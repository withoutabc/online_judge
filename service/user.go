package service

import (
	"online_judge/dao"
	"online_judge/model"
)

func SearchUserByUsername(username string) (u model.User, err error) {
	u, err = dao.SearchUserByUsername(username)
	return
}

func SearchUserByUid(uid string) (u model.User, err error) {
	u, err = dao.SearchUserByUid(uid)
	return
}
func CreateUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}

func ChangePassword(newPassword, username string) error {
	err := dao.UpdatePassword(newPassword, username)
	return err
}
