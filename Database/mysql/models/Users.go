package models //用户模型和数据库新增记录
import (
	"gorm.io/gorm"
	"goweb02/encrypt"
)

type User struct {
	BaseModel
	Username string `gorm:"column:username;primary_key"`
	Email    string `gorm:"column:email;primary_key"`
	Password string `gorm:"column:password"`
	TimeRecord
}

func (User) TableName() string {
	return "userLogin"
}

func (this *User) BeforeSave(tx *gorm.DB) (err error) { //gorm自带的模型钩子，在创建和更新模型Create、Update、Save前被自动调用
	// 因为我们使用的是模型钩子，所以原有的注册逻辑不需要修改
	this.Password, err = encrypt.EncryptPassword(this.Password)
	return
}
