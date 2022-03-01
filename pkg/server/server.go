package server

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/graphql"
	"github.com/mensatt/mensatt-backend/internal/misc"
)

func Run(cfg *ServerConfig, dbtx db.DBTX) error {
	database := db.New(dbtx)

	r := gin.Default()
	r.Use(cors.Default())

	miscRouterGroup := r.Group(cfg.VersionedPath("/misc"))
	misc.Run(miscRouterGroup)

	gqlRouterGroup := r.Group(cfg.VersionedPath("/graphql"))
	gqlServerParams := graphql.GraphQLParams{
		DebugEnabled: cfg.DebugEnabled,
		Database:     database,
	}
	graphql.Run(gqlRouterGroup, &gqlServerParams)

	log.Println("Running @ " + cfg.SchemaVersionedEndpoint(""))

	return r.Run(cfg.ListenEndpoint()) // Start the server
}
