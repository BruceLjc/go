package routes

import (
	"MyStuSys/gorm_gin_student/handlers"

	"github.com/gin-gonic/gin"

	// 在你的 Golang 项目中导入 cors
	"github.com/gin-contrib/cors"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	//中间件,后来加的, 为了连前端
	r.Use(cors.Default())
	r.GET("/students", handlers.GetStudents)
	r.POST("/students", handlers.CreateStudent)
	r.PUT("/students/:id", handlers.UpdateStudent)
	r.DELETE("/students/:id", handlers.DeleteStudent)

	return r
}
