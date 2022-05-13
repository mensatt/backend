package server

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/graphql"
	"github.com/mensatt/mensatt-backend/internal/misc"
)

func Run(cfg *ServerConfig, dbtx db.DBTX) error {

	app := gin.Default()
	app.Use(cors.Default())
	app.Use(sentrygin.New(sentrygin.Options{}))

	miscRouterGroup := app.Group(cfg.VersionedPath("/misc"))
	misc.Run(miscRouterGroup)

	database := db.New(dbtx)

	gqlRouterGroup := app.Group(cfg.VersionedPath("/graphql"))
	gqlServerParams := graphql.GraphQLParams{
		DebugEnabled: cfg.DebugEnabled,
		Database:     database,
	}
	graphql.Run(gqlRouterGroup, &gqlServerParams)

	log.Println("Running @ " + cfg.SchemaVersionedEndpoint(""))

	return app.Run(cfg.ListenEndpoint()) // Start the server
}
