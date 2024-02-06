package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// username 用户名称
	Username string `gorm:"column:username;type:varchar(40);" json:"username"`
	// 密码
	Password string `gorm:"column:password;type:varchar(20);" json:"password"`
	// userid
	Identity int64 `gorm:"column:identity;type:int(11);" json:"identity"`
	// level
	Level int `gorm:"column:level;type:int(2);" json:"level"`
	// Content
	Content string `gorm:"column:content;type:varchar(300);" json:"content"`
	// email
	Email string `gorm:"column:email;type:varchar(40);" json:"email"`
	// phone number
	Phone string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	// Job
	Job string `gorm:"column:job;type:varchar(40);" json:"job"`
	// company
	Company string `gorm:"column:company;type:varchar(60);" json:"company"`
	// Birthday
	Birthday string `gorm:"column:birthday;type:varchar(40);" json:"birthday"`
	// user age
	Age int `gorm:"column:age;type:int(3);" json:"age"`
	// head sath the path
	Avatar string `gorm:"column:avatar;type:varchar(100);" json:"avator"`
	// Tags id
	Tags []int64 `gorm:"column:tags;" json:"tags"`
	// describe
	Describe int `gorm:"column:describe;type:int(6);" json:"describe"`
	// Collection
	Collection int `gorm:"column:collection;type:int(6);" json:"collection"`
}

func (u User) TableName() string {
	return "user"
}
