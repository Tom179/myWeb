package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"goweb02/Database/redis"
	"net/http"
)

func SendImage(c *gin.Context) { //随机生成6位验证码，设置过期时间储存到redis中
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore) //其实可以直接存储到redis中，这里先暂时不用，手动提取出来手动存储到redis中
	//请不要把他存储在内存当中，因为captcha库是根据存储容器来检测的。如果存储到内存，那么redis中的操作和内容不同步会很麻烦，所以直把存储容器改为redis为好。
	id, b64s, err := captcha.Generate() // 生成验证码
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "生成验证码失败",
		})
		return
	}

	codeValue := captcha.Store.Get(id, false) //获取验证码的值，第二个参数为是否从存储容器中删除，不过是在自定义的存储中。因为我们没有用它的存储容器，所以删不删都没关系
	redis.Create(id, codeValue)               //:redis存储
	// 返回验证码结果
	c.JSON(http.StatusOK, gin.H{
		"captcha_id": id, //redis中的键
		"image_data": b64s,
		//"值":          codeValue,
	})
}
