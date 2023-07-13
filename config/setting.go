package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func init() { //一个项目可以有多个init函数，在执行main函数之前顺序执行

	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件打开失败:", err)

	}
	LoadData(file)
}

func LoadData(file *ini.File) {
	HttpPort := file.Section("server").Key("httpport").String()
	for i := 0; i < 5; i++ {
		fmt.Println("调试：", HttpPort)
	}

}
