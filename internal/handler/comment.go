package handler

import (
	"strconv"
	"w3-task/internal/domain/model"
	"w3-task/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
}

type AddCommentParam struct {
	Content string `json:"content" binding:"required"`
	PostId  uint   `json:"postId" binding:"required"`
}

// 创建评论
func (c *CommentController) createComment(ctx *gin.Context) {
	var param AddCommentParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(400, Response{
			Message: err.Error(),
			Success: false,
		})
		return
	}
	userId, _ := ctx.Get("userId")
	db, _ := repository.GetDb()
	//检查帖子是否存在
	var post model.Post
	db.Where("id = ?", param.PostId).First(&post)
	if post.ID == 0 {
		ctx.JSON(400, Response{
			Success: false,
			Message: "post not found",
		})
		return
	}
	comment := model.Comment{Content: param.Content, PostID: param.PostId, UserID: uint(userId.(uint))}
	db.Create(&comment)
	ctx.JSON(200, Response{
		Data:    comment,
		Message: "success",
		Success: true,
	})
}

// 获取某帖子的所有评论
func (c *CommentController) getCommentsByPostId(ctx *gin.Context) {
	var conmets []model.Comment
	postId := ctx.Param("id")
	if postId == "" {
		ctx.JSON(400, Response{
			Success: false,
			Message: "postId is required",
		})
		return
	}
	//检查帖子是否存在
	postIdConv, _ := strconv.Atoi(postId)
	db, _ := repository.GetDb()

	var post model.Post
	db.Where("id = ?", postIdConv).First(&post)
	if post.ID == 0 {
		ctx.JSON(400, Response{
			Success: false,
			Message: "post is not exist",
		})
		return
	}
	db.Where("post_id = ?", postIdConv).Find(&conmets)
	ctx.JSON(200, Response{
		Data:    conmets,
		Success: true,
	})

}
