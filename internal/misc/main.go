package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(g *gin.RouterGroup) error {
	// Simple keep-alive/ping handler
	g.GET("/ping", pingHandler())
	return nil
}

// pingHandler is simple keep-alive/ping handler
func pingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
