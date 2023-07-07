package models

//用于存放表中的公共基础字段，比如自增id和
import "time"

type BaseModel struct {
	ID int `gorm:"column:id;primaryKey;autoIncrement;" ` //自增必须要是主键，，，，原始为uint64
}
type TimeRecord struct {
	CreateTime time.Time `gorm:"column:createTime;" `
	UpDateTime time.Time `gorm:"column:upDateTime;" `
}
