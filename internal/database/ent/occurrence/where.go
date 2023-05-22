// Code generated by ent, DO NOT EDIT.

package occurrence

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent/predicate"
	"github.com/mensatt/backend/internal/database/schema"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// Kj applies equality check predicate on the "kj" field. It's identical to KjEQ.
func Kj(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKj), v))
	})
}

// Kcal applies equality check predicate on the "kcal" field. It's identical to KcalEQ.
func Kcal(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKcal), v))
	})
}

// Fat applies equality check predicate on the "fat" field. It's identical to FatEQ.
func Fat(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFat), v))
	})
}

// SaturatedFat applies equality check predicate on the "saturated_fat" field. It's identical to SaturatedFatEQ.
func SaturatedFat(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSaturatedFat), v))
	})
}

// Carbohydrates applies equality check predicate on the "carbohydrates" field. It's identical to CarbohydratesEQ.
func Carbohydrates(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCarbohydrates), v))
	})
}

// Sugar applies equality check predicate on the "sugar" field. It's identical to SugarEQ.
func Sugar(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSugar), v))
	})
}

// Fiber applies equality check predicate on the "fiber" field. It's identical to FiberEQ.
func Fiber(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFiber), v))
	})
}

// Protein applies equality check predicate on the "protein" field. It's identical to ProteinEQ.
func Protein(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProtein), v))
	})
}

// Salt applies equality check predicate on the "salt" field. It's identical to SaltEQ.
func Salt(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalt), v))
	})
}

// PriceStudent applies equality check predicate on the "price_student" field. It's identical to PriceStudentEQ.
func PriceStudent(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriceStudent), v))
	})
}

// PriceStaff applies equality check predicate on the "price_staff" field. It's identical to PriceStaffEQ.
func PriceStaff(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriceStaff), v))
	})
}

// PriceGuest applies equality check predicate on the "price_guest" field. It's identical to PriceGuestEQ.
func PriceGuest(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriceGuest), v))
	})
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v schema.OccurrenceStatus) predicate.Occurrence {
	vc := v
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), vc))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v schema.OccurrenceStatus) predicate.Occurrence {
	vc := v
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), vc))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...schema.OccurrenceStatus) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...schema.OccurrenceStatus) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// KjEQ applies the EQ predicate on the "kj" field.
func KjEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKj), v))
	})
}

// KjNEQ applies the NEQ predicate on the "kj" field.
func KjNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKj), v))
	})
}

// KjIn applies the In predicate on the "kj" field.
func KjIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldKj), v...))
	})
}

// KjNotIn applies the NotIn predicate on the "kj" field.
func KjNotIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldKj), v...))
	})
}

// KjGT applies the GT predicate on the "kj" field.
func KjGT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKj), v))
	})
}

// KjGTE applies the GTE predicate on the "kj" field.
func KjGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKj), v))
	})
}

// KjLT applies the LT predicate on the "kj" field.
func KjLT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKj), v))
	})
}

// KjLTE applies the LTE predicate on the "kj" field.
func KjLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKj), v))
	})
}

// KcalEQ applies the EQ predicate on the "kcal" field.
func KcalEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKcal), v))
	})
}

// KcalNEQ applies the NEQ predicate on the "kcal" field.
func KcalNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKcal), v))
	})
}

// KcalIn applies the In predicate on the "kcal" field.
func KcalIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldKcal), v...))
	})
}

// KcalNotIn applies the NotIn predicate on the "kcal" field.
func KcalNotIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldKcal), v...))
	})
}

// KcalGT applies the GT predicate on the "kcal" field.
func KcalGT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKcal), v))
	})
}

// KcalGTE applies the GTE predicate on the "kcal" field.
func KcalGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKcal), v))
	})
}

// KcalLT applies the LT predicate on the "kcal" field.
func KcalLT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKcal), v))
	})
}

