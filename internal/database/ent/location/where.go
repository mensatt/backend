// Code generated by ent, DO NOT EDIT.

package location

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ExternalID applies equality check predicate on the "external_id" field. It's identical to ExternalIDEQ.
func ExternalID(v int) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExternalID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Visible applies equality check predicate on the "visible" field. It's identical to VisibleEQ.
func Visible(v bool) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVisible), v))
	})
}

// ExternalIDEQ applies the EQ predicate on the "external_id" field.
func ExternalIDEQ(v int) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExternalID), v))
	})
}

// ExternalIDNEQ applies the NEQ predicate on the "external_id" field.
func ExternalIDNEQ(v int) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExternalID), v))
	})
}

// ExternalIDIn applies the In predicate on the "external_id" field.
func ExternalIDIn(vs ...int) predicate.Location {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExternalID), v...))
	})
}

// ExternalIDNotIn applies the NotIn predicate on the "external_id" field.
func ExternalIDNotIn(vs ...int) predicate.Location {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExternalID), v...))
	})
}

// ExternalIDGT applies the GT predicate on the "external_id" field.
func ExternalIDGT(v int) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExternalID), v))
	})
}

// ExternalIDGTE applies the GTE predicate on the "external_id" field.
func ExternalIDGTE(v int) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExternalID), v))
	})
}

// ExternalIDLT applies the LT predicate on the "external_id" field.
func ExternalIDLT(v int) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExternalID), v))
	})
}

// ExternalIDLTE applies the LTE predicate on the "external_id" field.
func ExternalIDLTE(v int) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExternalID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Location {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Location {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// VisibleEQ applies the EQ predicate on the "visible" field.
func VisibleEQ(v bool) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVisible), v))
	})
}

// VisibleNEQ applies the NEQ predicate on the "visible" field.
func VisibleNEQ(v bool) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVisible), v))
	})
}

// HasOccurrences applies the HasEdge predicate on the "occurrences" edge.
func HasOccurrences() predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OccurrencesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OccurrencesTable, OccurrencesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOccurrencesWith applies the HasEdge predicate on the "occurrences" edge with a given conditions (other predicates).
func HasOccurrencesWith(preds ...predicate.Occurrence) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OccurrencesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OccurrencesTable, OccurrencesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Location) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Location) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
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
func Not(p predicate.Location) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		p(s.Not())
	})
}
