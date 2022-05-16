package resolvers

import (
	"github.com/mensatt/mensatt-backend/internal/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Database db.Querier
}
