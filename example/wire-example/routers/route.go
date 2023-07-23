package routers

import (
	"kart-io/kart/example/wire-example/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute(g *gin.Engine, ctr *controller.ApiController) {
	g.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "test")
	})

	// users
	g.GET("/users", ctr.UserController.Users)
}
