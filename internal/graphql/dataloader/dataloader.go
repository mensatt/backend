package dataloader

import (
	"context"

	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/mensatt/backend/internal/db"
	"github.com/mensatt/backend/internal/db/sqlc"
)

// DataLoader offers data loaders scoped to a context
type DataLoader struct {
	reviewLoader *dataloader.Loader[uuid.UUID, *sqlc.Review]
}

// NewDataLoader returns the instantiated Loaders struct for use in a request
func NewDataLoader(db db.ExtendedQuerier) *DataLoader {
	return &DataLoader{
		reviewLoader: newBatchedLoaderFromDB(getReviewKey, db.GetBulkReviewsByID),
	}
}

// GetReview wraps the Review dataloader for efficient retrieval by Review ID
func (i *DataLoader) GetReview(ctx context.Context, id uuid.UUID) (*sqlc.Review, error) {
	thunk := i.reviewLoader.Load(ctx, id)
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result, nil
}
