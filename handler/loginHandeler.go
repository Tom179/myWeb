package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"goweb02/Database/mysql"
	"goweb02/jwt"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) { //email、password
	mysql.Connect()
	loginRequest := LoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请求参数获取失败，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"error":   err.Error(),
		})
	}

	user, err := checkUser(loginRequest.Email, loginRequest.Password)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "登录失败,此账号不存在",
			"error":   err.Error(), //自定义的错误类型，要加上.Error() 因为自定义的类型无法成功完成json转换
		})
	} else {
		jwt := jwt.GetJWT().CreateToken(*user) //登录成功后生成jwt
		c.JSON(200, gin.H{
			"token":     jwt,
			"message":   "登录成功",
			"loginUser": user,
		})
	}
}

func checkUser(email, password string) (*mysql.User, error) {
	user, err := mysql.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("邮件不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}

	return &user, nil
}
