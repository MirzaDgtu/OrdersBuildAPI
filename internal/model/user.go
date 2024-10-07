package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name"`
	FirstName string `gorm:"column:firstname" json:"firstname" required`
	LastName  string `gorm:"column:lastname" json:"lastname" required`
	Email     string `gorm:"column:email" json:"email" required`
	Pass      string `gorm:"column:pass" json:"pass" required`
	LoggedIn  bool   `gorm:"column:loggedin" json:"loggedin"`
}
