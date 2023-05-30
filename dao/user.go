package dao

import (
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
	"online_judge/model"
)

func NewUserDao() *UserDaoImpl {
	return &UserDaoImpl{
		db: DB,
		e:  E,
	}
}

type UserDaoImpl struct {
	db *gorm.DB
	e  *casbin.Enforcer
}

func (u *UserDaoImpl) GetRole(uid string) (bool, error) {
	return E.Enforce(uid, "admin_data", "read")
}

func (u *UserDaoImpl) CreateUser(tx *gorm.DB, user *model.User) error {
	result := tx.Create(&user)
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
