package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mensatt/mensatt-backend/internal/db/sqlc"
	"github.com/mensatt/mensatt-backend/internal/graphql/gqlserver"
	"github.com/mensatt/mensatt-backend/internal/graphql/models"
	"github.com/mensatt/mensatt-backend/pkg/utils"
)

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := r.Database.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", fmt.Errorf("wrong password")
	}

	return r.JWTKeyStore.GenerateJWT(user.ID)
}

func (r *mutationResolver) CreateTag(ctx context.Context, tag sqlc.CreateTagParams) (*sqlc.Tag, error) {
	return r.Database.CreateTag(ctx, &tag)
}

func (r *mutationResolver) CreateDish(ctx context.Context, name string) (*sqlc.Dish, error) {
	return r.Database.CreateDish(ctx, name)
}

func (r *mutationResolver) CreateOccurrence(ctx context.Context, occurrence models.OccurrenceInputHelper) (*sqlc.Occurrence, error) {
	return r.Database.CreateOccurrenceWithSideDishesAndTags(ctx, &occurrence.CreateOccurrenceParams, occurrence.SideDishes, occurrence.Tags)
}

func (r *mutationResolver) CreateReview(ctx context.Context, review sqlc.CreateReviewParams) (*sqlc.Review, error) {
	return r.Database.CreateReview(ctx, &review)
}

// Mutation returns gqlserver.MutationResolver implementation.
func (r *Resolver) Mutation() gqlserver.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
