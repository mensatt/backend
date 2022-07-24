package resolvers

import (
	"context"
	"errors"
	"io/ioutil"

	"github.com/mensatt/backend/internal/graphql/models"
)

func (r *Resolver) transcodeAndStoreImage(ctx context.Context, image *models.ImageInput) (string, error) {
	if image.Image.Size > r.ImageProcessor.GetMaxImageSize() {
		return "", errors.New("image size is too big")
	}

	imageBytes, err := ioutil.ReadAll(image.Image.File)
	if err != nil {
		return "", err
	}
	return r.ImageProcessor.StoreImage(imageBytes)
}
