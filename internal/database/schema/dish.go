package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Dish holds the schema definition for the Dish entity.
type Dish struct {
	ent.Schema
}

// Annotations of the Dish.
func (Dish) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "dish"},
	}
}

// Fields of the Dish.
func (Dish) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("name_de").Unique().NotEmpty(),
		field.String("name_en").Optional().Nillable().NotEmpty(),
	}
}

// Edges of the Dish.
func (Dish) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dish_occurrences", Occurrence.Type).StorageKey(edge.Column("dish")),
		edge.To("aliases", DishAlias.Type).StorageKey(edge.Column("dish")),
		edge.From("side_dish_occurrence", Occurrence.Type).Ref("side_dishes"),
	}
}
