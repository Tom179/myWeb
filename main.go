package main

import (
	"github.com/gin-gonic/gin"
	"goweb02/handler"
)

func main() {
	r := gin.Default()
	setUpRoutes(r)
	r.Run()
}

func setUpRoutes(r *gin.Engine) {
	/*r.LoadHTMLGlob("./Pages/*")                     //加载html
	    r.StaticFile("/style.css", "./Pages/style.css") //加载css

		r.GET("/", func(c *gin.Context) {
			c.JSON(200, "主页面")
		})
		r.GET("/login", func(c *gin.Context) {
			c.HTML(200, "login.html", nil)
		})
		r.GET("/regist", func(c *gin.Context) {
			c.HTML(200, "regist.html", nil)
		})*/

	r.POST("/regist", handler.Regist)
	r.POST("/login", handler.Login)
	//注册登录（没有jwt，之后完善）
	r.POST("/createImageCaptcha", handler.SendImage)

}
