//go:generate go run github.com/99designs/gqlgen generate

package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/neoSigfood/neosigfood-backend/internal/graphql/gqlserver"
	"github.com/neoSigfood/neosigfood-backend/internal/graphql/resolvers"
)

type GraphQL struct {
	DebugEnabled bool
}

func (gql *GraphQL) Init(g *gin.RouterGroup) {
	g.POST("", gql.graphqlHandler())
	if gql.DebugEnabled {
		g.GET("/playground", gql.playgroundHandler(g.BasePath()))
	}
}

// graphqlHandler defines the GQLGen GraphQL server handler
func (gql *GraphQL) graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(
		gqlserver.NewExecutableSchema(
			gqlserver.Config{
				Resolvers: &resolvers.Resolver{},
			},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// playgroundHandler defines a handler to expose the Playground
func (gql *GraphQL) playgroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("GraphQL", path)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