// KcalLTE applies the LTE predicate on the "kcal" field.
func KcalLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKcal), v))
	})
}

// FatEQ applies the EQ predicate on the "fat" field.
func FatEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFat), v))
	})
}

// FatNEQ applies the NEQ predicate on the "fat" field.
func FatNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFat), v))
	})
}

// FatIn applies the In predicate on the "fat" field.
func FatIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFat), v...))
	})
}

// FatNotIn applies the NotIn predicate on the "fat" field.
func FatNotIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFat), v...))
	})
}

// FatGT applies the GT predicate on the "fat" field.
func FatGT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFat), v))
	})
}

// FatGTE applies the GTE predicate on the "fat" field.
func FatGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFat), v))
	})
}

// FatLT applies the LT predicate on the "fat" field.
func FatLT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFat), v))
	})
}

// FatLTE applies the LTE predicate on the "fat" field.
func FatLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFat), v))
	})
}

// SaturatedFatEQ applies the EQ predicate on the "saturated_fat" field.
func SaturatedFatEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSaturatedFat), v))
	})
}

// SaturatedFatNEQ applies the NEQ predicate on the "saturated_fat" field.
func SaturatedFatNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSaturatedFat), v))
	})
}

// SaturatedFatIn applies the In predicate on the "saturated_fat" field.
func SaturatedFatIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSaturatedFat), v...))
	})
}

// SaturatedFatNotIn applies the NotIn predicate on the "saturated_fat" field.
func SaturatedFatNotIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSaturatedFat), v...))
	})
}

// SaturatedFatGT applies the GT predicate on the "saturated_fat" field.
func SaturatedFatGT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSaturatedFat), v))
	})
}

// SaturatedFatGTE applies the GTE predicate on the "saturated_fat" field.
func SaturatedFatGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSaturatedFat), v))
	})
}

// SaturatedFatLT applies the LT predicate on the "saturated_fat" field.
func SaturatedFatLT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSaturatedFat), v))
	})
}

// SaturatedFatLTE applies the LTE predicate on the "saturated_fat" field.
func SaturatedFatLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSaturatedFat), v))
	})
}

// CarbohydratesEQ applies the EQ predicate on the "carbohydrates" field.
func CarbohydratesEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCarbohydrates), v))
	})
}

// CarbohydratesNEQ applies the NEQ predicate on the "carbohydrates" field.
func CarbohydratesNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCarbohydrates), v))
	})
}

// CarbohydratesIn applies the In predicate on the "carbohydrates" field.
func CarbohydratesIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCarbohydrates), v...))
	})
}

// CarbohydratesNotIn applies the NotIn predicate on the "carbohydrates" field.
func CarbohydratesNotIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCarbohydrates), v...))
	})
}

// CarbohydratesGT applies the GT predicate on the "carbohydrates" field.
func CarbohydratesGT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCarbohydrates), v))
	})
}

// CarbohydratesGTE applies the GTE predicate on the "carbohydrates" field.
func CarbohydratesGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCarbohydrates), v))
	})
}

// CarbohydratesLT applies the LT predicate on the "carbohydrates" field.
func CarbohydratesLT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCarbohydrates), v))
	})
}

// CarbohydratesLTE applies the LTE predicate on the "carbohydrates" field.
func CarbohydratesLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCarbohydrates), v))
	})
}

// SugarEQ applies the EQ predicate on the "sugar" field.
func SugarEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSugar), v))
	})
}

// SugarNEQ applies the NEQ predicate on the "sugar" field.
func SugarNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSugar), v))
	})
}

// SugarIn applies the In predicate on the "sugar" field.
func SugarIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSugar), v...))
	})
}

// SugarNotIn applies the NotIn predicate on the "sugar" field.
func SugarNotIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSugar), v...))
	})
}

// SugarGT applies the GT predicate on the "sugar" field.
func SugarGT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSugar), v))
	})
}

