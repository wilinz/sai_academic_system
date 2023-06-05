package account

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server_template/base_type"
	"server_template/constant/code_type"
	"server_template/constant/error_code"
	"server_template/db"
	"server_template/model"
	"server_template/service"
	"server_template/util"
)

func RegisterHandler(c *gin.Context) {
	var p model.RegistrationParameters
	err := c.Bind(&p)
	if err != nil {
		service.HttpParameterError(c)
		return
	}

	if !CheckPasswordIsLegal(c, p.Password) {
		return
	}

	if IsAccountExists(c, p.Username) {
		return
	}

	//检查学号是否存在
	var student model.Student
	err = db.Mysql.Model(model.Student{}).Where(map[string]any{"student_no": p.StudentNo}).First(&student).Error
	if gorm.ErrRecordNotFound == err {
		c.JSON(200, model.JsonResponse{
			Code: error_code.StudentNotExist,
			Msg:  "该学生不存在，请联系管理员添加您的学生信息",
			Data: nil,
		})
		return
	}

	if len(student.Username.String) != 0 {
		c.JSON(200, model.JsonResponse{
			Code: error_code.StudentAssociated,
			Msg:  "该学生已被注册关联，请联系管理员处理",
			Data: nil,
		})
		return
	}

	if IsVerificationCodeCorrect(c, p.VerificationCode, code_type.Register, p.Username) {
		salt := util.GetRandomString(6)
		var user = model.User{
			ID:               0,
			UserInfoReadOnly: model.UserInfoReadOnly{Username: p.Username},
			Password:         util.Sha256Sum(util.Sha256Sum(p.Password) + salt),
			Salt:             salt,
		}

		if util.IsEmail(p.Username) {
			user.Email = p.Username
		} else if util.IsPhone(p.Username) {
			user.Phone = p.Username
		} else {
			c.JSON(200, model.JsonResponse{
				Code: error_code.UsernameError,
				Msg:  "手机号或邮箱错误",
				Data: nil,
			})
			return
		}

		student.Username = base_type.NullableString{
			NullString: sql.NullString{
				String: p.Username,
				Valid:  true,
			},
		}
		user.StudentNo = p.StudentNo

		err := db.Mysql.Transaction(func(tx *gorm.DB) error {
			err2 := tx.Updates(&student).Error
			err2 = tx.Create(&user).Error
			return err2
		})

		if err != nil {
			service.HttpServerInternalError(c)
			return
		}
		service.HttpOK(c)
	}

}
