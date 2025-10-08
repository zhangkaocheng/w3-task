package util

import (
	"fmt"
	"strings"
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
	expireTime := time.Now().Add(time.Hour * time.Duration(configs.GetGWTConfig().ExpireHour))
	claims := Claims{
		UserId:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.GetGWTConfig().Secret))
}

// 解析token
func ParseToken(tokenString string) (*Claims, error) {
	//去掉Bearer前缀
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	// 检查 token 基本格式
	if tokenString == "" {
		return nil, fmt.Errorf("token is empty")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.GetGWTConfig().Secret), nil
	})

	if err != nil {
		// 返回具体的错误信息，帮助调试
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
