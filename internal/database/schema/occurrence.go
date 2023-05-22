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
		field.Int("kj").Optional().Nillable(),
		field.Int("kcal").Optional().Nillable(),
		field.Int("fat").Optional().Nillable(),
		field.Int("saturated_fat").Optional().Nillable(),
		field.Int("carbohydrates").Optional().Nillable(),
		field.Int("sugar").Optional().Nillable(),
		field.Int("fiber").Optional().Nillable(),
		field.Int("protein").Optional().Nillable(),
		field.Int("salt").Optional().Nillable(),
		field.Int("price_student").Optional().Nillable(),
		field.Int("price_staff").Optional().Nillable(),
		field.Int("price_guest").Optional().Nillable(),
	}
}

// Edges of the Occurrence.
func (Occurrence) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("location", Location.Type).Ref("occurrences").Unique().Required(),
		edge.From("dish", Dish.Type).Ref("occurrences").Unique().Required(),
		edge.From("tag", Tag.Type).Ref("occurrences"),
		edge.To("side_dishes", Dish.Type), // todo: create reverse edge
		edge.To("reviews", Review.Type).StorageKey(edge.Column("occurrence")),
	}
}
