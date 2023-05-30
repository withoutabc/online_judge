package model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	UserId   int64  `json:"user_id" form:"user_id" binding:"-" gorm:"primarykey" `
	Username string `json:"username" form:"username" binding:"required" gorm:"type:varchar(40);not null"`
	Password string `json:"password" form:"password" binding:"required" gorm:"not null;type:longblob"`
	Salt     []byte `json:"salt" form:"salt" binding:"-" gorm:"not null"`
}

type MyClaims struct {
	UserId int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

//request

type ReqChangePwd struct {
	UserId      int64  `json:"user_id"`
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}

//response

type RespLogin struct {
	UserId    int64  `json:"user_id"`
	LoginTime string `json:"login_time"`
	Token     string `json:"token"`
}

type RespLoginRole struct {
	RespLogin
	Role string `json:"role"`
}

// BeforeCreate uses snowflake to generate an ID.
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	//if u.UserId != 0 {
	//	return nil
	//}
	//sf, err := snowflake.NewNode(0)
	//if err != nil {
	//	log.Fatalf("generate id failed: %s", err.Error())
	//	return err
	//}
	//u.UserId = sf.Generate().Int64()
	return nil
}

func (u *User) AfterCreate(_ *gorm.DB) (err error) {

	return nil
}
