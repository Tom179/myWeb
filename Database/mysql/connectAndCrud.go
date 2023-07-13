package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //连接在Database类中，作为成员变量，但是需要connect赋值

func Connect() { //连接的库、并自动建表
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		"root",
		"Qy85891607",
		"127.0.0.1",
		"3306",
		"MyGoweb", //库名，gorm不能自动创库，只能自动创表。
		"utf8mb4")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库打开失败")
	}
	DB = db
	if err := DB.AutoMigrate(&User{}); err != nil {
		fmt.Println("自动建User表失败:", err)
	}
	if err := DB.AutoMigrate(&Topic{}); err != nil {
		fmt.Println("自动建Topic表失败:", err)
	}
}
