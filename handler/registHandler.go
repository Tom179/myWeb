package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"goweb02/Database/mysql"
	"goweb02/Database/redis"
	"net/http"
	"time"
)

type RegistRequest struct {
	EmailCaptchaTry string `json:"emailCaptchaTry"` //邮件验证码，填好了才能注册
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
}

func Regist(c *gin.Context) {
	mysql.Connect() //连接数据库
	redis.Connect()
	req := RegistRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请求参数获取失败，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"error":   err.Error(),
		})
	}
	if err := verifyCaptcha(req); err != nil { //验证并删除redis中的邮件记录
		c.JSON(400, gin.H{
			"message": "注册失败",
			"error":   errors.New("邮件验证码错误").Error(),
		})
		return
	}
	//email, username, password := c.PostForm("email"), c.PostForm("username"), c.PostForm("password")
	//👆传统表单提交方式，我们这里用json请求参数方式
	if !mysql.Existed(mysql.User{}, "username", req.Username) /*&& !mysql.Existed("email", req.Email) */ {
		//fmt.Println("用户名和邮箱都没有使用过")
		userModel := mysql.User{
			Email:      req.Email,
			Username:   req.Username,
			Password:   req.Password, //加密存储
			TimeRecord: mysql.TimeRecord{CreateTime: time.Now()},
		}
		mysql.CreateUser(&userModel)
		c.JSON(200, gin.H{
			"message":      "注册成功",
			"registedUser": userModel,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "注册失败",
			"err":     "该邮件或用户名已被注册",
		})
	}
}

// 问题1；前端如何将请求数据发送给后端？html表单？还是全部携带在请求中?
// 前后端分离架构下，前端是可以把表单中的数据打包在请求中的，所以可以在请求惨始终获取

func verifyCaptcha(req RegistRequest) error {
	ctx := context.Background()
	val, err := redis.RDB.Get(ctx, req.Email).Result()
	if err != nil {
		return err
	}

	if val != req.EmailCaptchaTry {
		return errors.New("邮件验证码错误")
	}

	//redis.Delete(req.Email)//👈验证完后应该删除
	return nil
}
