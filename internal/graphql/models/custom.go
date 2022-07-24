package models

import (
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/db/sqlc"
)

type CreateOccurrenceInputHelper struct {
	sqlc.CreateOccurrenceParams
	SideDishes []uuid.UUID `json:"sideDishes"`
	Tags       []string    `json:"tags"`
}
