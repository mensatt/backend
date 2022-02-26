package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/graphql/gqlserver"
	models1 "github.com/mensatt/mensatt-backend/internal/graphql/models"
)

func (r *queryResolver) GetAllergies(ctx context.Context) ([]db.Allergy, error) {
	return r.Database.GetAllergies(ctx)
}

func (r *queryResolver) GetTags(ctx context.Context) ([]models1.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns gqlserver.QueryResolver implementation.
func (r *Resolver) Query() gqlserver.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
