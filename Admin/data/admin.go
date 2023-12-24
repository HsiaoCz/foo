package data

type Admin struct{}

func (a Admin) TableName() string {
	return "admin"
}

