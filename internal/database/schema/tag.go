package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Unique(),
		field.String("name").NotEmpty(),
		field.String("description").NotEmpty(),
		field.String("short_name").NotEmpty().Optional(),
		field.Enum("priority").GoType(TagPriority("")).Default("HIDE"),
		field.Bool("is_allergy").Default(false),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("occurrences", Occurrence.Type),
	}
}
