package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mensatt/backend/internal/db"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
	"github.com/mensatt/backend/internal/graphql/models"
)

// Images is the resolver for the images field.
func (r *createReviewInputResolver) Images(ctx context.Context, obj *db.CreateReviewWithImagesParams, data []*models.ImageInput) error {
	if len(data) == 0 {
		return nil
	}

	imagesParams := make([]*db.CreateImageParams, len(data))
	for i, image := range data {
		imageStoreID, err := r.transcodeAndStoreImage(ctx, image)
		if err != nil {
			// TODO: delete previous images
			return err
		}
		imagesParams[i] = &db.CreateImageParams{
			ImageStoreID: imageStoreID,
		}
	}
	obj.Images = imagesParams

	return nil
}

// CreateReviewInput returns gqlserver.CreateReviewInputResolver implementation.
func (r *Resolver) CreateReviewInput() gqlserver.CreateReviewInputResolver {
	return &createReviewInputResolver{r}
}

type createReviewInputResolver struct{ *Resolver }
