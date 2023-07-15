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

//type GetTopicRequest struct { //查询topic，请设计请求参数
//
//}

func GetOneTopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	fmt.Println(id)
	topic := mysql.GetTopicByID(id)
	c.JSON(200, topic)
}

type GetTopicRequst struct {
	//Keyword string `json:"keyword"` //搜索关键字
	Num int `json:"per_page"` //一页数量
	//PageNum int `json:"page_num"`
}

func GetAllTopics(c *gin.Context) { //请求传入关键字，查询列表
	req := GetTopicRequst{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("查询请求参数绑定错误")
	}

	//fmt.Println("获取参数为：", req.Keyword)
	topics := mysql.GetTopicList("", req.Num, 1) //分页查询第一页
	totalPage := mysql.GetTopicPage(req.Num)
	c.JSON(200, gin.H{
		"":          topics,
		"totalPage": totalPage,
	})

} //返回共有多少页数据

type GetCertainPageRequest struct {
	Perpage int `json:"per_page"`
}

func GetCertainPage(c *gin.Context) {
	req := GetCertainPageRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("请求页数参数绑定失败")
		return
	}

	page, err := strconv.Atoi(c.Param("CertainPage"))
	if err != nil {
		fmt.Println("字符串转换失败")
		return
	}
	topics := mysql.GetTopicList("", req.Perpage, page)
	c.JSON(200, gin.H{
		"msg":    fmt.Sprintf("查询第%d页", page),
		"topics": topics,
	})

}

type MovePageReq struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"per_page"`
}

func MovePage(c *gin.Context) {
	dir := c.Param("dir")
	req := MovePageReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("下一页请求参数绑定失败")
		return
	}
	var topics []mysql.Topic
	var movedPage int
	if dir == "next" {
		movedPage = req.CurrentPage + 1
	} else if dir == "last" {
		movedPage = req.CurrentPage - 1
	}
	topics = mysql.GetTopicList("", req.PerPage, movedPage) //传入每页数、显示第n页
	c.JSON(200, gin.H{
		"msg":    fmt.Sprintf("第%d页", movedPage),
		"topics": topics,
	})
	return
}

type ModifyRequest struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func ModifyTopic(c *gin.Context) {
	req := ModifyRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("修改topic请求参数绑定失败")
		return
	}
	err := mysql.Modify(req.Id, mysql.Topic{}, "description", req.Description)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "修改简介成功",
	})
}
