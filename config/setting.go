package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var ServerPort string
var JWTsecretKey string
var JWTexpireTime int64

func init() { //一个项目可以有多个init函数，在执行main函数之前顺序执行

	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件打开失败:", err)

	}
	LoadData(file)
}

func LoadData(file *ini.File) {
	ServerPort = file.Section("server").Key("httpport").String()
	JWTsecretKey = file.Section("jwt").Key("secretkey").String()
	JWTexpireTime, _ = file.Section("jwt").Key("expireTime").Int64()

	fmt.Println("调试：", ServerPort)

}
