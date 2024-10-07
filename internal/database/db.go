package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	Db *gorm.DB
}

func InitDB() (s *Store, err error) {
	if s.Db != nil {
		return &Store{}, nil
	} else {
		db, err := gorm.Open(mysql.Open("pmp:pmp1226@(172.16.1.25:3306)/ordersbuild?parseTime=true"), &gorm.Config{})
		if err != nil {
			return nil, err
		} else {
			return &Store{
				Db: db,
			}, nil
		}
	}
}
