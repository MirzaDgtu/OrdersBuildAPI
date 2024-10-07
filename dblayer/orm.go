package dblayer

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBORM struct {
	gorm.DB
}

func NewORM(connectionString string) (*DBORM, error) {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	return &DBORM{
		DB: *db,
	}, err
}
