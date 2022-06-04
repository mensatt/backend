package resolvers

import (
	"github.com/mensatt/backend/internal/db"
	"github.com/mensatt/backend/pkg/utils"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Database     db.ExtendedQuerier
	JWTKeyStore  *utils.JWTKeyStore
	VCSBuildInfo *utils.VCSBuildInfo
	ImageProcessor *utils.ImageProcessor
}
