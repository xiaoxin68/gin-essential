package common

import (
	"gin-essential/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtkey = []byte("a_secret_create")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claimes := &Claims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "xiaoxin.com",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimes)
	tokenStr, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claimes := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claimes, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	return token, claimes, err
}
