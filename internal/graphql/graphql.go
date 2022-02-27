//go:generate go run github.com/99designs/gqlgen generate

package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/graphql/gqlserver"
	"github.com/mensatt/mensatt-backend/internal/graphql/resolvers"
)

type GraphQLParams struct {
	DebugEnabled bool
	Database     db.Querier
}

func Init(g *gin.RouterGroup, params *GraphQLParams) {
	g.POST("", graphqlHandler(params.Database))
	if params.DebugEnabled {
		g.GET("/playground", playgroundHandler(g.BasePath()))
	}
}

// graphqlHandler defines the GQLGen GraphQL server handler
func graphqlHandler(database db.Querier) gin.HandlerFunc {
	h := handler.NewDefaultServer(
		gqlserver.NewExecutableSchema(
			gqlserver.Config{
				Resolvers: &resolvers.Resolver{
					Database: database,
				},
			},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// playgroundHandler defines a handler to expose the Playground
func playgroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("GraphQL", path)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}