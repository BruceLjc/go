package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

var DB *gorm.DB

// ConnectDB 连接数据库
func ConnectDB() {
	dsn := "root:root@tcp(0.0.0.0:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db

	// 自动迁移数据库结构
	err = DB.AutoMigrate(&Student{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
