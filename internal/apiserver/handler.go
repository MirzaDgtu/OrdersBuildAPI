package apiserver

import (
	"net/http"
	"ordersbuild/dblayer"
	"ordersbuild/internal/model"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	AddUser(ctx *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (HandlerInterface, error) {
	db, err := dblayer.NewORM("pmp:pmp1226@(172.16.1.25:3306)/ordersbuild")

	if err != nil {
		return nil, err
	} else {

		return &Handler{
			db: db,
		}, err
	}
}

func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = h.db.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
