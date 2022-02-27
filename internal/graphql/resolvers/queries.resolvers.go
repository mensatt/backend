package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/graphql/gqlserver"
)

func (r *queryResolver) GetAllergies(ctx context.Context) ([]db.Allergy, error) {
	return r.Database.GetAllergies(ctx)
}

func (r *queryResolver) GetTags(ctx context.Context) ([]db.Tag, error) {
	return r.Database.GetTags(ctx)
}

func (r *queryResolver) GetOccurrences(ctx context.Context) ([]db.Occurrence, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns gqlserver.QueryResolver implementation.
func (r *Resolver) Query() gqlserver.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
