package middleware

import (
	"context"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/db/sqlc"
	"github.com/mensatt/mensatt-backend/pkg/utils"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type AuthParams struct {
	KeyStore   *utils.JWTKeyStore
	Database   db.ExtendedQuerier
	CookieName string
}

// Auth wraps the request with auth middleware
func Auth(params AuthParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return // No token, no auth
		}

		userID, err := params.KeyStore.ParseJWT(tokenString)
		if err != nil {
			log.Println("[Auth Middleware]: Failed to parse or invalid JWT:", err)
			return
		}

		currentUser, err := params.Database.GetUserByID(c, userID)
		if err != nil {
			log.Printf("[Auth Middleware]: Failed to find user with id: %s - %v\n", userID, err)
			return
		}

		ctx := context.WithValue(c.Request.Context(), userCtxKey, currentUser)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetUserIDFromCtx(ctx context.Context) *sqlc.User {
	user := ctx.Value(userCtxKey)
	if user == nil {
		return nil
	}

	return user.(*sqlc.User)
}