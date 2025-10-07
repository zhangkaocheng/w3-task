package handler

import (
	"net/http"
	"w3-task/internal/domain/model"
	"w3-task/internal/domain/repository"
	"w3-task/pkg/util"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

type LoginParam struct {
	Username string `json:"username" bunding:"required"`
	Password string `json:"password" bunding:"required"`
}

type RegisterParam struct {
	Username string `json:"username" bunding:"required"`
	Password string `json:"password" bunding:"required"`
	Email    string `json:"email" bunding:"required,min=6"`
}

// 注册
func (c *AuthController) Register(ctx *gin.Context) {
	var param RegisterParam

	//参数校验
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(400, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	//用户名/邮箱唯一性校验
	db, err := repository.GetDb()
	if err != nil {
		ctx.JSON(500, Response{
			Success: false,
			Message: "数据库连接失败",
		})
		return
	}
	var user model.User
	result := db.Where("username = ? OR email = ?", param.Username, param.Email).First(&user)
	if result.Error == nil {
		ctx.JSON(400, Response{
			Success: false,
			Message: "用户名或邮箱已存在",
		})
		return
	}
	//密码加密
	hashedPassword, err := util.EncryptePassword(param.Password)
	if err != nil {
		ctx.JSON(500, Response{
			Success: false,
			Message: "密码加密失败",
		})
		return
	}
	//创建用户
	if err := db.Create(&model.User{Username: param.Username, Password: hashedPassword, Email: param.Email}); err != nil {
		ctx.JSON(500, Response{
			Success: false,
			Message: "用户创建失败",
		})
	}
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "用户创建成功",
	})

}

//登录

func (c *AuthController) Login(ctx *gin.Context) {

	var param LoginParam

	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(400, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	//获取用户
	db, err := repository.GetDb()
	if err != nil {
		ctx.JSON(500, Response{
			Success: false,
			Message: "数据库连接失败",
		})
		return
	}
	var user model.User
	result := db.Where("username = ?", param.Username).First(&user)
	if result.Error != nil {
		ctx.JSON(400, Response{
			Success: false,
			Message: "用户不存在",
		})
		return
	}
	if !util.CheckPassword(user.Password, param.Password) {
		ctx.JSON(400, Response{
			Success: false,
			Message: "密码错误",
		})
		return
	}

	//生成token
	token, err := util.GenerateToken(user.ID, user.Username)
	if err != nil {
		ctx.JSON(500, Response{
			Success: false,
			Message: "token生成失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "登录成功",
		Data:    token,
	})
}
