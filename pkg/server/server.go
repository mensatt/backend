package server

import (
	"log"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/graphql"
	"github.com/mensatt/mensatt-backend/internal/middleware"
	"github.com/mensatt/mensatt-backend/internal/misc"
)

func Run(cfg *ServerConfig, pool *pgxpool.Pool) error {

	app := gin.Default()
	app.Use(cors.Default())
	app.Use(sentrygin.New(sentrygin.Options{}))
	app.Use(middleware.Auth(middleware.AuthParams{
		CookieName: "token",
	}))

	miscRouterGroup := app.Group(cfg.VersionedPath("/misc"))
	misc.Run(miscRouterGroup)

	database := db.NewExtended(pool)

	gqlRouterGroup := app.Group(cfg.VersionedPath("/graphql"))
	gqlServerParams := graphql.GraphQLParams{
		DebugEnabled: cfg.DebugEnabled,
		Database:     database,
	}
	graphql.Run(gqlRouterGroup, &gqlServerParams)

	log.Println("Running @ " + cfg.SchemaVersionedEndpoint(""))

	return app.Run(cfg.ListenEndpoint()) // Start the server
}
