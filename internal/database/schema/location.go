package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.Int("external_id").Unique(),
		field.String("name").Unique().NotEmpty(),
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return nil
}
