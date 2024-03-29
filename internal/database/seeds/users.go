package seeds

import (
	"context"
	"github.com/mensatt/backend/internal/database/ent"
)

func seedUsers(ctx context.Context, client *ent.Client) error {
	err := client.User.CreateBulk(
		client.User.Create().SetEmail("admin@mensatt.de").SetPasswordHash("$2a$10$pdvY6v8k2McSYbFk3HRDl.h8QfMjOxfpm2CywkDDzfOzlYDZV8NUm"),
	).OnConflict().DoNothing().Exec(ctx)
	return err
}
