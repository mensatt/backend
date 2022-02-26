package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/graphql"
)

func Run(cfg *ServerConfig, pool *pgxpool.Pool) error {
	database := db.New(pool)

	r := gin.Default()

	// miscRouterGroup := r.Group(cfg.VersionedPath("/misc"))
	// misc.Init(miscRouterGroup)

	gqlRouterGroup := r.Group(cfg.VersionedPath("/graphql"))
	gqlServer := graphql.GraphQL{
		DebugEnabled: cfg.DebugEnabled,
		Database:     database,
	}
	gqlServer.Init(gqlRouterGroup)

	log.Println("Running @ " + cfg.SchemaVersionedEndpoint(""))

	return r.Run(cfg.ListenEndpoint()) // Run the server
}
