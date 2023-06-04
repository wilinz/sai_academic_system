package course

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"server_template/db"
	"server_template/model"
	"server_template/service"
	"server_template/service/account"
)

// AddCourse 添加学生信息
func AddCourse(c *gin.Context) {
	isAdmin, _ := account.IsAdmin(c)
	if !isAdmin {
		return
	}

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
	isAdmin, _ := account.IsAdmin(c)
	if !isAdmin {
		return
	}

	var courses []model.Course

	if err := db.Mysql.Find(&courses).Error; err != nil {
		service.HttpNotFound(c)
		return
	}

	service.HttpOK1(c, courses)
}

// GetCourse 获取学生信息
func GetCourse(c *gin.Context) {
	isAdmin, _ := account.IsAdmin(c)
	if !isAdmin {
		return
	}

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
	isAdmin, _ := account.IsAdmin(c)
	if !isAdmin {
		return
	}

	var course model.Course

	if err := c.BindJSON(&course); err != nil {
		service.HttpParameterError(c)
		return
	}

	//设为零值
	course.Selected = 0

	if err := db.Mysql.Updates(&course).Error; err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK1(c, course)
}

// 删除学生信息

func DeleteCourse(c *gin.Context) {
	isAdmin, _ := account.IsAdmin(c)
	if !isAdmin {
		return
	}

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

func GetSelectableCourse(c *gin.Context) {
	//logged, _ := account.IsLogged(c)
	//if !logged {
	//	return
	//}

	var student model.Student
	student.Username = sql.NullString{
		String: "3397733901@qq.com",
		Valid:  true,
	}
	err := db.Mysql.Where(map[string]any{"username": student.Username}).First(&student).Error
	if err != nil {
		service.HttpServerInternalError(c)
		return
	}

	var courses []model.Course
	err = db.Mysql.Where("grade = ? AND selected < capacity", student.Grade).Find(&courses).Error
	if err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK1(c, courses)
}
