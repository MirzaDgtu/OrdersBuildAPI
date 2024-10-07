package dblayer

import (
	"net/http"
	"ordersbuild/dblayer"
	"ordersbuild/internal/model"

	"github.com/gin-gonic/gin"
)

type OrderRepository interface {
	AddOrder(model.Order) (model.Order, error)
	UpdateStatusById(int, bool) error
	GetOrderList() ([]model.Order, error)
	GetOrderByUID(int) (model.Order, error)
}

type OrdersHandler struct {
	db dblayer.OrderRepository
}

func (h *OrdersHandler) AddOrder(ctx *gin.Context) {
	if h.db == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}

	var orders []model.Order
	err := ctx.ShouldBindJSON(&orders)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createOrders []model.Order

	for order := range orders {

		orderCr, err := h.db.AddOrder(order)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		createOrders = append(createOrders, orderCr)
	}

	ctx.JSON(http.StatusOK, createOrders)
}

func (h *OrdersHandler) GetOrderList(ctx *gin.Context) {
	if h.db == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}

	orders, err := h.db.GetOrderList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
