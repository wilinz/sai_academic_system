package account

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server_template/common"
	"server_template/constant/error_code"
	"server_template/constant/redis_prefix"
	"server_template/constant/user_constant"
	"server_template/constant/verification_code"
	"server_template/db"
	"server_template/model"
	"server_template/util"
	"time"
)

func IsPasswordCorrect(c *gin.Context, username string, password string) bool {
	//如果密码错误次数超过限制，则禁止登录一段时间
	key := util.GetKey(redis_prefix.PasswordTryCount, username)
	errorCount, _ := db.Redis.Get(db.Context, key).Int()
	if errorCount >= user_constant.PasswordTryMax {
		ttl := db.Redis.TTL(db.Context, key).Val()
		c.JSON(200, model.JsonResponse{
			Code: error_code.RequestTooFrequent,
			Msg:  "请求过于频繁",
			Data: ttl / time.Second,
		})
		return false
	}

	//验证密码
	var dbUser = model.User{}
	err := db.Mysql.Select("password", "salt").Where("username=?", username).Take(&dbUser).Error
	//如果不存在
	if err != nil {
		c.JSON(200, model.JsonResponse{
			Code: error_code.UserNotExist,
			Msg:  "账号不存在",
			Data: nil,
		})
		return false
	}
	encryption := util.Sha256Sum(password + dbUser.Salt)
	//密码正确
	if encryption == dbUser.Password {
		db.Redis.Del(db.Context, key)
		return true
	}

	//获取ttl
	var ttl = time.Hour * 1
	if errorCount > 0 {
		if ttlCmd := db.Redis.TTL(db.Context, key); ttlCmd.Err() == nil {
			ttl = ttlCmd.Val()
		}
	}

	//更新计数
	db.Redis.Set(db.Context, key, errorCount+1, ttl)

	c.JSON(200, model.JsonResponse{
		Code: error_code.PasswordError,
		Msg:  "密码错误",
		Data: nil,
	})
	return false
}

func IsAccountExists(c *gin.Context, username string) bool {
	var count int64
	db.Mysql.Model(&model.User{}).Where(map[string]any{"username": username}).Count(&count)
	if count > 0 {
		c.JSON(200, model.JsonResponse{
			Code: error_code.UserAlreadyExists,
			Msg:  "账号已注册",
			Data: nil,
		})
		return true
	}
	return false
}

func IsVerificationCodeCorrect(c *gin.Context, code, codeType, username string) bool {
	codeKey := util.GetKey(redis_prefix.Code, codeType, username)
	//如果验证码错误次数超过限制，则把验证码删除，以重新获取
	TryCountKey := util.GetKey(redis_prefix.CodeTryCount, codeKey)
	errorCount, _ := db.Redis.Get(db.Context, TryCountKey).Int()
	if errorCount >= verification_code.TryMax {
		db.Redis.Del(db.Context, codeKey)
		db.Redis.Del(db.Context, TryCountKey)
		c.JSON(200, model.JsonResponse{
			Code: error_code.PleaseReAcquire,
			Msg:  "请重新获取",
			Data: nil,
		})
		return false
	}

	//从数据库中取出验证码
	codeCmd := db.Redis.Get(db.Context, codeKey)
	log.Println(codeCmd)
	//不存在
	if err := codeCmd.Err(); err != nil {
		log.Println(err)
		log.Println("验证码不存在")
		c.JSON(200, model.JsonResponse{
			Code: error_code.Unverified,
			Msg:  "验证码错误",
			Data: nil,
		})
		return false
	}

	dbCode := codeCmd.Val()
	log.Println(codeCmd)
	//正确
	if dbCode == code {
		db.Redis.Del(db.Context, codeKey)
		db.Redis.Del(db.Context, TryCountKey)
		return true
	}

	//获取ttl
	var ttl = verification_code.Interval
	if errorCount > 0 {
		if ttlCmd := db.Redis.TTL(db.Context, codeKey); ttlCmd.Err() == nil {
			ttl = ttlCmd.Val()
		}
	}
	//更新计数
	db.Redis.Set(db.Context, TryCountKey, errorCount+1, ttl)

	log.Println("验证码错误")
	c.JSON(200, model.JsonResponse{
		Code: error_code.Unverified,
		Msg:  "Unverified",
		Data: nil,
	})
	return false
}

func IsLogged(c *gin.Context) (bool, string) {
	logged, username, err := isLogged(c)
	if !logged || err != nil {
		c.JSON(http.StatusUnauthorized, model.JsonResponse{
			Code: error_code.Unverified,
			Msg:  "请登录",
			Data: nil,
		})
	}
	return logged, username
}

func isLogged(c *gin.Context) (bool, string, error) {
	session, err := common.Sessions.Get(c.Request, "session-key")
	var username, exist = session.Values["username"].(string)
	return exist, username, err
}
