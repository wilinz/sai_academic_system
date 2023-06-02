package service

import (
	"github.com/gin-gonic/gin"
	"server_template/model"
)

func HttpParameterError(c *gin.Context) {
	c.JSON(200, model.JsonResponse{
		Code: 403,
		Msg:  "参数错误",
		Data: nil,
	})
}

func HttpServerInternalError(c *gin.Context) {
	c.JSON(200, model.JsonResponse{
		Code: 500,
		Msg:  "Http服务器内部错误",
		Data: nil,
	})
}

func HttpNotFound(c *gin.Context) {
	c.JSON(404, model.JsonResponse{
		Code: 404,
		Msg:  "Record not found",
		Data: nil,
	})
}

func HttpOK(c *gin.Context) {
	c.JSON(200, model.JsonResponse{
		Code: 200,
		Msg:  "ok",
		Data: nil,
	})
}

func HttpOK1(c *gin.Context, data any) {
	c.JSON(200, model.JsonResponse{
		Code: 200,
		Msg:  "ok",
		Data: data,
	})
}
