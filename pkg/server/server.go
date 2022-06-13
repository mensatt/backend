package server

import (
	"log"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mensatt/backend/internal/assets"
	"github.com/mensatt/backend/internal/db"
	"github.com/mensatt/backend/internal/graphql"
	"github.com/mensatt/backend/internal/middleware"
	"github.com/mensatt/backend/internal/misc"
	"github.com/mensatt/backend/pkg/utils"
)

func Run(config *ServerConfig, pool *pgxpool.Pool) error {
	jwtKeyStore, err := utils.InitJWTKeyStore(&config.JWT)
	if err != nil {
		return err
	}
	
	imageProcessor := utils.NewImageProcessor(config.ImageProcessor)

	database := db.NewExtended(pool)

	app := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	app.Use(cors.New(corsConfig))

	app.Use(sentrygin.New(sentrygin.Options{}))
	app.Use(middleware.Auth(middleware.AuthParams{
		JWTKeyStore: jwtKeyStore,
		Database:    database,
	}))

	assetsRouterGroup := app.Group(config.VersionedPath("/assets"))
	assetsParams := assets.AssetParams{
		AssetDirectory: config.AssetsDir,
	}
	err = assets.Run(assetsRouterGroup, &assetsParams)
	if err != nil {
		return err
	}

	miscRouterGroup := app.Group(config.VersionedPath("/misc"))
	err = misc.Run(miscRouterGroup)
	if err != nil {
		return err
	}

	gqlRouterGroup := app.Group(config.VersionedPath("/graphql"))
	gqlServerParams := graphql.GraphQLParams{
		DebugEnabled:   config.DebugEnabled,
		Database:       database,
		JWTKeyStore:    jwtKeyStore,
		ImageProcessor: imageProcessor,
	}
	err = graphql.Run(gqlRouterGroup, &gqlServerParams)
	if err != nil {
		return err
	}

	log.Println("Running @ " + config.SchemaVersionedEndpoint(""))

	return app.Run(config.ListenEndpoint()) // Start the server
}
