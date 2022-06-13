package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/mensatt/backend/internal/db"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mensatt/backend/pkg/server"
	"github.com/mensatt/backend/pkg/utils"
)

func main() {
	assetDir := utils.MustGet("ASSETS_DIR")
	imageDir := filepath.Join(assetDir, "images")
	err := utils.CreateDirIfNotExists(imageDir)
	if err != nil {
		log.Fatalf("failed to create image dir(%s): %s\n", imageDir, err)
	}

	config := server.ServerConfig{
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
		ImageProcessor: utils.ImageProcessorConfig{
			ImageDirectory:  imageDir,
			MaxOutputSizeMB: utils.MustGetInt32("IMAGE_PROCESSOR_MAX_OUTPUT_SIZE_MB"),
			MaxResolution:   utils.MustGetInt32("IMAGE_PROCESSOR_MAX_RESOLUTION"),
		},
	}

	sentryDSN, err := utils.GetOrFile("SENTRY_DSN")
	if err != nil {
		log.Println("Sentry DSN secret could not be retrieved:", err)
	}

	// Initialize sentry.io client for error reporting & logging
	err = sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
	})
	if err != nil {
		log.Fatalln("Sentry initialization failed:", err)
	}
	// Flush buffered events before the program terminates
	defer sentry.Flush(2 * time.Second)

	username, err := utils.GetOrFile("DATABASE_USERNAME")
	if err != nil {
		log.Fatalln("Database username secret could not be retrieved:", err)
	}

	password, err := utils.GetOrFile("DATABASE_PASSWORD")
	if err != nil {
		log.Fatalln("Database password secret could not be retrieved:", err)
	}
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, utils.MustGet("DATABASE_HOST"), utils.MustGet("DATABASE_NAME"))

	pool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	defer pool.Close()

	// Run migrations to keep the database up to date
	err = db.UpgradeDatabase(databaseUrl)
	if err != nil {
		log.Fatalln("Error upgrading database:", err)
	}

	log.Fatal(server.Run(&config, pool))
}
