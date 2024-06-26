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
	"os"

	"github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/pkg/server"
	"github.com/mensatt/backend/pkg/utils"
	"log"
	"time"
)

func main() {
	fmt.Println("Starting Mensatt Backend...")
	fmt.Println("Environment:", utils.GetEnvironment())
	fmt.Println("Version:", os.Getenv("VERSION"))

	config := server.Config{
		Host:           utils.MustGet("HOST"),
		Port:           utils.MustGetInt32("PORT"),
		MaxConnections: utils.MustGetInt32("MAX_CONNECTIONS"),
		DebugEnabled:   utils.MustGetBool("DEBUG_ENABLED"),
		JWT: utils.JWTKeyStoreConfig{
			PrivateKeyPath: utils.MustGet("JWT_PRIVATE_KEY_PATH"),
			PublicKeyPath:  utils.MustGet("JWT_PUBLIC_KEY_PATH"),
			Algorithm:      utils.MustGet("JWT_ALGORITHM"),
			TimeoutSec:     utils.MustGetInt32("JWT_TIMEOUT_SEC"),
		},
		ImageAPIURL: utils.MustGet("IMAGE_API_URL"),
		ImageAPIKey: utils.MustGetOrFile("IMAGE_API_KEY"),
	}

	sentryDSN, err := utils.GetOrFile("SENTRY_DSN")
	if err != nil {
		log.Println("Sentry DSN secret could not be retrieved:", err)
	}

	// Initialize sentry.io client for error reporting & logging
	sentryOptions := sentry.ClientOptions{
		Dsn:                sentryDSN,
		Environment:        utils.GetEnvironment(),
		Release:            os.Getenv("VERSION"),
		AttachStacktrace:   true,
		EnableTracing:      true,
		TracesSampleRate:   0.5, // tracks journey of a request
		ProfilesSampleRate: 0.5, // performance profiling
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

	if config.DebugEnabled {
		client = client.Debug() // print sql queries, good enough for now... add logger in the future
	}

	ctx := context.Background()

	// todo: run migration tool here in the future

	// todo: this might delete the database, so be careful!!... uncommented for now!!
	//if err := client.Schema.Create(ctx); err != nil {
	//	log.Fatalln("Error creating schema:", err)
	//}

	if err := seeds.Seed(ctx, client); err != nil {
		log.Fatalln("Error seeding database:", err)
	}

	// Run the server
	log.Fatal(server.Run(&config, client))
}
