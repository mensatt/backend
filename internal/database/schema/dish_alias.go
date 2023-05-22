package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DishAlias holds the schema definition for the DishAlias entity.
type DishAlias struct {
	ent.Schema
}

// Fields of the DishAlias.
func (DishAlias) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("alias_name"),
		field.String("normalized_alias_name").NotEmpty(),
	}
}

// Edges of the DishAlias.
func (DishAlias) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dish", Dish.Type).Ref("aliases").Unique().Required(),
	}
}
