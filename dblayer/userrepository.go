package dblayer

import "ordersbuild/internal/model"

type UserRepository interface {
	AddUser(model.User) (model.User, error)
	//SignInUser(email, password string) (model.User, error)
	//SignOutUserById(int) error
}
