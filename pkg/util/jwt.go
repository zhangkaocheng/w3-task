package util

import (
	"time"
	"w3-task/configs"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成token
func GenerateToken(userId uint, username string) (string, error) {
	// 设置过期时间24小时
	expireTime := time.Now().Add(time.Hour * time.Duration(configs.JWTConf.ExpireHour))
	claims := Claims{
		UserId:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.JWTConf.Secret))
}

// 解析token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.JWTConf.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, nil
}
