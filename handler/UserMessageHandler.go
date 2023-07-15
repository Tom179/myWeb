package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb02/Database/mysql"
)

type ShowRequest struct {
	Num  int `json:"num"`
	Page int `json:"page"` //第几页
}

func ShowUsers(c *gin.Context) {
	req := ShowRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("err:", err)
	}
	num := req.Num
	page := req.Page
	fmt.Println(num, page)

	users := mysql.GetUserList((page-1)*num, page*num)
	c.JSON(200, users)
}
