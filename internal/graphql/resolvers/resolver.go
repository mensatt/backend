package resolvers

import (
	ent "github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/pkg/imageuploader"
	"github.com/mensatt/backend/pkg/utils"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Database      *ent.Client
	JWTKeyStore   *utils.JWTKeyStore
	VCSBuildInfo  *utils.VCSBuildInfo
	ImageUploader *imageuploader.ImageUploader
	ImageBaseURL  string
}
