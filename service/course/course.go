package course

import (
	"github.com/gin-gonic/gin"
	"server_template/db"
	"server_template/model"
	"server_template/service"
)

// AddCourse 添加学生信息
func AddCourse(c *gin.Context) {
	var course model.Course
	if err := c.BindJSON(&course); err != nil {
		service.HttpParameterError(c)
		return
	}

	if err := db.Mysql.Create(&course).Error; err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK1(c, course)
}

// GetCourses 获取所有学生信息
func GetCourses(c *gin.Context) {
	var courses []model.Course

	if err := db.Mysql.Find(&courses).Error; err != nil {
		service.HttpNotFound(c)
		return
	}

	service.HttpOK1(c, courses)
}

// GetCourse 获取学生信息
func GetCourse(c *gin.Context) {
	var course model.Course
	id := c.Query("id")

	if err := db.Mysql.Where("id = ?", id).First(&course).Error; err != nil {
		service.HttpNotFound(c)
		return
	}

	service.HttpOK1(c, course)
}

// UpdateCourse 更新学生信息
func UpdateCourse(c *gin.Context) {
	var course model.Course

	if err := c.BindJSON(&course); err != nil {
		service.HttpParameterError(c)
		return
	}

	if err := db.Mysql.Updates(&course).Error; err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK1(c, course)
}

// 删除学生信息

func DeleteCourse(c *gin.Context) {
	var course model.Course
	id := c.Query("id")

	if err := db.Mysql.Where("id = ?", id).First(&course).Error; err != nil {
		service.HttpNotFound(c)
		return
	}

	if err := db.Mysql.Delete(&course).Error; err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK(c)
}
