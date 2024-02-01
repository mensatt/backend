package resolvers

import (
	ent "github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/pkg/utils"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Database               *ent.Client
	JWTKeyStore            *utils.JWTKeyStore
	VCSBuildInfo           *utils.VCSBuildInfo
	ImageBaseURL           string
	ReviewCreatedChannels  map[string]chan *ent.Review
	ReviewAcceptedChannels map[string]chan *ent.Review
	mutex                  sync.Mutex
}
