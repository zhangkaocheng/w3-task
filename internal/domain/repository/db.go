package repository

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDb() error {
	dsn := "root:root@/my-blog?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}
	fmt.Println("success to connect database")
	return nil
}

// 设置连接池参数
func SetConnectionPool(maxOpenConns, maxIdleConns, connMaxLifetime int) {

	db, err := DB.DB()
	if err != nil {
		panic("failed to get database connection pool: " + err.Error())
	}
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
}
