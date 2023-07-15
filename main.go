package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb02/Database/mysql"
	"goweb02/Database/redis"
	_ "goweb02/config"
	"goweb02/handler"
	"goweb02/jwt"
)

func main() {

	r := gin.Default()
	redis.Connect()
	mysql.Connect()
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
	r.POST("/createImageCaptcha", handler.SendImage)
	r.POST("/sendEmailCaptcha", handler.SendEmail)
	r.POST("/testAuth", jwt.ParseJWTMiddleWare(), nextMethod)
	r.POST("/getUsers", handler.ShowUsers) //获取用户列表接口：查询第几页，还有bug：未计算总共有几页
	//r.POST("/testINI", getIni)

	r1 := r.Group("/topics") //话题增删查改
	r1.POST("/add", handler.AddTopic)
	r1.DELETE("/delete/:id", handler.DeleteTopic)
	r11 := r1.Group("/get")
	{
		r11.GET("/one/:id", handler.GetOneTopic)
		r11.GET("/all", handler.GetAllTopics) //前端传入每页需要多少
		r11.GET("/all/MovePage/:dir", handler.MovePage)
		r11.GET("/page/:CertainPage", handler.GetCertainPage)
	}
	r1.PUT("/modify", handler.ModifyTopic) //传入id和修改内容

}

func nextMethod(c *gin.Context) {
	for i := 0; i < 5; i++ {
		fmt.Println("进入下一个函数")
	}
}

/*func getIni(c *gin.Context) {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("读取配置文件错误:", err)
	}
	for i := 0; i < 5; i++ {
		fmt.Println("进入测试函数")
	}

	port := cfg.Section("server").Key("httpport")
	fmt.Println(port)
}
*/
