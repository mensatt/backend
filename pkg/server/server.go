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
	"github.com/mensatt/mensatt-backend/pkg/utils"
)

func Run(config *ServerConfig, pool *pgxpool.Pool) error {
	jwtKeyStore, err := utils.InitJWTKeyStore(&config.JWT)
	if err != nil {
		return err
	}

	database := db.NewExtended(pool)

	app := gin.Default()
	app.Use(cors.Default())
	app.Use(sentrygin.New(sentrygin.Options{}))
	app.Use(middleware.Auth(middleware.AuthParams{
		JWTKeyStore: jwtKeyStore,
		Database:    database,
	}))

	miscRouterGroup := app.Group(config.VersionedPath("/misc"))
	misc.Run(miscRouterGroup)

	gqlRouterGroup := app.Group(config.VersionedPath("/graphql"))
	gqlServerParams := graphql.GraphQLParams{
		DebugEnabled: config.DebugEnabled,
		Database:     database,
		JWTKeyStore:  jwtKeyStore,
	}
	graphql.Run(gqlRouterGroup, &gqlServerParams)

	log.Println("Running @ " + config.SchemaVersionedEndpoint(""))

	return app.Run(config.ListenEndpoint()) // Start the server
}
