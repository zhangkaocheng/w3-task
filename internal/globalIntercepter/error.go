package globalintercepter

import (
	"net/http"
	"w3-task/internal/handler"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		var errMsg string
		switch v := err.(type) {
		case string:
			c.JSON(http.StatusBadRequest, handler.Response{
				Success: false,
				Message: errMsg,
			})
		case error:
			errMsg = v.Error()
			c.JSON(http.StatusInternalServerError, handler.Response{
				Success: false,
				Message: errMsg,
			})
		default:
			errMsg = "Internal Server Error"
			c.JSON(http.StatusInternalServerError, handler.Response{
				Success: false,
				Message: errMsg,
			})

		}
		c.Abort()
	})

}
