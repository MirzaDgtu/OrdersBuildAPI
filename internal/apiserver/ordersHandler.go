package apiserver

import "github.com/gin-gonic/gin"

type OrdersHandler interface {
	AddOrder(ctx *gin.Context)
	UpdateStatusById(ctx *gin.Context)
	GetOrderList(ctx *gin.Context)
	GetOrderByUID(ctx *gin.Context)
}
