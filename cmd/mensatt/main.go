package main

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/getsentry/sentry-go"
	_ "github.com/lib/pq"
	"github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/internal/database/seeds"
	"github.com/mensatt/backend/pkg/imageuploader"
	"github.com/mensatt/backend/pkg/server"
	"github.com/mensatt/backend/pkg/utils"
	"log"
	"path/filepath"
	"time"
)

func main() {
	assetDir := utils.MustGet("ASSETS_DIR")
	err := utils.CreateDirIfNotExists(assetDir)
	if err != nil {
		log.Fatalf("failed to create asset dir(%s): %s\n", assetDir, err)
	}

	imageDir := filepath.Join(assetDir, "images")
	err = utils.CreateDirIfNotExists(imageDir)
	if err != nil {
		log.Fatalf("failed to create image dir(%s): %s\n", imageDir, err)
	}

	config := server.Config{
		Host:           utils.MustGet("HOST"),
		Port:           utils.MustGetInt32("PORT"),
		ServiceVersion: utils.MustGet("SERVER_PATH_VERSION"),
		DebugEnabled:   utils.MustGetBool("DEBUG_ENABLED"),
		JWT: utils.JWTKeyStoreConfig{
			PrivateKeyPath: utils.MustGet("JWT_PRIVATE_KEY_PATH"),
			PublicKeyPath:  utils.MustGet("JWT_PUBLIC_KEY_PATH"),
			Algorithm:      utils.MustGet("JWT_ALGORITHM"),
			TimeoutSec:     utils.MustGetInt32("JWT_TIMEOUT_SEC"),
		},
		AssetsDir: assetDir,
		ImageUploader: imageuploader.Config{
			ImageDirectory: imageDir,
			MaxImageSizeMB: utils.MustGetInt32("IMAGE_PROCESSOR_MAX_OUTPUT_SIZE_MB"),
			MaxResolution:  utils.MustGetInt32("IMAGE_PROCESSOR_MAX_RESOLUTION"),
		},
	}

	sentryDSN, err := utils.GetOrFile("SENTRY_DSN")
	if err != nil {
		log.Println("Sentry DSN secret could not be retrieved:", err)
	}

	// Initialize sentry.io client for error reporting & logging
	sentryOptions := sentry.ClientOptions{
		Dsn:              sentryDSN,
		AttachStacktrace: true,
	}

	if config.DebugEnabled {
		sentryOptions.Debug = true
	}

	err = sentry.Init(sentryOptions)
	if err != nil {
		log.Fatalln("Sentry initialization failed:", err)
	}
	// Flush buffered events before the program terminates
	defer sentry.Flush(2 * time.Second)

	// Details for database connection
	username, err := utils.GetOrFile("DATABASE_USERNAME")
	if err != nil {
		log.Fatalln("Database username secret could not be retrieved:", err)
	}

	password, err := utils.GetOrFile("DATABASE_PASSWORD")
	if err != nil {
		log.Fatalln("Database password secret could not be retrieved:", err)
	}
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, utils.MustGet("DATABASE_HOST"), utils.MustGet("DATABASE_NAME"))

	// Connect to database
	client, err := ent.Open(dialect.Postgres, databaseUrl)
	if err != nil {
		log.Fatalln("Error opening database connection:", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalln("Error closing database connection:", err)
		}
	}(client)

	ctx := context.Background()

	// todo: run migration tool here in the future
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalln("Error creating schema:", err)
	}

	if err := seeds.Seed(ctx, client); err != nil {
		log.Fatalln("Error seeding database:", err)
	}

	// Run the server
	log.Fatal(server.Run(&config, client))
}
