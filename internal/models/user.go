package models

type User struct {
	ID   int64  `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"uniqueIndex;not null"`
}

// TableName specifies the table name
func (User) TableName() string {
	return "user"
}
