package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Password []byte `json:"password"`
	Salt     []byte `json:"salt"`
}

type MyClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}
