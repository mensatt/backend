package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mensatt/mensatt-backend/pkg/server"
	"github.com/mensatt/mensatt-backend/pkg/utils"
)

func main() {
	// TODO: load from env variables
	sc := mustGetServerConfig()

	pool, err := pgxpool.Connect(context.Background(), utils.MustGet("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	defer pool.Close()

	log.Fatal(server.Run(sc, pool))
}
