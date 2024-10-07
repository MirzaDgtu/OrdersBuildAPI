package dblayer

import (
	"ordersbuild/internal/database"
	"ordersbuild/internal/model"

	"gorm.io/gorm"
)

type OrderRepo struct {
	Db *gorm.DB
}

func New() (*OrderRepo, error) {
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	return &OrderRepo{Db: db.Db}, nil
}

func (o *OrderRepo) AddOrder(order model.Order) (model.Order, error) {
	err := o.Db.Create(&order).Error
	return order, err
}

func (o OrderRepo) UpdateStatusById(id int, Done bool) error {
	result := o.Db.Table("orders").Where("OrderUID=?", id)
	return result.Update("done", Done).Error
}

func (o *OrderRepo) GetOrderList() (orders []model.Order, err error) {
	return orders, o.Db.Table("orders").Select("*").
		Joins("left join orderdetails d on d.orderuid = orders.orderuid").
		Scan(&orders).Error
}

func (o *OrderRepo) GetOrderByUID(id int) (order model.Order, err error) {
	return order, o.Db.Table("orders").Select("*").
		Joins("left join orderdetails d on d.orderuid = orders.orderuid").
		Where("OrderUID=", id).
		Scan(&order).Error
}
