//go:generate go run github.com/99designs/gqlgen generate

package graphql

import (
	ent "github.com/mensatt/backend/internal/database/ent"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/mensatt/backend/internal/graphql/directives"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
	"github.com/mensatt/backend/internal/graphql/resolvers"
	"github.com/mensatt/backend/pkg/utils"
)

type GraphQLParams struct {
	DebugEnabled bool
	Database     *ent.Client
	JWTKeyStore  *utils.JWTKeyStore
	ImageAPIURL  string
	ImageAPIKey  string
}

func Run(g *gin.RouterGroup, params *GraphQLParams) error {
	g.Any("", graphqlHandler(params))
	if params.DebugEnabled {
		g.GET("/playground", playgroundHandler(g.BasePath()))
	}
	return nil
}

// graphqlHandler defines the GQLGen GraphQL server handler
func graphqlHandler(params *GraphQLParams) gin.HandlerFunc {
	var vscBuildInfo *utils.VCSBuildInfo

	// if debug is not enabled, hide vsc build info
	if params.DebugEnabled {
		var err error
		vscBuildInfo, err = utils.GetVscBuildInfo()
		if err != nil {
			log.Println("Error: VCS build info could not be retrieved:", err)
		}
	}

	h := handler.NewDefaultServer(
		gqlserver.NewExecutableSchema(
			gqlserver.Config{
				Resolvers: &resolvers.Resolver{
					Database:               params.Database,
					JWTKeyStore:            params.JWTKeyStore,
					VCSBuildInfo:           vscBuildInfo,
					ImageAPIURL:            params.ImageAPIURL,
					ImageAPIKey:            params.ImageAPIKey,
					ReviewAcceptedChannels: map[string]chan *ent.Review{},
					ReviewCreatedChannels:  map[string]chan *ent.Review{},
				},
				Directives: gqlserver.DirectiveRoot{
					Authenticated: directives.Authenticated,
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
