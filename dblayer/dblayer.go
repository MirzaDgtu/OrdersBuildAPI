package dblayer

import "ordersbuild/internal/model"

type DBLayer interface {
	AddUser(model.User) (model.User, error)
}
