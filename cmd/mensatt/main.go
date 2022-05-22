package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/mensatt/mensatt-backend/internal/db"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mensatt/mensatt-backend/pkg/server"
	"github.com/mensatt/mensatt-backend/pkg/utils"
)

func main() {
	sc := server.ServerConfig{
		Host:           utils.MustGet("HOST"),
		Port:           utils.MustGetInt32("PORT"),
		ServiceVersion: utils.MustGet("SERVER_PATH_VERSION"),
		DebugEnabled:   utils.MustGetBool("DEBUG_ENABLED"),
	}

	sentryDSN, err := utils.GetOrFile("SENTRY_DSN")
	if err != nil {
		log.Fatalln("Sentry DSN secret could not be retrieved:", err)
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

	log.Fatal(server.Run(&sc, pool))
}
