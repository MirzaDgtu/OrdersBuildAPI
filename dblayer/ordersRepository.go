package dblayer

import (
	"ordersbuild/internal/model"
)

type OrderRepository interface {
	AddOrder(model.Order) (model.Order, error)
	UpdateStatusById(int, bool) error
	GetOrderList() ([]model.Order, error)
	GetOrderByUID(int) (model.Order, error)
}
