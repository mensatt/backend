package main

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/getsentry/sentry-go"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mensatt/backend/internal/database/seeds"

	"github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/pkg/server"
	"github.com/mensatt/backend/pkg/utils"
	"log"
	"time"
)

func main() {
	config := server.Config{
		Host:           utils.MustGet("HOST"),
		Port:           utils.MustGetInt32("PORT"),
		MaxConnections: utils.MustGetInt32("MAX_CONNECTIONS"),
		ServiceVersion: utils.MustGet("SERVER_PATH_VERSION"),
		DebugEnabled:   utils.MustGetBool("DEBUG_ENABLED"),
		JWT: utils.JWTKeyStoreConfig{
			PrivateKeyPath: utils.MustGet("JWT_PRIVATE_KEY_PATH"),
			PublicKeyPath:  utils.MustGet("JWT_PUBLIC_KEY_PATH"),
			Algorithm:      utils.MustGet("JWT_ALGORITHM"),
			TimeoutSec:     utils.MustGetInt32("JWT_TIMEOUT_SEC"),
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
		EnableTracing:    true,
		TracesSampleRate: 1.0,
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
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatalln("Error opening database connection:", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln("Error closing database connection:", err)
		}
	}(db)

	db.SetMaxOpenConns(int(config.MaxConnections))

	driver := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(driver))

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
