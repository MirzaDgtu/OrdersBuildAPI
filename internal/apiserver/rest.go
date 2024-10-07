package apiserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h HandlerInterface) error {
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
