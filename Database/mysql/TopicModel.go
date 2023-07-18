package mysql

import (
	"errors"
	"fmt"
)

type Topic struct {
	IdRecord
	Name          string `gorm:"primaryKey"`
	Description   string `gorm:"default:该话题暂无简介"`
	CreatedByUser int    `gorm:"not null"` //创建人？？？？
	Like          int    `gorm:""`         //点赞？？？？？
	TimeRecord
}

func (Topic) TableName() string {
	return "Topics"
}

func CreateTopic(topic *Topic) { //把该结构体新增到表中
	DB.Table(topic.TableName()).Create(topic)
}

func DeleteTopic(id int) { //根据id删除
	DB.Delete(&Topic{}, id)
}

func GetTopicByID(id int) *Topic {
	topic := Topic{}
	DB.Where("id =?", id).Find(&topic)
	return &topic
}

func GetTopicList(keyword string, option ...int) []Topic { //关键字| 1.即一页查询多少 2.第几页
	var topics []Topic
	if len(option) == 0 { //不分页
		DB.Where("name LIKE ?", "%"+keyword+"%").Find(&topics)
		return topics
	}
	if len(option) == 1 { //分页，第一页
		num := option[0]
		DB.Where("name LIKE ?", "%"+keyword+"%").Offset(0).Limit(num).Find(&topics) //查询第一页的记录
		return topics
	}
	if len(option) == 2 { //分页，指定页
		num := option[0]
		DB.Where("name LIKE ?", "%"+keyword+"%").Offset((option[1] - 1) * num).Limit(num).Find(&topics) //模糊查询
		fmt.Println(keyword)
		return topics
	}

	return nil
}

func GetTopicPage(perPage int) int {
	var count int64
	DB.Model(&Topic{}).Count(&count)
	if perPage == 0 {
		return int(count)
	}
	if int(count)%perPage == 0 {
		return int(count) / perPage
	}
	return int(count)/perPage + 1
}

func Modify(id int, table interface{}, field string, value interface{}) error {
	result := DB.Model(&table).Where("id = ?", id).Update(field, value)
	if result.Error != nil {
		fmt.Println("数据库修改失败")
		return result.Error
	} else if result.RowsAffected == 0 {
		fmt.Println("没有匹配的数据库记录")
		return errors.New("没有匹配的数据库记录")
	}

	return nil
}
