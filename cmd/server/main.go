package main

import (
	"fmt"
	"log"
	"w3-task/configs"
	"w3-task/internal/domain/model"
	"w3-task/internal/domain/repository"
	globalintercepter "w3-task/internal/globalIntercepter"

	"github.com/gin-gonic/gin"
)

func main() {

	// 加载配置
	cfg, err := configs.LoadConfig("../../configs/config.yaml")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// 初始化数据库
	err = repository.GetDb()
	if err != nil {
		log.Fatal("Failed to init database: ", err)
	}

	// 设置连接池参数
	repository.SetConnectionPool(
		cfg.MySQL.MaxOpenConns,
		cfg.MySQL.MaxIdleConns,
		cfg.MySQL.ConnMaxLifetime,
	)
	//自动迁移
	repository.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	fmt.Println("AutoMigrate...")

	// gin框架初始化路由
	server := gin.New()
	//设置错误处理拦截
	server.Use(globalintercepter.ErrorHandler())
	//设置权限处理拦截

	// 启动服务
	err = server.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
