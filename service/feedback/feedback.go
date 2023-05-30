package feedback

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/weili71/go-filex"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"server_template/common"
	"server_template/db"
	"server_template/model"
	"server_template/service"
	"strings"
)

// PostFeedbackHandler 定义处理反馈的函数
func PostFeedbackHandler(c *gin.Context) {

	feedbackData := model.Feedback{}

	err := c.Bind(&feedbackData)
	if err != nil {
		service.HttpParameterError(c)
		return
	}

	// 读取图片文件并保存到指定路径
	files := c.Request.MultipartForm.File["picture"]
	var picturePaths []string
	for _, file := range files {
		fmt.Println(file.Filename)
		relativePath := filepath.Join("/picture", uuid.NewString()+filepath.Ext(file.Filename))

		f := filex.NewFile(filepath.Join(common.UploadSavePath, relativePath))
		dir := f.ParentFile()
		if !dir.IsExist() {
			err := dir.MkdirAll(0666)
			if err != nil {
				log.Panicln(err)
				return
			}
		}
		destPath := f.Pathname
		if err := c.SaveUploadedFile(file, destPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		relativePath = strings.ReplaceAll(relativePath, "\\", "/")
		picturePaths = append(picturePaths, path.Join(common.UploadDir, relativePath))
	}

	feedbackData.Picture = strings.Join(picturePaths, ",")

	// 将反馈信息插入数据库

	if err := db.Mysql.Create(&feedbackData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回处理结果
	c.JSON(http.StatusOK, model.JsonResponse{
		Code: 200,
		Msg:  "ok",
		Data: nil,
	})

}
