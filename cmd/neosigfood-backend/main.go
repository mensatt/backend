package main

import (
	"log"

	"github.com/neoSigfood/neosigfood-backend/pkg/server"
)

func main() {
	// move to config file or env variables
	sc := server.ServerConfig{
		Host:           "localhost",
		Port:           4000,
		ServiceVersion: "v1",
		DebugEnabled:   true,
	}

	log.Fatal(server.Run(&sc))
}
