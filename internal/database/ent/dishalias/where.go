// Code generated by ent, DO NOT EDIT.

package dishalias

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mensatt/backend/internal/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AliasName applies equality check predicate on the "alias_name" field. It's identical to AliasNameEQ.
func AliasName(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAliasName), v))
	})
}

// NormalizedAliasName applies equality check predicate on the "normalized_alias_name" field. It's identical to NormalizedAliasNameEQ.
func NormalizedAliasName(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNormalizedAliasName), v))
	})
}

// AliasNameEQ applies the EQ predicate on the "alias_name" field.
func AliasNameEQ(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAliasName), v))
	})
}

// AliasNameNEQ applies the NEQ predicate on the "alias_name" field.
func AliasNameNEQ(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAliasName), v))
	})
}

// AliasNameIn applies the In predicate on the "alias_name" field.
func AliasNameIn(vs ...string) predicate.DishAlias {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAliasName), v...))
	})
}

// AliasNameNotIn applies the NotIn predicate on the "alias_name" field.
func AliasNameNotIn(vs ...string) predicate.DishAlias {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAliasName), v...))
	})
}

// AliasNameGT applies the GT predicate on the "alias_name" field.
func AliasNameGT(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAliasName), v))
	})
}

// AliasNameGTE applies the GTE predicate on the "alias_name" field.
func AliasNameGTE(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAliasName), v))
	})
}

// AliasNameLT applies the LT predicate on the "alias_name" field.
func AliasNameLT(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAliasName), v))
	})
}

// AliasNameLTE applies the LTE predicate on the "alias_name" field.
func AliasNameLTE(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAliasName), v))
	})
}

// AliasNameContains applies the Contains predicate on the "alias_name" field.
func AliasNameContains(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAliasName), v))
	})
}

// AliasNameHasPrefix applies the HasPrefix predicate on the "alias_name" field.
func AliasNameHasPrefix(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAliasName), v))
	})
}

// AliasNameHasSuffix applies the HasSuffix predicate on the "alias_name" field.
func AliasNameHasSuffix(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAliasName), v))
	})
}

// AliasNameEqualFold applies the EqualFold predicate on the "alias_name" field.
func AliasNameEqualFold(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAliasName), v))
	})
}

// AliasNameContainsFold applies the ContainsFold predicate on the "alias_name" field.
func AliasNameContainsFold(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAliasName), v))
	})
}

// NormalizedAliasNameEQ applies the EQ predicate on the "normalized_alias_name" field.
func NormalizedAliasNameEQ(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNormalizedAliasName), v))
	})
}

// NormalizedAliasNameNEQ applies the NEQ predicate on the "normalized_alias_name" field.
func NormalizedAliasNameNEQ(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNormalizedAliasName), v))
	})
}

// NormalizedAliasNameIn applies the In predicate on the "normalized_alias_name" field.
func NormalizedAliasNameIn(vs ...string) predicate.DishAlias {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNormalizedAliasName), v...))
	})
}

// NormalizedAliasNameNotIn applies the NotIn predicate on the "normalized_alias_name" field.
func NormalizedAliasNameNotIn(vs ...string) predicate.DishAlias {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNormalizedAliasName), v...))
	})
}

// NormalizedAliasNameGT applies the GT predicate on the "normalized_alias_name" field.
func NormalizedAliasNameGT(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNormalizedAliasName), v))
	})
}

// NormalizedAliasNameGTE applies the GTE predicate on the "normalized_alias_name" field.
func NormalizedAliasNameGTE(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNormalizedAliasName), v))
	})
}

// NormalizedAliasNameLT applies the LT predicate on the "normalized_alias_name" field.
func NormalizedAliasNameLT(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNormalizedAliasName), v))
	})
}

// NormalizedAliasNameLTE applies the LTE predicate on the "normalized_alias_name" field.
func NormalizedAliasNameLTE(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNormalizedAliasName), v))
	})
}

// NormalizedAliasNameContains applies the Contains predicate on the "normalized_alias_name" field.
func NormalizedAliasNameContains(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNormalizedAliasName), v))
	})
}

// NormalizedAliasNameHasPrefix applies the HasPrefix predicate on the "normalized_alias_name" field.
func NormalizedAliasNameHasPrefix(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNormalizedAliasName), v))
	})
}

// NormalizedAliasNameHasSuffix applies the HasSuffix predicate on the "normalized_alias_name" field.
func NormalizedAliasNameHasSuffix(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNormalizedAliasName), v))
	})
}

// NormalizedAliasNameEqualFold applies the EqualFold predicate on the "normalized_alias_name" field.
func NormalizedAliasNameEqualFold(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNormalizedAliasName), v))
	})
}

// NormalizedAliasNameContainsFold applies the ContainsFold predicate on the "normalized_alias_name" field.
func NormalizedAliasNameContainsFold(v string) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNormalizedAliasName), v))
	})
}

// HasDish applies the HasEdge predicate on the "dish" edge.
func HasDish() predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DishTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DishTable, DishColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDishWith applies the HasEdge predicate on the "dish" edge with a given conditions (other predicates).
func HasDishWith(preds ...predicate.Dish) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DishInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DishTable, DishColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DishAlias) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DishAlias) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
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
func Not(p predicate.DishAlias) predicate.DishAlias {
	return predicate.DishAlias(func(s *sql.Selector) {
		p(s.Not())
	})
}
