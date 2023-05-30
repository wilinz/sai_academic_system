package account

import (
	"github.com/gin-gonic/gin"
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

		db.Mysql.Create(&user)
		service.HttpOK(c)
	}

}
