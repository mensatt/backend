package directives

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/mensatt/backend/internal/database/schema"
	"github.com/mensatt/backend/internal/middleware"
)

func Authenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	user := middleware.GetUserIDFromCtx(ctx)
	if user == nil {
		return nil, fmt.Errorf("not authenticated")
	}
	return next(ctx)
}

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver, requiredRole *schema.UserRole) (interface{}, error) {
	user := middleware.GetUserIDFromCtx(ctx)
	if user == nil {
		return nil, fmt.Errorf("not authenticated")
	}

	if requiredRole != nil && (user.Role == nil || *user.Role != *requiredRole) {
		return nil, fmt.Errorf("not authorized")
	}

	return next(ctx)
}
