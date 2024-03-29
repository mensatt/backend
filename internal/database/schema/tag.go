package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Annotations of the Tag.
func (Tag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tag"},
	}
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("key"),
		field.String("name").NotEmpty(),
		field.String("description").NotEmpty(),
		field.String("short_name").Optional().Nillable().NotEmpty(),
		field.Enum("priority").GoType(TagPriority("")).Default("HIDE"),
		field.Bool("is_allergy").Default(false),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("occurrence", Occurrence.Type).Ref("tags"),
	}
}
