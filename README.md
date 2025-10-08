# 博客系统

## 项目介绍
该项目使用 Go 语言结合 Gin 框架和 GORM 库开发的个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、
读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能
 
## 运行环境
1.go1.24.5  
2.mysql5.7  
3.gin1.6.3  
4.git 2.20.1  
5.gorm.io/gorm v1.31.0


## 依赖安装步骤
1.进入项目根目录并执行：  
go mod init w3-task  
go mod tidy  

2.相关依赖包安装  
go get -u github.com/gin-gonic/gin  
go get -u github.com/jinzhu/gorm  
go get -u github.com/go-sql-driver/mysql  

## 项目启动
1.进入项目根目录并执行：  
go run main.go  

2.访问：   
http://localhost:8080/api/XXX  

## 测试用例
详见 /w3-task/doc/my-blog.postman_collection.json

