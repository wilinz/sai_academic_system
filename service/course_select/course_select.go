package course_select

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server_template/db"
	"server_template/model"
	"server_template/service"
	"server_template/service/account"
	"strconv"
)

// SelectCourse 选课接口
func SelectCourse(c *gin.Context) {
	logged, username := account.IsLogged(c)
	if !logged {
		return
	}
	student0, err := account.GetStudentByUsername(username)
	if err != nil {
		service.HttpParameterError(c)
		return
	}
	// 获取课程 ID 和学生 ID
	courseID, err := strconv.ParseInt(c.PostForm("course_id"), 10, 64)
	studentID := student0.ID
	if err != nil {
		service.HttpParameterError(c)
		return
	}

	// 检查课程是否存在
	var course model.Course
	if err := db.Mysql.Where("id = ?", courseID).First(&course).Error; err != nil {
		service.HttpNotFound(c)
		return
	}

	// 检查是否已经选过这门课程
	var selection model.CourseSelection
	if err := db.Mysql.Where("course_id = ? AND student_id = ?", courseID, studentID).First(&selection).Error; err == nil {
		c.JSON(200, model.JsonResponse{
			Code: 403,
			Msg:  "不能重复选课",
			Data: nil,
		})
		return
	}

	// 检查课程是否已经超容量
	if course.Selected >= course.Capacity {
		c.JSON(200, model.JsonResponse{
			Code: 403,
			Msg:  "课程已经超容量",
			Data: nil,
		})
		return
	}

	// 创建选课信息
	selection = model.CourseSelection{CourseID: courseID, StudentID: studentID}
	err = db.Mysql.Transaction(func(tx *gorm.DB) error {
		err := db.Mysql.Create(&selection).Error
		// 更新课程已选人数
		err = db.Mysql.Model(&course).UpdateColumn("selected", course.Selected+1).Error
		return err
	})

	if err != nil {
		service.HttpServerInternalError(c)
	}

	service.HttpOK(c)
}

// DropCourse 退课接口
func DropCourse(c *gin.Context) {
	logged, username := account.IsLogged(c)
	if !logged {
		return
	}
	student, err := account.GetStudentByUsername(username)
	if err != nil {
		service.HttpParameterError(c)
		return
	}
	// 获取选课信息 ID
	courseID, err := strconv.ParseInt(c.Query("course_id"), 10, 64)
	if err != nil {
		service.HttpParameterError(c)
		return
	}

	//检查选课信息是否存在
	var selection model.CourseSelection
	if err := db.Mysql.Where(map[string]any{"course_id": courseID, "student_id": student.ID}).First(&selection).Error; err != nil {
		service.HttpNotFound(c)
		return
	}

	err = db.Mysql.Transaction(func(tx *gorm.DB) error {
		err := db.Mysql.Delete(&selection).Error
		err = tx.Model(&model.Course{}).Where("id = ?", selection.CourseID).
			UpdateColumn("selected", gorm.Expr("selected - ?", 1)).Error
		return err
	})

	if err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK(c)
}

// GetSelectedCourses 获取已选课程接口
func GetSelectedCourses(c *gin.Context) {
	logged, username := account.IsLogged(c)
	if !logged {
		return
	}
	student, err := account.GetStudentByUsername(username)
	if err != nil {
		service.HttpParameterError(c)
		return
	}

	// 查询选课信息和课程信息
	var courses []model.Course

	if err := db.Mysql.Table("course_selections").
		Select("courses.*").
		Joins("INNER JOIN courses ON courses.id = course_selections.course_id").
		Where("course_selections.student_id = ?", student.ID).
		Find(&courses).Error; err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK1(c, courses)
}

func getCourseSelectionsByCourseName(c *gin.Context) {

	var selections []model.CourseSelectionInfo
	courseName := c.Query("course_name")

	err := db.Mysql.
		Table("course_selections").
		Select("course_selections.*, courses.*, students.*").
		Joins("JOIN courses ON course_selections.course_id = courses.id").
		Joins("JOIN students ON course_selections.student_id = students.id").
		Where("courses.course_name LIKE ?", "%"+courseName+"%").
		Find(&selections).Error

	if err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK1(c, selections)
}

func getCourseSelectionsByStudentName(c *gin.Context) {

	var selections []model.CourseSelectionInfo
	studentName := c.Query("student_name")

	err := db.Mysql.
		Table("course_selections").
		Select("course_selections.*, courses.*, students.*").
		Joins("JOIN courses ON course_selections.course_id = courses.id").
		Joins("JOIN students ON course_selections.student_id = students.id").
		Where("students.name LIKE ?", "%"+studentName+"%").
		Find(&selections).Error

	if err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK1(c, selections)
}

func getAllCourseSelections(c *gin.Context) {

	var selections []model.CourseSelectionInfo

	err := db.Mysql.
		Table("course_selections").
		Select("course_selections.*, courses.*, students.*").
		Joins("JOIN courses ON course_selections.course_id = courses.id").
		Joins("JOIN students ON course_selections.student_id = students.id").
		Find(&selections).Error

	if err != nil {
		service.HttpServerInternalError(c)
		return
	}

	service.HttpOK1(c, selections)
}

func GetCourseSelections(c *gin.Context) {
	isAdmin, _ := account.IsAdmin(c)
	if !isAdmin {
		return
	}

	if c.Query("course_name") != "" {
		getCourseSelectionsByCourseName(c)
	} else if c.Query("student_name") != "" {
		getCourseSelectionsByStudentName(c)
	} else {
		getAllCourseSelections(c)
	}
}
