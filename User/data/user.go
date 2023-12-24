package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// username 用户名称
	Username string
	// 密码
	Password string
	// userid
	Identity int64
	// level
	Level string
	// Content
	Content string
	// email
	Email string
	// phone number
	Phone string
	// Job
	Job string
	// Birthday
	Birthday string
	// Tag
	Tag []string
	// head sath the path
	Avatar string
	// describe
	Describe int
	// Collection
	Collection int
}

func (u User) TableName() string {
	return "user"
}
