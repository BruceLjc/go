package models

import "gorm.io/gorm"

// Student 学生模型定义
type Student struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}
