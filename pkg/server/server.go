package server

import (
	ent "github.com/mensatt/backend/internal/database/ent"
	"log"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/mensatt/backend/internal/graphql"
	"github.com/mensatt/backend/internal/middleware"
	"github.com/mensatt/backend/internal/misc"
	"github.com/mensatt/backend/pkg/utils"
)

func Run(config *Config, client *ent.Client) error {
	jwtKeyStore, err := utils.InitJWTKeyStore(&config.JWT)
	if err != nil {
		return err
	}

	if !config.DebugEnabled {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()

	err = app.SetTrustedProxies(nil)
	if err != nil {
		return err
	}

	app.Use(cors)

	app.Use(sentrygin.New(sentrygin.Options{}))
	app.Use(middleware.Auth(middleware.AuthParams{
		JWTKeyStore: jwtKeyStore,
		Database:    client,
	}))

	miscRouterGroup := app.Group("/data/misc")
	err = misc.Run(miscRouterGroup)
	if err != nil {
		return err
	}

	gqlRouterGroup := app.Group("/data/graphql")
	gqlServerParams := graphql.GraphQLParams{
		DebugEnabled: config.DebugEnabled,
		Database:     client,
		JWTKeyStore:  jwtKeyStore,
		ImageAPIURL:  config.ImageAPIURL,
		ImageAPIKey:  config.ImageAPIKey,
	}
	err = graphql.Run(gqlRouterGroup, &gqlServerParams)
	if err != nil {
		return err
	}

	log.Println("Running @ " + config.SchemaVersionedEndpoint(""))

	return app.Run(config.ListenEndpoint()) // Start the server
}
