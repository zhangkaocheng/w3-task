package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(server *gin.Engine) {

	//控制器
	authController := AuthController{}

	//公开的路由
	public := server.Group("/api")
	{
		//用户注册/登录
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
		//文章
		//评论
	}

	//需要认证的路由
	// protected := server.Group("/api")
	// protected.Use(AuthHandler())

	{
		//文章
		//评论
	}

}
