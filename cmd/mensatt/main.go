package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mensatt/mensatt-backend/pkg/server"
)

func main() {
	// move to config file or env variables
	sc := server.ServerConfig{
		Host:           "localhost",
		Port:           4000,
		ServiceVersion: "v1",
		DebugEnabled:   true,
	}

	pool, err := pgxpool.Connect(context.Background(), "DATABASE_URL")
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	
	log.Fatal(server.Run(&sc, pool))
}
