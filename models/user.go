package models

type User struct {
	ID   int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (t *User) TableName() string {
	return "user"
}
