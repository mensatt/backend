package dataloader

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mensatt/backend/internal/db"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var dataloaderCtxKey = &contextKey{"dataloader"}

type contextKey struct {
	name string
}

// Middleware injects a DataLoader into the request context so it can be
// used later in the schema resolvers
func Middleware(db db.ExtendedQuerier) gin.HandlerFunc {
	dl := NewDataLoader(db)

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), dataloaderCtxKey, dl)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// For returns the dataloader for a given context
func FromContext(ctx context.Context) *DataLoader {
	return ctx.Value(dataloaderCtxKey).(*DataLoader)
}
