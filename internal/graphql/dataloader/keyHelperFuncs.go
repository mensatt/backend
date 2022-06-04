package dataloader

import (
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/db/sqlc"
)

func getReviewKey(review *sqlc.Review) uuid.UUID {
	return review.ID
}
