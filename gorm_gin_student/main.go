package main

import (
	"MyStuSys/gorm_gin_student/database"
	"MyStuSys/gorm_gin_student/routes"
)

func main() {
	// 连接数据库
	database.ConnectDB()

	// 设置路由
	r := routes.SetupRouter()

	// 启动应用
	r.Run("localhost:8080")
}
