package handler

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"gopkg.in/gomail.v2"
	"goweb02/Database/redis"
	"math/rand"
	"time"
)

type sendEmailRequest struct {
	CaptchaID  string `json:"captcha_id"` //凡是要绑定json必须要大写
	CaptchaTry string `json:"captchaTry"`
	Email      string `json:"email"`
}

func SendEmail(c *gin.Context) {
	req := sendEmailRequest{}
	c.ShouldBindJSON(&req)
	//验证图片验证码
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	redis.Connect()
	store := redis.RedisStore{}
	captcha := base64Captcha.NewCaptcha(driver, &store)

	if !captcha.Verify(req.CaptchaID, req.CaptchaTry, true) { //第三个参数，设定验证码是否未一次性的，如果为true，验证一次就作废验证码(不管是否对错)
		c.JSON(400, gin.H{
			"msg":    "图片验证码不正确",
			"id":     req.CaptchaID,
			"answer": req.CaptchaTry,
		})
		return
	}
	//发送邮件验证码

	verifyCode := RandomVerifyCode()
	if !Send("1537946665@qq.com", req.Email, "验证码", verifyCode) {
		fmt.Println("验证码发送失败")
		return
	}
	redis.Create(req.Email, verifyCode)
	c.JSON(200, gin.H{
		"message": "验证码发送成功",
		"验证码":  verifyCode,
	})
}

func Send(senderMail string, reciverMail string, subject string, content string) bool { //返回是否发送成
	m := gomail.NewMessage()
	m.SetHeader("From", senderMail)  // 发件人
	m.SetHeader("To", reciverMail)   // 收件人
	m.SetHeader("Subject", subject)  // 主题
	m.SetBody("text/plain", content) //发送格式

	d := gomail.NewDialer("smtp.qq.com", 587, senderMail, "cpygzacjajawfghc") //smtp服务器地址、端口、发件人邮箱地址、发件人smtp授权码
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}                       // true为关闭tls认证，跳过证书验证
	if err := d.DialAndSend(m); err != nil {                                  // 发送邮件
		fmt.Println(fmt.Sprintf("发送邮件失败: %v", err))
		return false
	}
	return true
}

func RandomVerifyCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
