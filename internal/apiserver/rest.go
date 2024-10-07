package apiserver

import (
	"fmt"
	"ordersbuild/dblayer"

	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}

	o, errO := dblayer.New()
	if errO != nil {
		return errO
	}
	return RunAPIWithHandler(address, h, o)
}

func RunAPIWithHandler(address string, h HandlerInterface, o OrdersInterface) error {
	//Get gin's default engine
	r := gin.Default()
	r.Use(MyCustomLogger())

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOutUserById)
		//userGroup.GET("/:id/orders", h.GetOrders)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
	}

	orderGroup := r.Group("/order", o.GetOrderByUID)
	{
		orderGroup.GET("/:id")
		orderGroup.POST("/:id/:done", o.GetOrderByUID)
	}

	ordersGroup := r.Group("/orders")
	{
		ordersGroup.POST("", o.AddOrder)
		ordersGroup.GET("", o.GetOrderList)
	}

	//r.Use(static.ServeRoot("/", "../public/build"))
	return r.Run(address)
}

func MyCustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("************************************")
		c.Next()
		fmt.Println("************************************")
	}
}
