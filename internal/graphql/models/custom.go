package models

import (
	"github.com/google/uuid"
	"github.com/mensatt/mensatt-backend/internal/db"
)

type OccurrenceInputHelper struct {
	db.CreateOccurrenceParams
	SideDishes []uuid.UUID `json:"sideDishes"`
	Tags       []string    `json:"tags"`
}
