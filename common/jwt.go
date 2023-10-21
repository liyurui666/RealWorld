package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	Username string
	Email    string
	jwt.StandardClaims
}

// 密钥
var jwtKey = []byte("MDk4ZjZiY2Q0NjIxZDM3M2NhZGU0ZTgzMjYyN2I0ZjY")

// GetToken 生成token
func GetToken(username string, email string) (string, error) {
	expirationTime := time.Now().Add(3600 * 1000 * time.Minute)
	UserClaim := &UserClaims{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *UserClaims, error) {
	UserClaim := &UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, UserClaim, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, UserClaim, err
}
