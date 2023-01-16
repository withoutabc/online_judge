package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"online_judge/model"
	"time"
)

var Secret = []byte("YJX")

func keyFunc(*jwt.Token) (i interface{}, err error) {
	return Secret, nil
}

// GenToken GenToken生成access token和refresh token
func GenToken(uid string) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	c := model.MyClaims{
		Uid: uid, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // 过期时间
			Issuer:    "YJX",                                // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(Secret)
	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(), // 过期时间
		Issuer:    "YJX",                                       // 签发人
	}).SignedString(Secret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken, uid string, err error) {
	// 从旧access token中解析出claims数据 解析出payload负载信息
	var claims model.MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", "", "", errors.New("invalid token signature")
		}
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				//当access token是过期错误并且refresh token没有过期时就创建一个新的access token
				_, err = jwt.ParseWithClaims(rToken, &claims, keyFunc)
				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						return "", "", "", errors.New("invalid refresh token signature")
					}
					if ve, ok = err.(*jwt.ValidationError); ok {
						if ve.Errors&jwt.ValidationErrorExpired != 0 {
							return "", "", "", errors.New("refresh token expired")
						}
					}
				}
				//生成新的token
				newAToken, newRToken, err = GenToken(claims.Uid)
				if err != nil {
					return "", "", "", err
				}
				return newAToken, newRToken, claims.Uid, nil
			}
		}
	}
	return "", "", "", errors.New("invalid token")
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*model.MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
