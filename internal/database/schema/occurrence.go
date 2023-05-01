package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Occurrence holds the schema definition for the Occurrence entity.
type Occurrence struct {
	ent.Schema
}

// Fields of the Occurrence.
func (Occurrence) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.Time("date"),
		field.Enum("status").GoType(OccurrenceStatus("")).Default("AWAITING_APPROVAL"),
		field.Int("kj"),
		field.Int("kcal"),
		field.Int("fat"),
		field.Int("saturated_fat"),
		field.Int("carbohydrates"),
		field.Int("sugar"),
		field.Int("fiber"),
		field.Int("protein"),
		field.Int("salt"),
		field.Int("price_student"),
		field.Int("price_staff"),
		field.Int("price_guest"),
	}
}

// Edges of the Occurrence.
func (Occurrence) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("location", Location.Type).Unique(),
		edge.To("dish", Dish.Type).Unique(),
		edge.To("side_dishes", Dish.Type),
		edge.To("tags", Tag.Type),
	}
}
