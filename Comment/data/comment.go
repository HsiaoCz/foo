package data

type Comment struct{}

func (c Comment) TableName() string {
	return "comment"
}
