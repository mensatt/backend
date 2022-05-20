package main

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/mensatt/mensatt-backend/pkg/secrets"
	"log"
	"time"

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

	c, err := secrets.CreateSecretsContainer()
	if err != nil {
		log.Panic(err)
	}

	sentryDSN, err := c.Get(utils.MustGet("SENTRY_DSN_FILE"))
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

	databasePassword, err := c.Get(utils.MustGet("DATABASE_PASSWORD_FILE"))
	if err != nil {
		log.Fatalln("Database password secret could not be retrieved:", err)
	}
	databaseUrl := utils.GetDatabaseHost(databasePassword)

	pool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	defer pool.Close()

	/*
		// Run migrations to keep the database up to date
		err = db.UpgradeDatabase()
		if err != nil {
			log.Fatalln("Error upgrading database:", err)
		}
	*/

	log.Fatal(server.Run(&sc, pool))
}
