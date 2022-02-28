package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mensatt/mensatt-backend/pkg/utils"
)

func main() {
	fileToSeed := "./internal/db/sql/seed.sql"

	pool, err := pgxpool.Connect(context.Background(), utils.MustGet("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	defer pool.Close()

	b, err := ioutil.ReadFile(fileToSeed)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	_, err = pool.Exec(context.Background(), string(b))
	if err != nil {
		log.Fatalln("Error writing seed to database:", err)
	}

	log.Println("db successfully seeded.")
}
