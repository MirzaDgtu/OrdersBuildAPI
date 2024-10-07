package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `gorm:"column:firstname" json:"firstname" required`
	LastName  string    `gorm:"column:lastname" json:"lastname" required`
	Email     string    `gorm:"column:email" json:"email" required`
	Pass      string    `gorm:"column:pass" json:"pass" required`
	LoggedIn  bool      `gorm:"column:loggedin" json:"loggedin"`
	Inn       string    `gorm:"column:inn" json:"inn"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}
