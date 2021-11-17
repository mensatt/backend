package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/neoSigfood/neosigfood-backend/internal/graphql/gqlserver"
	"github.com/neoSigfood/neosigfood-backend/internal/graphql/models"
)

func (r *mutationResolver) AddReview(ctx context.Context, review models.ReviewInput) (*models.Review, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns gqlserver.MutationResolver implementation.
func (r *Resolver) Mutation() gqlserver.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
