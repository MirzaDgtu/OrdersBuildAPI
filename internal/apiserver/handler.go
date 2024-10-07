package apiserver

import "ordersbuild/dblayer"

type HandlerInterface struct {
	*dblayer.DBORM
}

func NewHandler() (HandlerInterface, error) {
	db, err := dblayer.NewORM("pmp:pmp1226@(127.0.0.1:3306)/ordersbuild")
	return db, err
}
