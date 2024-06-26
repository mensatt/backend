package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/internal/database/ent/dish"
	"github.com/mensatt/backend/internal/database/ent/image"
	"github.com/mensatt/backend/internal/database/ent/occurrence"
	"github.com/mensatt/backend/internal/database/ent/review"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
	"github.com/mensatt/backend/internal/graphql/models"
)

// Aliases is the resolver for the aliases field.
func (r *dishResolver) Aliases(ctx context.Context, obj *ent.Dish) ([]*ent.DishAlias, error) {
	return r.Database.Dish.QueryAliases(obj).All(ctx)
}

// ReviewData is the resolver for the reviewData field.
func (r *dishResolver) ReviewData(ctx context.Context, obj *ent.Dish, filter *models.ReviewFilter) (*models.ReviewDataDish, error) {
	queryBuilder := r.Database.Review.Query().Where(review.HasOccurrenceWith(occurrence.HasDishWith(dish.ID(obj.ID))))

	if filter != nil {
		if *filter.Approved {
			queryBuilder.Where(review.AcceptedAtNotNil())
		} else {
			queryBuilder.Where(review.AcceptedAtIsNil())
		}
	}

	reviews, err := queryBuilder.All(ctx)
	if err != nil {
		return nil, err
	}

	images, err := r.Database.Image.Query().Where(image.HasReviewWith(review.HasOccurrenceWith(occurrence.HasDishWith(dish.ID(obj.ID))))).All(ctx)
	if err != nil {
		return nil, err
	}

	averageStars := 0.0
	for _, review := range reviews {
		averageStars += float64(review.Stars)
	}
	if len(reviews) > 0 {
		averageStars /= float64(len(reviews))
	}

	return &models.ReviewDataDish{
		Reviews: reviews,
		Images:  images,
		Metadata: &models.ReviewMetadataDish{
			AverageStars: &averageStars,
			ReviewCount:  len(reviews),
		},
	}, nil
}

// Dish is the resolver for the dish field.
func (r *dishAliasResolver) Dish(ctx context.Context, obj *ent.DishAlias) (*ent.Dish, error) {
	return r.Database.DishAlias.QueryDish(obj).Only(ctx)
}

// AliasName is the resolver for the aliasName field.
func (r *dishAliasResolver) AliasName(ctx context.Context, obj *ent.DishAlias) (string, error) {
	return obj.ID, nil
}

// Review is the resolver for the review field.
func (r *imageResolver) Review(ctx context.Context, obj *ent.Image) (*ent.Review, error) {
	return r.Database.Image.QueryReview(obj).Only(ctx)
}

// Location is the resolver for the location field.
func (r *occurrenceResolver) Location(ctx context.Context, obj *ent.Occurrence) (*ent.Location, error) {
	return r.Database.Occurrence.QueryLocation(obj).Only(ctx)
}

// Dish is the resolver for the dish field.
func (r *occurrenceResolver) Dish(ctx context.Context, obj *ent.Occurrence) (*ent.Dish, error) {
	return r.Database.Occurrence.QueryDish(obj).Only(ctx)
}

// SideDishes is the resolver for the sideDishes field.
func (r *occurrenceResolver) SideDishes(ctx context.Context, obj *ent.Occurrence) ([]*ent.Dish, error) {
	return r.Database.Occurrence.QuerySideDishes(obj).All(ctx)
}

// Tags is the resolver for the tags field.
func (r *occurrenceResolver) Tags(ctx context.Context, obj *ent.Occurrence) ([]*ent.Tag, error) {
	return r.Database.Occurrence.QueryTags(obj).All(ctx)
}

// ReviewData is the resolver for the reviewData field.
func (r *occurrenceResolver) ReviewData(ctx context.Context, obj *ent.Occurrence, filter *models.ReviewFilter) (*models.ReviewDataOccurrence, error) {
	queryBuilder := r.Database.Review.Query().Where(review.HasOccurrenceWith(occurrence.ID(obj.ID)))

	if filter != nil {
		if *filter.Approved {
			queryBuilder.Where(review.AcceptedAtNotNil())
		} else {
			queryBuilder.Where(review.AcceptedAtIsNil())
		}
	}

	reviews, err := queryBuilder.All(ctx)
	if err != nil {
		return nil, err
	}

	images, err := r.Database.Image.Query().Where(image.HasReviewWith(review.HasOccurrenceWith(occurrence.ID(obj.ID)))).All(ctx)
	if err != nil {
		return nil, err
	}

	averageStars := 0.0
	for _, review := range reviews {
		averageStars += float64(review.Stars)
	}
	if len(reviews) > 0 {
		averageStars /= float64(len(reviews))
	}

	return &models.ReviewDataOccurrence{
		Reviews: reviews,
		Images:  images,
		Metadata: &models.ReviewMetadataOccurrence{
			AverageStars: &averageStars,
			ReviewCount:  len(reviews),
		},
	}, nil
}

// Occurrence is the resolver for the occurrence field.
func (r *reviewResolver) Occurrence(ctx context.Context, obj *ent.Review) (*ent.Occurrence, error) {
	return r.Database.Review.QueryOccurrence(obj).Only(ctx)
}

// Images is the resolver for the images field.
func (r *reviewResolver) Images(ctx context.Context, obj *ent.Review) ([]*ent.Image, error) {
	return r.Database.Review.QueryImages(obj).All(ctx)
}

// Key is the resolver for the key field.
func (r *tagResolver) Key(ctx context.Context, obj *ent.Tag) (string, error) {
	return obj.ID, nil
}

// Dish returns gqlserver.DishResolver implementation.
func (r *Resolver) Dish() gqlserver.DishResolver { return &dishResolver{r} }

// DishAlias returns gqlserver.DishAliasResolver implementation.
func (r *Resolver) DishAlias() gqlserver.DishAliasResolver { return &dishAliasResolver{r} }

// Image returns gqlserver.ImageResolver implementation.
func (r *Resolver) Image() gqlserver.ImageResolver { return &imageResolver{r} }

// Occurrence returns gqlserver.OccurrenceResolver implementation.
func (r *Resolver) Occurrence() gqlserver.OccurrenceResolver { return &occurrenceResolver{r} }

// Review returns gqlserver.ReviewResolver implementation.
func (r *Resolver) Review() gqlserver.ReviewResolver { return &reviewResolver{r} }

// Tag returns gqlserver.TagResolver implementation.
func (r *Resolver) Tag() gqlserver.TagResolver { return &tagResolver{r} }

type dishResolver struct{ *Resolver }
type dishAliasResolver struct{ *Resolver }
type imageResolver struct{ *Resolver }
type occurrenceResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
type tagResolver struct{ *Resolver }
