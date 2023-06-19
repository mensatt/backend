package server

import (
	"github.com/davidbyttow/govips/v2/vips"
	ent "github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/internal/images"
	"log"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/mensatt/backend/internal/graphql"
	"github.com/mensatt/backend/internal/middleware"
	"github.com/mensatt/backend/internal/misc"
	"github.com/mensatt/backend/pkg/imageuploader"
	"github.com/mensatt/backend/pkg/utils"
)

func Run(config *Config, client *ent.Client) error {
	jwtKeyStore, err := utils.InitJWTKeyStore(&config.JWT)
	if err != nil {
		return err
	}

	imageUploader, err := imageuploader.NewImageUploader(config.ImageUploader)
	if err != nil {
		return err
	}
	defer vips.Shutdown() // Shutdown libvips on exit

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

	assetsRouterGroup := app.Group(config.VersionedPath("/assets"))
	imagesRouterGroup := assetsRouterGroup.Group("/images")
	err = images.Run(imagesRouterGroup, &images.ImageParams{
		ImageUploader: imageUploader,
	})
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
		DebugEnabled:  config.DebugEnabled,
		Database:      client,
		JWTKeyStore:   jwtKeyStore,
		ImageUploader: imageUploader,
		ImageBaseURL:  imagesRouterGroup.BasePath(),
	}
	err = graphql.Run(gqlRouterGroup, &gqlServerParams)
	if err != nil {
		return err
	}

	log.Println("Running @ " + config.SchemaVersionedEndpoint(""))

	return app.Run(config.ListenEndpoint()) // Start the server
}
