package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(server *gin.Engine) {

	//控制器
	authController := &AuthController{}
	postController := &PostController{}
	commentController := &CommentController{}

	//公开的路由
	public := server.Group("/api")
	{
		//用户注册/登录
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
		//文章
		public.GET("/post/getPostById/:id", postController.GetPostById)
		public.GET("/post/getAllPosts", postController.getAllPosts)
		//评论
		public.GET("/comment/getCommentsByPostId/:id", commentController.getCommentsByPostId)

	}

	//需要认证的路由
	protected := server.Group("/api")
	protected.Use(AuthHandler())
	{
		//文章
		protected.POST("/post/create", postController.CreatePost)
		protected.POST("/post/update", postController.updatePost)
		protected.DELETE("/post/delete/:id", postController.deletePostById)
		//评论
		protected.POST("/comment/create", commentController.createComment)
	}

}
