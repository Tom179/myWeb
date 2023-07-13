package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb02/Database/mysql"
	"strconv"
)

type AddTopicRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func AddTopic(c *gin.Context) {
	req := AddTopicRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("添加话题请求参数绑定失败:", err)
		return
	}

	if mysql.Existed(mysql.Topic{}, "Name", req.Name) {
		fmt.Println("该topic已存在，添加失败")
		c.JSON(400, "topic已经存在，添加失败")
		return
	}

	topic := mysql.Topic{
		Name:        req.Name,
		Description: req.Description,
	}
	fmt.Println("该topic不存在，可以添加")
	mysql.CreateTopic(&topic)
	c.JSON(200, gin.H{
		"msg":   "添加话题成功",
		"topic": topic,
	})
}

func DeleteTopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("字符串转换失败:", err)
		return
	}
	mysql.DeleteTopic(id)
	c.JSON(200, gin.H{
		"msg":    "删除成功",
		"删除话题id": id,
	})
}
