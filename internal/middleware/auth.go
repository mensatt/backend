package middleware

import (
	"context"
	"github.com/getsentry/sentry-go"
	ent "github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/internal/database/ent/user"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mensatt/backend/pkg/utils"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type AuthParams struct {
	JWTKeyStore *utils.JWTKeyStore
	Database    *ent.Client
}

// Auth wraps the request with auth middleware
func Auth(params AuthParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return // No token, no auth
		}

		userID, err := params.JWTKeyStore.ParseJWT(tokenString)
		if err != nil {
			log.Println("[Auth Middleware]: Failed to parse or invalid JWT:", err)
			return
		}

		currentUser, err := params.Database.User.Query().Where(user.ID(userID)).Only(c)
		//print(currentUser.ID)

		//currentUser, err := params.Database.GetUserByID(c, userID)
		if err != nil {
			log.Printf("[Auth Middleware]: Failed to find user with id: %s - %v\n", userID, err)
			return
		}

		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetUser(sentry.User{
				ID:    currentUser.ID.String(),
				Email: currentUser.Email,
			})
		})

		ctx := context.WithValue(c.Request.Context(), userCtxKey, currentUser)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetUserIDFromCtx(ctx context.Context) *ent.User {
	user := ctx.Value(userCtxKey)
	if user == nil {
		return nil
	}

	return user.(*ent.User)
}
