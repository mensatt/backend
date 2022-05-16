package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/graphql/gqlserver"
	"github.com/mensatt/mensatt-backend/internal/graphql/models"
)

func (r *mutationResolver) CreateTag(ctx context.Context, tag db.CreateTagParams) (*db.Tag, error) {
	return r.Database.CreateTag(ctx, &tag)
}

func (r *mutationResolver) CreateDish(ctx context.Context, name string) (*db.Dish, error) {
	return r.Database.CreateDish(ctx, name)
}

func (r *mutationResolver) CreateOccurrence(ctx context.Context, occurrence models.OccurrenceInput) (*db.Occurrence, error) {
	// TODO wrap inside transaction, eliminate for loops, maybe map input to db params better

	occ, err := r.Database.CreateOccurrence(ctx, &db.CreateOccurrenceParams{
		Dish:         occurrence.Dish,
		Date:         occurrence.Date,
		PriceStudent: int32(occurrence.PriceStudent),
		PriceStaff:   int32(occurrence.PriceStaff),
		PriceGuest:   int32(occurrence.PriceGuest),
	})
	if err != nil {
		return nil, err
	}
	for _, t := range occurrence.Tags {
		_, err := r.Database.AddOccurrenceTag(ctx, &db.AddOccurrenceTagParams{
			OccurrenceID: occ.ID,
			TagKey:       t,
		})
		if err != nil {
			return nil, err
		}
	}
	for _, d := range occurrence.SideDishes {
		_, err := r.Database.AddOccurrenceSideDish(ctx, &db.AddOccurrenceSideDishParams{
			OccurrenceID: occ.ID,
			DishID:       d,
		})
		if err != nil {
			return nil, err
		}
	}
	return occ, nil
}

func (r *mutationResolver) CreateReview(ctx context.Context, review db.CreateReviewParams) (*db.Review, error) {
	return r.Database.CreateReview(ctx, &review)
}

// Mutation returns gqlserver.MutationResolver implementation.
func (r *Resolver) Mutation() gqlserver.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
