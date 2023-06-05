package route

import (
	"github.com/gin-gonic/gin"
	"server_template/common"
	"server_template/service/account"
	"server_template/service/appversion"
	"server_template/service/bing"
	"server_template/service/course"
	"server_template/service/course_select"
	"server_template/service/feedback"
	"server_template/service/proxy"
	"server_template/service/student"
	"server_template/service/welcome"
)

func Run() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(func(c *gin.Context) {
		/*	log.Println(c.Request.Header)
			c.Next()
			log.Println(c.Writer.Header())*/
	})
	r.Static(common.UploadDir, common.UploadSavePath)
	accountRouter := r.Group("/account")
	accountRouter.POST("/login", account.LoginHandler)
	accountRouter.POST("/login_with_code", account.LoginWithCodeHandler)
	accountRouter.POST("/register", account.RegisterHandler)
	accountRouter.POST("/verify", account.VerificationCodeHandler)
	accountRouter.PUT("/password/reset", account.ResetPasswordHandler)
	accountRouter.PUT("/password/change", account.ChangePasswordHandler)
	accountRouter.PUT("/info", account.SetUserInfoHandler)
	accountRouter.GET("/info", account.GetUserInfoHandler)
	accountRouter.DELETE("/logout", account.LogoutHandler)

	r.GET("/welcome", welcome.WelcomeHandler)

	r.POST("/feedback", feedback.PostFeedbackHandler)

	r.GET("/app_version", appversion.GetAppVersion)

	r.Any("/proxy", proxy.HttpProxyHandler)

	r.GET("/student", student.GetStudent)
	r.GET("/students", student.GetStudents)
	r.DELETE("/student", student.DeleteStudent)
	r.PUT("/student", student.UpdateStudent)
	r.POST("/student", student.AddStudent)

	r.GET("/course", course.GetCourse)
	r.GET("/course/selectable", course.GetSelectableCourse)
	r.GET("/courses", course.GetCourses)
	r.DELETE("/course", course.DeleteCourse)
	r.PUT("/course", course.UpdateCourse)
	r.POST("/course", course.AddCourse)

	r.GET("/bing/daily_image", bing.GetDailyImage)

	r.POST("/course/select", course_select.SelectCourse)
	r.DELETE("/course/unselect", course_select.DropCourse)
	r.GET("/course/selected", course_select.GetSelectedCourses)
	r.GET("/course/selected/admin", course_select.GetCourseSelections)

	err := r.Run(":10011")
	if err != nil {
		panic(err)
	}
}
