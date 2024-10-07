package apiserver

import (
	"net/http"
	"ordersbuild/dblayer"
	"ordersbuild/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	AddUser(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	SignOutUserById(ctx *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (HandlerInterface, error) {
	db, err := dblayer.NewORM("pmp:pmp1226@(172.16.1.25:3306)/ordersbuild?parseTime=true")

	if err != nil {
		return nil, err
	} else {

		return &Handler{
			db: db,
		}, err
	}
}

func (h *Handler) AddUser(ctx *gin.Context) {
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
	user, err = h.db.AddUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) SignIn(ctx *gin.Context) {
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

	user, err = h.db.SignInUser(user.Email, user.Pass)
	if err != nil {
		if err == dblayer.ErrINVALIDPASSWORD {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) SignOutUserById(ctx *gin.Context) {
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

	err = h.db.SignOutUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
