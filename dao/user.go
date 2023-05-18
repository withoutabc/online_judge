package dao

import (
	"gorm.io/gorm"
	"online_judge/model"
)

func NewUserDao() *UserDaoImpl {
	return &UserDaoImpl{
		db: DB,
	}
}

type UserDaoImpl struct {
	db *gorm.DB
}

func (u *UserDaoImpl) CreateUser(user *model.User) error {
	result := DB.Create(&user)
	return result.Error
}

func (u *UserDaoImpl) SearchUserById(uid int64) (user model.User, err error) {
	result := u.db.Where(&model.User{UserId: uid}).First(&user)
	return user, result.Error
}

func (u *UserDaoImpl) SearchUserByName(username string) (user model.User, err error) {
	result := u.db.Where(&model.User{Username: username}).First(&user)
	return user, result.Error
}

func (u *UserDaoImpl) ChangePwd(uid int64, password string, salt []byte) error {
	result := u.db.Model(&model.User{}).Where(&model.User{UserId: uid}).Updates(model.User{Password: password, Salt: salt})
	return result.Error
}
