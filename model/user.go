package model

type User struct {
	ID      int    `gorm:"column:id;primary_key;auto_increment"`
	Name    string `gorm:"column:name;not null"`
}
