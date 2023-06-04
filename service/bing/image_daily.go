package bing

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"server_template/model"
	"sync"
	"time"
)

var (
	imageCache  string
	endDateTime time.Time
	mutex       sync.Mutex
)

// 从Bing获取每日一图的链接，并返回链接和图片结束时间
// 从Bing获取每日一图的链接，并返回链接和图片结束时间
func fetchBingImageLink() (string, time.Time, error) {
	// 创建一个新的HTTP请求，设置请求头中的User-Agent字段为浏览器的标识字符串
	req, err := http.NewRequest("GET", "https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", nil)
	if err != nil {
		return "", time.Time{}, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	// 发送HTTP请求，获取Bing每日一图的API响应
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", time.Time{}, err
	}
	defer response.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", time.Time{}, err
	}

	// 解析JSON响应
	var bingResponse model.BingResponse
	if err = json.Unmarshal(body, &bingResponse); err != nil {
		return "", time.Time{}, err
	}

	// 获取图片链接和结束时间
	imageUrl := "https://www.bing.com" + bingResponse.Images[0].Url
	endDate, err := time.ParseInLocation("20060102", bingResponse.Images[0].EndDate, time.UTC)
	if err != nil {
		return "", time.Time{}, err
	}

	return imageUrl, endDate, nil
}

// 获取Bing每日一图的链接，如果缓存未过期，则返回缓存中的链接
func getBingImageLink() (string, error) {
	// 检查缓存是否过期
	mutex.Lock()
	defer mutex.Unlock()
	now := time.Now().UTC()
	if now.Before(endDateTime) {
		return imageCache, nil
	}

	// 缓存过期，从Bing获取链接并更新缓存
	imageLink, endDate, err := fetchBingImageLink()
	if err != nil {
		return "", err
	}

	endOfToday := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, time.UTC)
	if now.After(endOfToday) {
		// 如果今天已经过了图片的过期时间，则再次获取图片链接并更新缓存
		imageLink, endDate, err = fetchBingImageLink()
		if err != nil {
			return "", err
		}
	}

	imageCache = imageLink
	endDateTime = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, time.UTC).Add(24 * time.Hour)

	return imageLink, nil
}

func GetDailyImage(c *gin.Context) {
	// 获取Bing每日一图的链接，如果缓存未过期，则返回缓存中的链接
	//imageLink, err := getBingImageLink()
	//if err != nil {
	//	c.AbortWithError(http.StatusInternalServerError, err)
	//	return
	//}

	imageLink, _, _ := fetchBingImageLink()
	// 将客户端重定向到图片链接
	c.Redirect(http.StatusFound, imageLink)
}
