package data

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagID   int64  `gorm:"column:tag_id;type:int(11);"  json:"tag_id"`
	TagName string `gorm:"column:tag_name;type:varchar(20);" json:"tag_name"`
}

func (t Tag) TableName() string {
	return "tag"
}
