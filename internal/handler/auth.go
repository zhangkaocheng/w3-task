package handler

import (
	"net/http"
	"strings"
	"w3-task/pkg/util"

	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, Response{
				Success: false,
				Message: "Unauthorized",
			})
			ctx.Abort()
			return
		}
		tokenStr := strings.TrimPrefix(token, "Bearer ")
		Claims, err := util.ParseToken(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, Response{
				Success: false,
				Message: "Unauthorized",
			})
			ctx.Abort()
			return
		}
		ctx.Set("userId", Claims.UserId)
		ctx.Set("username", Claims.Username)
		ctx.Next()
	}
}
