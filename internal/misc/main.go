package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mensatt/mensatt-backend/internal/middleware"
)

func Run(g *gin.RouterGroup) error {
	// Simple keep-alive/ping handler
	g.GET("/ping", pingHandler())
	g.GET("/secure-ping", securePingHandler())
	return nil
}

// pingHandler is simple keep-alive/ping handler
func pingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

// securePingHandler is simple keep-alive/ping handler with auth
func securePingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := middleware.GetUserIDFromCtx(c.Request.Context())
		if user == nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			return
		}

		c.String(http.StatusOK, "OK")
	}
}
