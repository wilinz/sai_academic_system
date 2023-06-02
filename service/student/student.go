package student

import (
	"github.com/gin-gonic/gin"
	"server_template/db"
	"server_template/model"
	"server_template/service"
)

// AddStudent 添加学生信息
func AddStudent(c *gin.Context) {
	var student model.Student
	if err := c.BindJSON(&student); err != nil {
		service.HttpParameterError(c)
		return
	}

	if err := db.Mysql.Create(&student).Error; err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK1(c, student)
}

// GetStudents 获取所有学生信息
func GetStudents(c *gin.Context) {
	var students []model.Student

	if err := db.Mysql.Find(&students).Error; err != nil {
		service.HttpNotFound(c)
		return
	}

	service.HttpOK1(c, students)
}

// GetStudent 获取学生信息
func GetStudent(c *gin.Context) {
	var student model.Student
	id := c.Query("id")

	if err := db.Mysql.Where("id = ?", id).First(&student).Error; err != nil {
		service.HttpNotFound(c)
		return
	}

	service.HttpOK1(c, student)
}

// UpdateStudent 更新学生信息
func UpdateStudent(c *gin.Context) {
	var student model.Student

	if err := c.BindJSON(&student); err != nil {
		service.HttpParameterError(c)
		return
	}

	if err := db.Mysql.Updates(&student).Error; err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK1(c, student)
}

// 删除学生信息

func DeleteStudent(c *gin.Context) {
	var student model.Student
	id := c.Query("id")

	if err := db.Mysql.Where("id = ?", id).First(&student).Error; err != nil {
		service.HttpNotFound(c)
		return
	}

	if err := db.Mysql.Delete(&student).Error; err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK(c)
}
