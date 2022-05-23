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

func Run(cfg *ServerConfig, pool *pgxpool.Pool) error {
	jwtKeyStore, err := utils.InitJWTKeyStore(&utils.JWTKeyStoreConfig{})
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

	miscRouterGroup := app.Group(cfg.VersionedPath("/misc"))
	misc.Run(miscRouterGroup)

	gqlRouterGroup := app.Group(cfg.VersionedPath("/graphql"))
	gqlServerParams := graphql.GraphQLParams{
		DebugEnabled: cfg.DebugEnabled,
		Database:     database,
	}
	graphql.Run(gqlRouterGroup, &gqlServerParams)

	log.Println("Running @ " + cfg.SchemaVersionedEndpoint(""))

	return app.Run(cfg.ListenEndpoint()) // Start the server
}
