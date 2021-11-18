package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mensatt/mensatt-backend/internal/graphql"
	"github.com/mensatt/mensatt-backend/internal/misc"
)

func Run(cfg *ServerConfig) error {
	r := gin.Default()

	miscRouterGroup := r.Group(cfg.VersionedPath("/misc"))
	misc.Init(miscRouterGroup)

	gqlRouterGroup := r.Group(cfg.VersionedPath("/graphql"))
	gqlServer := graphql.GraphQL{
		DebugEnabled: cfg.DebugEnabled,
	}
	gqlServer.Init(gqlRouterGroup)

	log.Println("Running @ " + cfg.SchemaVersionedEndpoint(""))
	// Run the server
	return r.Run(cfg.ListenEndpoint())
}
