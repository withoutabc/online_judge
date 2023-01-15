package service

import (
	"crypto/rand"
	"crypto/sha256"
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

func ChangePassword(newPassword []byte, username string, salt []byte) error {
	err := dao.UpdatePassword(newPassword, username, salt)
	return err
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	return salt, err

}

func HashWithSalt(password string, salt []byte) []byte {
	salted := append(salt, []byte(password)...)
	hashed := sha256.Sum256(salted)
	return hashed[:]
}
