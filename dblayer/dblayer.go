package dblayer

import (
	"errors"
	"ordersbuild/internal/model"
)

type DBLayer interface {
	AddUser(model.User) (model.User, error)
	SignInUser(username, password string) (model.User, error)
	SignOutUserById(int) error
}

var ErrINVALIDPASSWORD = errors.New("Invalid password")
