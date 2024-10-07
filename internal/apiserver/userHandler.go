package apiserver

import (
	"net/http"
	"ordersbuild/dblayer"
	"ordersbuild/internal/model"

	"github.com/gin-gonic/gin"
)

type UserRepositoryInterface interface {
	AddUser(ctx *gin.Context)
}

type UserRepository struct {
	db *dblayer.DBORM
}

func (h *UserRepository) AddUser(ctx *gin.Context) {
	if h.db == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}

	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
