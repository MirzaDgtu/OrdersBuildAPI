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

func checkPassword(existingHash, incomingPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass)) == nil
}

func (db *DBORM) SignInUser(email, pass string) (user model.User, err error) {
	result := db.Table("users").Where(&model.User{Email: email})
	err = result.First(&user).Error
	if err != nil {
		return user, err
	}

	if !checkPassword(user.Pass, pass) {
		return user, ErrINVALIDPASSWORD
	}

	user.Pass = ""

	err = result.Update("loggedin", 1).Error
	if err != nil {
		return user, err
	}

	return user, result.Find(&user).Error
}

func (db *DBORM) SignOutUserById(id int) error {
	user := model.User{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	return db.Table("users").Where(&user).Update("loggedin", 0).Error
}
