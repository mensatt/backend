package resolvers_generated

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/neoSigfood/neosigfood-backend/internal/graphql/models"
	"github.com/neoSigfood/neosigfood-backend/internal/graphql/server"
)

func (r *queryResolver) GetReviews(ctx context.Context, dishID string) ([]*models.Review, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
