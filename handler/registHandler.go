package handler

import (
	"github.com/gin-gonic/gin"
	"goweb02/Database/mysql"
	"goweb02/Database/mysql/models"
	"net/http"
	"time"
)

type RegistRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Regist(c *gin.Context) {
	mysql.Connect() //连接数据库
	request := RegistRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请求参数获取失败，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"error":   err.Error(),
		})
	}
	//email, username, password := c.PostForm("email"), c.PostForm("username"), c.PostForm("password")
	//👆传统表单提交方式，我们这里用json请求参数方式
	if !mysql.Existed("username", request.Username) && !mysql.Existed("email", request.Email) {
		//fmt.Println("用户名和邮箱都没有使用过")
		userModel := models.User{
			Email:      request.Email,
			Username:   request.Username,
			Password:   request.Password, //加密存储
			TimeRecord: models.TimeRecord{CreateTime: time.Now()},
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
