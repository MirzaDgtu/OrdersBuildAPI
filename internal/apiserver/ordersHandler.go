package apiserver

import (
	"net/http"
	"ordersbuild/dblayer"
	"ordersbuild/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrdersInterface interface {
	AddOrder(ctx *gin.Context)
	UpdateStatusById(ctx *gin.Context)
	GetOrderList(ctx *gin.Context)
	GetOrderByUID(ctx *gin.Context)
}

type OrdersHand struct {
	db dblayer.OrderRepository
}

func (h *OrdersHand) AddOrder(ctx *gin.Context) {
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

	for _, order := range orders {

		orderCr, err := h.db.AddOrder(order)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		createOrders = append(createOrders, orderCr)
	}

	ctx.JSON(http.StatusOK, createOrders)
}

func (h *OrdersHand) GetOrderList(ctx *gin.Context) {
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

func (h *OrdersHand) GetOrderByUID(ctx *gin.Context) {
	if h.db == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}

	p := ctx.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, errO := h.db.GetOrderByUID(id)
	if errO != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (h *OrdersHand) UpdateStatusById(ctx *gin.Context) {
	if h.db == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}

	pId := ctx.Param("id")
	pDone := ctx.Param("done")

	id, err := strconv.Atoi(pId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	done, errD := strconv.ParseBool(pDone)
	if errD != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errU := h.db.UpdateStatusById(id, done)
	if errU != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Статус документа изменен"})
}
