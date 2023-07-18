package mysql

type Comment struct {
	IdRecord
	Content       string //图片、表情等如何处理
	CreatedByUser int
	TopicID       int
	TimeRecord
	Title string //
}
