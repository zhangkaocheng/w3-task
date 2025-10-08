package handler

import (
	"strconv"
	"w3-task/internal/domain/model"
	"w3-task/internal/domain/repository"
	"w3-task/pkg/util"

	"github.com/gin-gonic/gin"
)

type PostController struct {
}

type AddPostParam struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostParam struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Id      uint
}

// 创建帖子
func (c *PostController) CreatePost(ctx *gin.Context) {
	var param AddPostParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(400, Response{
			Success: false,
			Message: "参数错误",
		})
		return
	}
	claims, err := util.ParseToken(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(401, Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}
	post := model.Post{
		Title:   param.Title,
		Content: param.Content,
		UserID:  claims.UserId,
	}
	db, err := repository.GetDb()
	if err != nil {
		ctx.JSON(500, Response{
			Success: false,
			Message: "数据库连接失败",
		})
		return
	}
	result := db.Create(&post)
	if result.Error != nil {
		ctx.JSON(500, Response{
			Success: false,
			Message: "创建帖子失败",
		})
		return
	}
	ctx.JSON(200, Response{
		Success: true,
		Message: "创建帖子成功",
	})
}

// 获取单个帖子
func (c *PostController) GetPostById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, Response{
			Success: false,
			Message: "参数错误",
		})
		return
	}
	strconvId, _ := strconv.Atoi(id)
	var post model.Post
	db, _ := repository.GetDb()
	result := db.Preload("User").First(&post, "id = ?", strconvId)
	if result.Error != nil {
		ctx.JSON(500, Response{
			Success: false,
			Message: "获取帖子失败",
		})
		return
	}
	ctx.JSON(200, Response{
		Success: true,
		Data:    post,
	})

}

// 获取所有帖子
func (c *PostController) getAllPosts(ctx *gin.Context) {
	var posts []model.Post
	db, _ := repository.GetDb()
	result := db.Preload("User").Find(&posts)
	if result.Error != nil {
		ctx.JSON(500, Response{
			Success: false,
			Message: "获取帖子失败",
		})
		return
	}
	ctx.JSON(200, Response{
		Success: true,
		Data:    posts,
	})

}

// 更新帖子
func (c *PostController) updatePost(ctx *gin.Context) {
	var param UpdatePostParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(400, Response{
			Success: false,
			Message: "参数错误",
		})
		return
	}
	claims, err := util.ParseToken(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(401, Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}
	updatePost := model.Post{Title: param.Title, Content: param.Content, UserID: claims.UserId}
	db, _ := repository.GetDb()
	UserId := ctx.MustGet("userId").(uint)
	//查询当前更新的帖子是否属于当前用户
	var post model.Post
	db.Where("id = ?", param.Id).First(&post)
	if post.UserID != UserId {
		ctx.JSON(400, Response{
			Success: false,
			Message: "current user is not the owner of this post",
		})
		return
	}
	//更新帖子
	db.Model(&post).Updates(updatePost)

}

// 删除帖子
func (c *PostController) deletePostById(ctx *gin.Context) {
	UserId := ctx.MustGet("userId").(uint)
	PostId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, Response{
			Success: false,
			Message: "Invalid id",
		})
		return
	}
	db, _ := repository.GetDb()

	var post model.Post

	err = db.Where("id=?", PostId).First(&post).Error
	if err != nil {
		ctx.JSON(400, Response{
			Success: false,
			Message: "Post not found",
		})
		return
	}

	db.Where("id=?", PostId).First(&post)
	if post.UserID != UserId {
		ctx.JSON(400, Response{
			Success: false,
			Message: "current user is not the owner of this post",
		})
		return
	}
	//删除帖子
	err = db.Where("id=?", PostId).Delete(&post).Error
	if err != nil {
		ctx.JSON(400, Response{
			Success: false,
			Message: "delete post failed",
		})
		return
	}
	ctx.JSON(200, Response{
		Success: true,
		Message: "delete post success",
	})

}
