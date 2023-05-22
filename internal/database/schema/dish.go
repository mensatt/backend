package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Dish holds the schema definition for the Dish entity.
type Dish struct {
	ent.Schema
}

// Fields of the Dish.
func (Dish) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("name_de").Unique(),
		field.String("name_en"),
	}
}

// Edges of the Dish.
func (Dish) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("occurrences", Occurrence.Type).StorageKey(edge.Column("dish")),
		edge.To("aliases", DishAlias.Type).StorageKey(edge.Column("dish")),
	}
}
