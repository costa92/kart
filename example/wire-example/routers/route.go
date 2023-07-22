package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute(g *gin.Engine) {
	g.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "test")
	})
}
