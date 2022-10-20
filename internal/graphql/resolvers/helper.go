package resolvers

import (
	"context"
	"database/sql"
	"io"
	"time"

	"github.com/jackc/pgtype"
	"github.com/mensatt/backend/internal/graphql/models"
)

func (r *Resolver) storeImage(ctx context.Context, image *models.ImageInput) (string, error) {
	imageBytes, err := io.ReadAll(image.Image.File)
	if err != nil {
		return "", err
	}

	return r.ImageUploader.ValidateAndStoreImage(imageBytes)
}

func approvedBoolToNullTime(approved bool) sql.NullTime {
	if approved {
		return sql.NullTime{Valid: true, Time: time.Now()}
	}
	return sql.NullTime{Valid: false}
}

func pgtypeNumericToFloat(numeric pgtype.Numeric) (*float64, error) {
	if numeric.Status == pgtype.Null {
		return nil, nil
	}

	var f float64
	err := numeric.AssignTo(&f)
	return &f, err
}

func boolPointerToNullBool(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{Valid: false}
	}
	return sql.NullBool{Valid: true, Bool: *b}
}
