package appversion

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server_template/db"
	"server_template/model"
	"server_template/service"
)

func GetAppVersion(c *gin.Context) {
	// 获取参数
	appid := c.Query("appid")
	// 从数据库中查询最新版本信息
	latestVersion := new(model.AppVersion)
	err := db.Mysql.Where("appid = ?", appid).Order("version_code desc").First(latestVersion).Error
	if err == gorm.ErrRecordNotFound {
		latestVersion = nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		service.HttpServerInternalError(c)
		return
	}
	// 返回结果
	c.JSON(200, model.JsonResponse{
		Code: 200,
		Msg:  "ok",
		Data: latestVersion,
	})
}
