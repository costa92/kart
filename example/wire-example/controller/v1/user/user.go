package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

var ProvideUserAPISet = wire.NewSet(NewUserController)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) Users(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "user")
}