// SugarGTE applies the GTE predicate on the "sugar" field.
func SugarGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSugar), v))
	})
}

// SugarLT applies the LT predicate on the "sugar" field.
func SugarLT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSugar), v))
	})
}

// SugarLTE applies the LTE predicate on the "sugar" field.
func SugarLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSugar), v))
	})
}

// FiberEQ applies the EQ predicate on the "fiber" field.
func FiberEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFiber), v))
	})
}

// FiberNEQ applies the NEQ predicate on the "fiber" field.
func FiberNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFiber), v))
	})
}

// FiberIn applies the In predicate on the "fiber" field.
func FiberIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFiber), v...))
	})
}

// FiberNotIn applies the NotIn predicate on the "fiber" field.
func FiberNotIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFiber), v...))
	})
}

// FiberGT applies the GT predicate on the "fiber" field.
func FiberGT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFiber), v))
	})
}

// FiberGTE applies the GTE predicate on the "fiber" field.
func FiberGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFiber), v))
	})
}

// FiberLT applies the LT predicate on the "fiber" field.
func FiberLT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFiber), v))
	})
}

// FiberLTE applies the LTE predicate on the "fiber" field.
func FiberLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFiber), v))
	})
}

// ProteinEQ applies the EQ predicate on the "protein" field.
func ProteinEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProtein), v))
	})
}

// ProteinNEQ applies the NEQ predicate on the "protein" field.
func ProteinNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProtein), v))
	})
}

// ProteinIn applies the In predicate on the "protein" field.
func ProteinIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldProtein), v...))
	})
}

// ProteinNotIn applies the NotIn predicate on the "protein" field.
func ProteinNotIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldProtein), v...))
	})
}

// ProteinGT applies the GT predicate on the "protein" field.
func ProteinGT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProtein), v))
	})
}

// ProteinGTE applies the GTE predicate on the "protein" field.
func ProteinGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProtein), v))
	})
}

// ProteinLT applies the LT predicate on the "protein" field.
func ProteinLT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProtein), v))
	})
}

// ProteinLTE applies the LTE predicate on the "protein" field.
func ProteinLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProtein), v))
	})
}

// SaltEQ applies the EQ predicate on the "salt" field.
func SaltEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalt), v))
	})
}

// SaltNEQ applies the NEQ predicate on the "salt" field.
func SaltNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSalt), v))
	})
}

// SaltIn applies the In predicate on the "salt" field.
func SaltIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSalt), v...))
	})
}

// SaltNotIn applies the NotIn predicate on the "salt" field.
func SaltNotIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSalt), v...))
	})
}

// SaltGT applies the GT predicate on the "salt" field.
func SaltGT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSalt), v))
	})
}

// SaltGTE applies the GTE predicate on the "salt" field.
func SaltGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSalt), v))
	})
}

// SaltLT applies the LT predicate on the "salt" field.
func SaltLT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSalt), v))
	})
}

// SaltLTE applies the LTE predicate on the "salt" field.
func SaltLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSalt), v))
	})
}

// PriceStudentEQ applies the EQ predicate on the "price_student" field.
func PriceStudentEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriceStudent), v))
	})
}

// PriceStudentNEQ applies the NEQ predicate on the "price_student" field.
func PriceStudentNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPriceStudent), v))
	})
}

// PriceStudentIn applies the In predicate on the "price_student" field.
func PriceStudentIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPriceStudent), v...))
	})
}

// PriceStudentNotIn applies the NotIn predicate on the "price_student" field.
func PriceStudentNotIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPriceStudent), v...))
	})
}

// PriceStudentGT applies the GT predicate on the "price_student" field.
func PriceStudentGT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPriceStudent), v))
	})
}

// PriceStudentGTE applies the GTE predicate on the "price_student" field.
func PriceStudentGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPriceStudent), v))
	})
}

