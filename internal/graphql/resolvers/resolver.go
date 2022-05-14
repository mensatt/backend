package resolvers

import (
	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/graphql/gqlserver"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Database db.Querier
}

func (r *Resolver) OccurrenceInput() gqlserver.OccurrenceInputResolver {
	//TODO implement me
	panic("implement me")
}
