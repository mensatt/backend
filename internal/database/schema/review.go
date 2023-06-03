package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Annotations of the Review.
func (Review) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "review"},
	}
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("display_name").Optional().Nillable().NotEmpty().MaxLen(32),
		field.Int("stars").Min(1).Max(5),
		field.String("text").Optional().Nillable().NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("accepted_at").Optional().Nillable(),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("occurrence", Occurrence.Type).Ref("reviews").Unique().Required(),
		edge.To("images", Image.Type).StorageKey(edge.Column("review")),
	}
}
