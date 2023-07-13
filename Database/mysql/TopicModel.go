package mysql

type Topic struct {
	IdRecord
	Name        string `gorm:""`
	Description string
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

func GetTopic(keyword string) *Topic {
	topic := Topic{}
	DB.Where("name LIKE ?", "%"+keyword+"%").Find(&topic)
	return &topic
}

func GetTopicList(start, end int) []Topic {
	var topics []Topic
	DB.Where("id >? and id <=? ", start, end).Find(&topics)

	return topics
}