// PriceStudentLT applies the LT predicate on the "price_student" field.
func PriceStudentLT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPriceStudent), v))
	})
}

// PriceStudentLTE applies the LTE predicate on the "price_student" field.
func PriceStudentLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPriceStudent), v))
	})
}

// PriceStaffEQ applies the EQ predicate on the "price_staff" field.
func PriceStaffEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriceStaff), v))
	})
}

// PriceStaffNEQ applies the NEQ predicate on the "price_staff" field.
func PriceStaffNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPriceStaff), v))
	})
}

// PriceStaffIn applies the In predicate on the "price_staff" field.
func PriceStaffIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPriceStaff), v...))
	})
}

// PriceStaffNotIn applies the NotIn predicate on the "price_staff" field.
func PriceStaffNotIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPriceStaff), v...))
	})
}

// PriceStaffGT applies the GT predicate on the "price_staff" field.
func PriceStaffGT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPriceStaff), v))
	})
}

// PriceStaffGTE applies the GTE predicate on the "price_staff" field.
func PriceStaffGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPriceStaff), v))
	})
}

// PriceStaffLT applies the LT predicate on the "price_staff" field.
func PriceStaffLT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPriceStaff), v))
	})
}

// PriceStaffLTE applies the LTE predicate on the "price_staff" field.
func PriceStaffLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPriceStaff), v))
	})
}

// PriceGuestEQ applies the EQ predicate on the "price_guest" field.
func PriceGuestEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriceGuest), v))
	})
}

// PriceGuestNEQ applies the NEQ predicate on the "price_guest" field.
func PriceGuestNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPriceGuest), v))
	})
}

// PriceGuestIn applies the In predicate on the "price_guest" field.
func PriceGuestIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPriceGuest), v...))
	})
}

// PriceGuestNotIn applies the NotIn predicate on the "price_guest" field.
func PriceGuestNotIn(vs ...int) predicate.Occurrence {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPriceGuest), v...))
	})
}

// PriceGuestGT applies the GT predicate on the "price_guest" field.
func PriceGuestGT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPriceGuest), v))
	})
}

// PriceGuestGTE applies the GTE predicate on the "price_guest" field.
func PriceGuestGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPriceGuest), v))
	})
}

// PriceGuestLT applies the LT predicate on the "price_guest" field.
func PriceGuestLT(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPriceGuest), v))
	})
}

// PriceGuestLTE applies the LTE predicate on the "price_guest" field.
func PriceGuestLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPriceGuest), v))
	})
}

// HasLocation applies the HasEdge predicate on the "location" edge.
func HasLocation() predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LocationTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LocationTable, LocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationWith applies the HasEdge predicate on the "location" edge with a given conditions (other predicates).
func HasLocationWith(preds ...predicate.Location) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LocationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LocationTable, LocationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDish applies the HasEdge predicate on the "dish" edge.
func HasDish() predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DishTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DishTable, DishColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDishWith applies the HasEdge predicate on the "dish" edge with a given conditions (other predicates).
func HasDishWith(preds ...predicate.Dish) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DishInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DishTable, DishColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSideDishes applies the HasEdge predicate on the "side_dishes" edge.
func HasSideDishes() predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SideDishesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SideDishesTable, SideDishesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSideDishesWith applies the HasEdge predicate on the "side_dishes" edge with a given conditions (other predicates).
func HasSideDishesWith(preds ...predicate.Dish) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SideDishesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SideDishesTable, SideDishesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagTable, TagPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagTable, TagPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReviews applies the HasEdge predicate on the "reviews" edge.
func HasReviews() predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReviewsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReviewsTable, ReviewsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReviewsWith applies the HasEdge predicate on the "reviews" edge with a given conditions (other predicates).
func HasReviewsWith(preds ...predicate.Review) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReviewsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReviewsTable, ReviewsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Occurrence) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Occurrence) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
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
func Not(p predicate.Occurrence) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		p(s.Not())
	})
}
