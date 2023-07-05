package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"goweb02/Database/redis"
	"net/http"
	"time"
)

func SendImage(c *gin.Context) { //随机生成6位验证码，设置过期时间储存到redis中
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore) //其实可以直接存储到redis中，这里先暂时不用，手动提取出来手动存储到redis中

	// 生成验证码
	id, b64s, err := captcha.Generate()
	codeValue := captcha.Store.Get(id, true) //获取验证码的值，第二个参数为是否从存储容器中删除。这里会删除，不过是内存中

	redis.Connect()                                                                        //r
	ctx := context.Background()                                                            //e
	if err := redis.RDB.Set(ctx, "图片验证码"+id, codeValue, 1*time.Minute).Err(); err != nil { //d
		fmt.Println(err) //i
		return           //s
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "生成验证码失败",
		})
		return
	}
	// 返回验证码结果
	c.JSON(http.StatusOK, gin.H{
		"captcha_id": id,
		"image_data": b64s,
		//"值":          codeValue,
	})

}
