package student

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server_template/base_type"
	"server_template/db"
	"server_template/model"
	"server_template/service"
	"server_template/service/account"
)

// AddStudent 添加学生信息
func AddStudent(c *gin.Context) {
	isAdmin, _ := account.IsAdmin(c)
	if !isAdmin {
		return
	}

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
	isAdmin, _ := account.IsAdmin(c)
	if !isAdmin {
		return
	}

	var students []model.Student

	if err := db.Mysql.Find(&students).Error; err != nil {
		service.HttpNotFound(c)
		return
	}

	service.HttpOK1(c, students)
}

// GetStudent 获取学生信息
func GetStudent(c *gin.Context) {
	isAdmin, _ := account.IsAdmin(c)
	if !isAdmin {
		return
	}

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
	isAdmin, _ := account.IsAdmin(c)
	if !isAdmin {
		return
	}

	var student model.Student

	if err := c.BindJSON(&student); err != nil {
		service.HttpParameterError(c)
		return
	}

	if err := db.Mysql.Updates(&student).Error; err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK(c)
}

// 删除学生信息

func DeleteStudent(c *gin.Context) {
	isAdmin, _ := account.IsAdmin1(c)
	if !isAdmin {
		return
	}

	id := c.Query("id")

	err := db.Mysql.Transaction(func(tx *gorm.DB) error {
		var student model.Student
		student.Username = base_type.NullableString{
			NullString: sql.NullString{
				String: id,
				Valid:  true,
			}}
		err1 := tx.Model(model.Student{}).First(&student, id).Error
		err1 = tx.Delete(&student).Error
		err1 = tx.Delete(&model.User{}, map[string]any{"username": student.Username}).Error
		return err1
	})

	if err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK(c)
}
