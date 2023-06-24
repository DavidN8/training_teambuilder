package models

type User struct {
	ID       uint   `gorm:"primary key;autoIncrement" json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Teams    []Team `json:"teams"`
}
