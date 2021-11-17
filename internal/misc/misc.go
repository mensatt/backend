package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(g *gin.RouterGroup) {
	g.GET("/ping", pingHandler())
}

// pingHandler is simple keep-alive/ping handler
func pingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
}
