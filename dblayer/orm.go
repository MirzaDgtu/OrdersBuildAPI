package dblayer

import (
	"errors"
	"ordersbuild/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(connectionString string) (*DBORM, error) {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) AddUser(user model.User) (model.User, error) {

	hashPassword(&user.Pass)
	user.LoggedIn = true
	err := db.Create(&user).Error
	user.Pass = ""
	return user, err
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}
	//converd password string to byte slice
	sBytes := []byte(*s)
	//Obtain hashed password
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//update password string with the hashed version
	*s = string(hashedBytes[:])
	return nil
}
