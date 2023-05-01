package seeds

import (
	"context"
	"github.com/mensatt/backend/internal/database/ent"
)

// todo: use atlas to seed data in the future
func Seed(ctx context.Context, client *ent.Client) error {
	err := seedTags(ctx, client)
	if err != nil {
		return err
	}
	err = seedLocations(ctx, client)
	if err != nil {
		return err
	}
	err = seedUsers(ctx, client)
	if err != nil {
		return err
	}
	return nil
}
