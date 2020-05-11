package common

import (
	"github.com/dgrijalva/jwt-go"
	"tasker/server/constant"
	"tasker/server/model"
	"time"
)

var jwtKey = []byte("a_secret")

// 定义返回结构
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 颁发 token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(constant.MaxAge)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "hh",
			Subject:   "user_token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// 解析 token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
