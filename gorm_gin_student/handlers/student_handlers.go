package handlers

import (
	"net/http"
	"strconv"

	"MyStuSys/gorm_gin_student/database"
	"MyStuSys/gorm_gin_student/models"
	"github.com/gin-gonic/gin"
)

// GetStudents 获取所有学生
func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

// CreateStudent 创建新学生
func CreateStudent(c *gin.Context) {
	var input models.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

// UpdateStudent 更新学生信息
func UpdateStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&student).Updates(input)
	c.JSON(http.StatusOK, student)
}

// DeleteStudent 删除学生
func DeleteStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
