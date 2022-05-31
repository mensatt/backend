package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/db/sqlc"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
)

func (r *editReviewInputResolver) Dish(ctx context.Context, obj *sqlc.EditReviewParams, data uuid.UUID) error {
	panic(fmt.Errorf("not implemented"))
}

// EditReviewInput returns gqlserver.EditReviewInputResolver implementation.
func (r *Resolver) EditReviewInput() gqlserver.EditReviewInputResolver {
	return &editReviewInputResolver{r}
}

type editReviewInputResolver struct{ *Resolver }
