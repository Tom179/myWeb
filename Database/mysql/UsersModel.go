package mysql //用户模型和数据库新增记录
import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"goweb02/encrypt"
)

type User struct {
	IdRecord
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

func CreateUser(user *User) { //把该结构体新增到表中
	DB.Create(user)
}

func GetUserByEmail(email string) (User, error) { //查询用户,只能传入email和password？
	user := User{}
	DB.Where("email =? ", email).Find(&user)
	//若未成功查询到记录，返回空的结构体，所以空的ID为0

	if user.ID == 0 {
		return User{}, errors.New("该邮箱不存在")
	}

	return user, nil
}
func GetUserList(start, end int) []User {
	var users []User
	DB.Where("id >? and id <=? ", start, end).Find(&users)

	return users
}

// 查询值为value的field域是否已经被使用
func Existed(myType interface{}, field, value string) bool { //传入要（查询的关键字、值）
	var count int64
	sqlStr := fmt.Sprintf("%s = ?", field)
	DB.Model(myType).Where(sqlStr, value).Count(&count)
	return count > 0
}
