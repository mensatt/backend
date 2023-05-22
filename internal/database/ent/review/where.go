// Code generated by ent, DO NOT EDIT.

package review

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplayName), v))
	})
}

// Stars applies equality check predicate on the "stars" field. It's identical to StarsEQ.
func Stars(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStars), v))
	})
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// AcceptedAt applies equality check predicate on the "accepted_at" field. It's identical to AcceptedAtEQ.
func AcceptedAt(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAcceptedAt), v))
	})
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplayName), v))
	})
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDisplayName), v))
	})
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Review {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDisplayName), v...))
	})
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Review {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDisplayName), v...))
	})
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDisplayName), v))
	})
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDisplayName), v))
	})
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDisplayName), v))
	})
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDisplayName), v))
	})
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDisplayName), v))
	})
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDisplayName), v))
	})
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDisplayName), v))
	})
}

// DisplayNameIsNil applies the IsNil predicate on the "display_name" field.
func DisplayNameIsNil() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDisplayName)))
	})
}

// DisplayNameNotNil applies the NotNil predicate on the "display_name" field.
func DisplayNameNotNil() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDisplayName)))
	})
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDisplayName), v))
	})
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDisplayName), v))
	})
}

// StarsEQ applies the EQ predicate on the "stars" field.
func StarsEQ(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStars), v))
	})
}

// StarsNEQ applies the NEQ predicate on the "stars" field.
func StarsNEQ(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStars), v))
	})
}

// StarsIn applies the In predicate on the "stars" field.
func StarsIn(vs ...int) predicate.Review {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStars), v...))
	})
}

// StarsNotIn applies the NotIn predicate on the "stars" field.
func StarsNotIn(vs ...int) predicate.Review {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStars), v...))
	})
}

// StarsGT applies the GT predicate on the "stars" field.
func StarsGT(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStars), v))
	})
}

// StarsGTE applies the GTE predicate on the "stars" field.
func StarsGTE(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStars), v))
	})
}

// StarsLT applies the LT predicate on the "stars" field.
func StarsLT(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStars), v))
	})
}

// StarsLTE applies the LTE predicate on the "stars" field.
func StarsLTE(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStars), v))
	})
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldText), v))
	})
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Review {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldText), v...))
	})
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Review {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldText), v...))
	})
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldText), v))
	})
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldText), v))
	})
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldText), v))
	})
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldText), v))
	})
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldText), v))
	})
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldText), v))
	})
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldText), v))
	})
}

// TextIsNil applies the IsNil predicate on the "text" field.
func TextIsNil() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldText)))
	})
}

// TextNotNil applies the NotNil predicate on the "text" field.
func TextNotNil() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldText)))
	})
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldText), v))
	})
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldText), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Review {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Review {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Review {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Review {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// AcceptedAtEQ applies the EQ predicate on the "accepted_at" field.
func AcceptedAtEQ(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAcceptedAt), v))
	})
}

// AcceptedAtNEQ applies the NEQ predicate on the "accepted_at" field.
func AcceptedAtNEQ(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAcceptedAt), v))
	})
}

// AcceptedAtIn applies the In predicate on the "accepted_at" field.
func AcceptedAtIn(vs ...time.Time) predicate.Review {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAcceptedAt), v...))
	})
}

// AcceptedAtNotIn applies the NotIn predicate on the "accepted_at" field.
func AcceptedAtNotIn(vs ...time.Time) predicate.Review {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAcceptedAt), v...))
	})
}

// AcceptedAtGT applies the GT predicate on the "accepted_at" field.
func AcceptedAtGT(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAcceptedAt), v))
	})
}

// AcceptedAtGTE applies the GTE predicate on the "accepted_at" field.
func AcceptedAtGTE(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAcceptedAt), v))
	})
}

// AcceptedAtLT applies the LT predicate on the "accepted_at" field.
func AcceptedAtLT(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAcceptedAt), v))
	})
}

// AcceptedAtLTE applies the LTE predicate on the "accepted_at" field.
func AcceptedAtLTE(v time.Time) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAcceptedAt), v))
	})
}

// AcceptedAtIsNil applies the IsNil predicate on the "accepted_at" field.
func AcceptedAtIsNil() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAcceptedAt)))
	})
}

// AcceptedAtNotNil applies the NotNil predicate on the "accepted_at" field.
func AcceptedAtNotNil() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAcceptedAt)))
	})
}

// HasOccurrence applies the HasEdge predicate on the "occurrence" edge.
func HasOccurrence() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OccurrenceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OccurrenceTable, OccurrenceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOccurrenceWith applies the HasEdge predicate on the "occurrence" edge with a given conditions (other predicates).
func HasOccurrenceWith(preds ...predicate.Occurrence) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OccurrenceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OccurrenceTable, OccurrenceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasImages applies the HasEdge predicate on the "images" edge.
func HasImages() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ImagesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImagesWith applies the HasEdge predicate on the "images" edge with a given conditions (other predicates).
func HasImagesWith(preds ...predicate.Image) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ImagesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		p(s.Not())
	})
}
