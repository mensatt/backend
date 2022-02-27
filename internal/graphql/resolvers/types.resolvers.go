package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/graphql/gqlserver"
)

func (r *imageResolver) Occurrence(ctx context.Context, obj *db.Image) (*db.Occurrence, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *occurrenceResolver) Dish(ctx context.Context, obj *db.Occurrence) (*db.Dish, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *occurrenceResolver) SideDishes(ctx context.Context, obj *db.Occurrence) ([]db.Dish, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *occurrenceResolver) Allergies(ctx context.Context, obj *db.Occurrence) ([]db.Allergy, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *occurrenceResolver) Tags(ctx context.Context, obj *db.Occurrence) ([]db.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *reviewResolver) Occurrence(ctx context.Context, obj *db.Review) (*db.Occurrence, error) {
	panic(fmt.Errorf("not implemented"))
}

// Image returns gqlserver.ImageResolver implementation.
func (r *Resolver) Image() gqlserver.ImageResolver { return &imageResolver{r} }

// Occurrence returns gqlserver.OccurrenceResolver implementation.
func (r *Resolver) Occurrence() gqlserver.OccurrenceResolver { return &occurrenceResolver{r} }

// Review returns gqlserver.ReviewResolver implementation.
func (r *Resolver) Review() gqlserver.ReviewResolver { return &reviewResolver{r} }

type imageResolver struct{ *Resolver }
type occurrenceResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
