package apiserver

import "github.com/gin-gonic/gin"

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}

	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h UserRepositoryInterface) error {
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.POST("", h.AddUser)
	}

	return r.Run(address)
}
