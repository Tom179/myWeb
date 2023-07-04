package mysql

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"goweb02/Database/mysql/models"
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
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		fmt.Println("自动建表失败:", err)
	}

}

func CreateUser(user *models.User) { //把该结构体新增到表中
	DB.Create(user)
}

func GetUserByEmail(email string) (models.User, error) { //查询用户,只能传入email和password？
	user := models.User{}
	DB.Where("email =? ", email).Find(&user)
	//若未成功查询到记录，返回空的结构体，所以空的ID为0

	if user.ID == 0 {
		return models.User{}, errors.New("该邮箱不存在")
	}

	return user, nil
}

// 查询值为value的field域是否已经被使用
func Existed(field, value string) bool { //传入要（查询的关键字、值）
	var count int64
	sqlStr := fmt.Sprintf("%s = ?", field)
	DB.Model(models.User{}).Where(sqlStr, value).Count(&count)
	return count > 0
}
