// Code generated by ent, DO NOT EDIT.

package occurrence

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldID, id))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldDate, v))
}

// Kj applies equality check predicate on the "kj" field. It's identical to KjEQ.
func Kj(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldKj, v))
}

// Kcal applies equality check predicate on the "kcal" field. It's identical to KcalEQ.
func Kcal(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldKcal, v))
}

// Fat applies equality check predicate on the "fat" field. It's identical to FatEQ.
func Fat(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldFat, v))
}

// SaturatedFat applies equality check predicate on the "saturated_fat" field. It's identical to SaturatedFatEQ.
func SaturatedFat(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldSaturatedFat, v))
}

// Carbohydrates applies equality check predicate on the "carbohydrates" field. It's identical to CarbohydratesEQ.
func Carbohydrates(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldCarbohydrates, v))
}

// Sugar applies equality check predicate on the "sugar" field. It's identical to SugarEQ.
func Sugar(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldSugar, v))
}

// Fiber applies equality check predicate on the "fiber" field. It's identical to FiberEQ.
func Fiber(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldFiber, v))
}

// Protein applies equality check predicate on the "protein" field. It's identical to ProteinEQ.
func Protein(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldProtein, v))
}

// Salt applies equality check predicate on the "salt" field. It's identical to SaltEQ.
func Salt(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldSalt, v))
}

// PriceStudent applies equality check predicate on the "price_student" field. It's identical to PriceStudentEQ.
func PriceStudent(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldPriceStudent, v))
}

// PriceStaff applies equality check predicate on the "price_staff" field. It's identical to PriceStaffEQ.
func PriceStaff(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldPriceStaff, v))
}

// PriceGuest applies equality check predicate on the "price_guest" field. It's identical to PriceGuestEQ.
func PriceGuest(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldPriceGuest, v))
}

// NotAvailableAfter applies equality check predicate on the "not_available_after" field. It's identical to NotAvailableAfterEQ.
func NotAvailableAfter(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldNotAvailableAfter, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldDate, v))
}

// KjEQ applies the EQ predicate on the "kj" field.
func KjEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldKj, v))
}

// KjNEQ applies the NEQ predicate on the "kj" field.
func KjNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldKj, v))
}

// KjIn applies the In predicate on the "kj" field.
func KjIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldKj, vs...))
}

// KjNotIn applies the NotIn predicate on the "kj" field.
func KjNotIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldKj, vs...))
}

// KjGT applies the GT predicate on the "kj" field.
func KjGT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldKj, v))
}

// KjGTE applies the GTE predicate on the "kj" field.
func KjGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldKj, v))
}

// KjLT applies the LT predicate on the "kj" field.
func KjLT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldKj, v))
}

// KjLTE applies the LTE predicate on the "kj" field.
func KjLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldKj, v))
}

// KjIsNil applies the IsNil predicate on the "kj" field.
func KjIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldKj))
}

// KjNotNil applies the NotNil predicate on the "kj" field.
func KjNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldKj))
}

// KcalEQ applies the EQ predicate on the "kcal" field.
func KcalEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldKcal, v))
}

// KcalNEQ applies the NEQ predicate on the "kcal" field.
func KcalNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldKcal, v))
}

// KcalIn applies the In predicate on the "kcal" field.
func KcalIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldKcal, vs...))
}

// KcalNotIn applies the NotIn predicate on the "kcal" field.
func KcalNotIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldKcal, vs...))
}

// KcalGT applies the GT predicate on the "kcal" field.
func KcalGT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldKcal, v))
}

// KcalGTE applies the GTE predicate on the "kcal" field.
func KcalGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldKcal, v))
}

// KcalLT applies the LT predicate on the "kcal" field.
func KcalLT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldKcal, v))
}

// KcalLTE applies the LTE predicate on the "kcal" field.
func KcalLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldKcal, v))
}

// KcalIsNil applies the IsNil predicate on the "kcal" field.
func KcalIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldKcal))
}

// KcalNotNil applies the NotNil predicate on the "kcal" field.
func KcalNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldKcal))
}

// FatEQ applies the EQ predicate on the "fat" field.
func FatEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldFat, v))
}

// FatNEQ applies the NEQ predicate on the "fat" field.
func FatNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldFat, v))
}

// FatIn applies the In predicate on the "fat" field.
func FatIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldFat, vs...))
}

// FatNotIn applies the NotIn predicate on the "fat" field.
func FatNotIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldFat, vs...))
}

// FatGT applies the GT predicate on the "fat" field.
func FatGT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldFat, v))
}

// FatGTE applies the GTE predicate on the "fat" field.
func FatGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldFat, v))
}

// FatLT applies the LT predicate on the "fat" field.
func FatLT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldFat, v))
}

// FatLTE applies the LTE predicate on the "fat" field.
func FatLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldFat, v))
}

// FatIsNil applies the IsNil predicate on the "fat" field.
func FatIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldFat))
}

// FatNotNil applies the NotNil predicate on the "fat" field.
func FatNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldFat))
}

// SaturatedFatEQ applies the EQ predicate on the "saturated_fat" field.
func SaturatedFatEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldSaturatedFat, v))
}

// SaturatedFatNEQ applies the NEQ predicate on the "saturated_fat" field.
func SaturatedFatNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldSaturatedFat, v))
}

// SaturatedFatIn applies the In predicate on the "saturated_fat" field.
func SaturatedFatIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldSaturatedFat, vs...))
}

// SaturatedFatNotIn applies the NotIn predicate on the "saturated_fat" field.
func SaturatedFatNotIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldSaturatedFat, vs...))
}

// SaturatedFatGT applies the GT predicate on the "saturated_fat" field.
func SaturatedFatGT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldSaturatedFat, v))
}

// SaturatedFatGTE applies the GTE predicate on the "saturated_fat" field.
func SaturatedFatGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldSaturatedFat, v))
}

// SaturatedFatLT applies the LT predicate on the "saturated_fat" field.
func SaturatedFatLT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldSaturatedFat, v))
}

// SaturatedFatLTE applies the LTE predicate on the "saturated_fat" field.
func SaturatedFatLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldSaturatedFat, v))
}

// SaturatedFatIsNil applies the IsNil predicate on the "saturated_fat" field.
func SaturatedFatIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldSaturatedFat))
}

// SaturatedFatNotNil applies the NotNil predicate on the "saturated_fat" field.
func SaturatedFatNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldSaturatedFat))
}

// CarbohydratesEQ applies the EQ predicate on the "carbohydrates" field.
func CarbohydratesEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldCarbohydrates, v))
}

// CarbohydratesNEQ applies the NEQ predicate on the "carbohydrates" field.
func CarbohydratesNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldCarbohydrates, v))
}

// CarbohydratesIn applies the In predicate on the "carbohydrates" field.
func CarbohydratesIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldCarbohydrates, vs...))
}

// CarbohydratesNotIn applies the NotIn predicate on the "carbohydrates" field.
func CarbohydratesNotIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldCarbohydrates, vs...))
}

// CarbohydratesGT applies the GT predicate on the "carbohydrates" field.
func CarbohydratesGT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldCarbohydrates, v))
}

// CarbohydratesGTE applies the GTE predicate on the "carbohydrates" field.
func CarbohydratesGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldCarbohydrates, v))
}

// CarbohydratesLT applies the LT predicate on the "carbohydrates" field.
func CarbohydratesLT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldCarbohydrates, v))
}

// CarbohydratesLTE applies the LTE predicate on the "carbohydrates" field.
func CarbohydratesLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldCarbohydrates, v))
}

// CarbohydratesIsNil applies the IsNil predicate on the "carbohydrates" field.
func CarbohydratesIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldCarbohydrates))
}

// CarbohydratesNotNil applies the NotNil predicate on the "carbohydrates" field.
func CarbohydratesNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldCarbohydrates))
}

// SugarEQ applies the EQ predicate on the "sugar" field.
func SugarEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldSugar, v))
}

// SugarNEQ applies the NEQ predicate on the "sugar" field.
func SugarNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldSugar, v))
}

// SugarIn applies the In predicate on the "sugar" field.
func SugarIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldSugar, vs...))
}

// SugarNotIn applies the NotIn predicate on the "sugar" field.
func SugarNotIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldSugar, vs...))
}

// SugarGT applies the GT predicate on the "sugar" field.
func SugarGT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldSugar, v))
}

// SugarGTE applies the GTE predicate on the "sugar" field.
func SugarGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldSugar, v))
}

// SugarLT applies the LT predicate on the "sugar" field.
func SugarLT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldSugar, v))
}

// SugarLTE applies the LTE predicate on the "sugar" field.
func SugarLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldSugar, v))
}

// SugarIsNil applies the IsNil predicate on the "sugar" field.
func SugarIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldSugar))
}

// SugarNotNil applies the NotNil predicate on the "sugar" field.
func SugarNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldSugar))
}

// FiberEQ applies the EQ predicate on the "fiber" field.
func FiberEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldFiber, v))
}

// FiberNEQ applies the NEQ predicate on the "fiber" field.
func FiberNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldFiber, v))
}

// FiberIn applies the In predicate on the "fiber" field.
func FiberIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldFiber, vs...))
}

// FiberNotIn applies the NotIn predicate on the "fiber" field.
func FiberNotIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldFiber, vs...))
}

// FiberGT applies the GT predicate on the "fiber" field.
func FiberGT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldFiber, v))
}

// FiberGTE applies the GTE predicate on the "fiber" field.
func FiberGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldFiber, v))
}

// FiberLT applies the LT predicate on the "fiber" field.
func FiberLT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldFiber, v))
}

// FiberLTE applies the LTE predicate on the "fiber" field.
func FiberLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldFiber, v))
}

// FiberIsNil applies the IsNil predicate on the "fiber" field.
func FiberIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldFiber))
}

// FiberNotNil applies the NotNil predicate on the "fiber" field.
func FiberNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldFiber))
}

// ProteinEQ applies the EQ predicate on the "protein" field.
func ProteinEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldProtein, v))
}

// ProteinNEQ applies the NEQ predicate on the "protein" field.
func ProteinNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldProtein, v))
}

// ProteinIn applies the In predicate on the "protein" field.
func ProteinIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldProtein, vs...))
}

// ProteinNotIn applies the NotIn predicate on the "protein" field.
func ProteinNotIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldProtein, vs...))
}

// ProteinGT applies the GT predicate on the "protein" field.
func ProteinGT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldProtein, v))
}

// ProteinGTE applies the GTE predicate on the "protein" field.
func ProteinGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldProtein, v))
}

// ProteinLT applies the LT predicate on the "protein" field.
func ProteinLT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldProtein, v))
}

// ProteinLTE applies the LTE predicate on the "protein" field.
func ProteinLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldProtein, v))
}

// ProteinIsNil applies the IsNil predicate on the "protein" field.
func ProteinIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldProtein))
}

// ProteinNotNil applies the NotNil predicate on the "protein" field.
func ProteinNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldProtein))
}

// SaltEQ applies the EQ predicate on the "salt" field.
func SaltEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldSalt, v))
}

// SaltNEQ applies the NEQ predicate on the "salt" field.
func SaltNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldSalt, v))
}

// SaltIn applies the In predicate on the "salt" field.
func SaltIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldSalt, vs...))
}

// SaltNotIn applies the NotIn predicate on the "salt" field.
func SaltNotIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldSalt, vs...))
}

// SaltGT applies the GT predicate on the "salt" field.
func SaltGT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldSalt, v))
}

// SaltGTE applies the GTE predicate on the "salt" field.
func SaltGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldSalt, v))
}

// SaltLT applies the LT predicate on the "salt" field.
func SaltLT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldSalt, v))
}

// SaltLTE applies the LTE predicate on the "salt" field.
func SaltLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldSalt, v))
}

// SaltIsNil applies the IsNil predicate on the "salt" field.
func SaltIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldSalt))
}

// SaltNotNil applies the NotNil predicate on the "salt" field.
func SaltNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldSalt))
}

// PriceStudentEQ applies the EQ predicate on the "price_student" field.
func PriceStudentEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldPriceStudent, v))
}

// PriceStudentNEQ applies the NEQ predicate on the "price_student" field.
func PriceStudentNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldPriceStudent, v))
}

// PriceStudentIn applies the In predicate on the "price_student" field.
func PriceStudentIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldPriceStudent, vs...))
}

// PriceStudentNotIn applies the NotIn predicate on the "price_student" field.
func PriceStudentNotIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldPriceStudent, vs...))
}

// PriceStudentGT applies the GT predicate on the "price_student" field.
func PriceStudentGT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldPriceStudent, v))
}

// PriceStudentGTE applies the GTE predicate on the "price_student" field.
func PriceStudentGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldPriceStudent, v))
}

// PriceStudentLT applies the LT predicate on the "price_student" field.
func PriceStudentLT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldPriceStudent, v))
}

// PriceStudentLTE applies the LTE predicate on the "price_student" field.
func PriceStudentLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldPriceStudent, v))
}

// PriceStudentIsNil applies the IsNil predicate on the "price_student" field.
func PriceStudentIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldPriceStudent))
}

// PriceStudentNotNil applies the NotNil predicate on the "price_student" field.
func PriceStudentNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldPriceStudent))
}

// PriceStaffEQ applies the EQ predicate on the "price_staff" field.
func PriceStaffEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldPriceStaff, v))
}

// PriceStaffNEQ applies the NEQ predicate on the "price_staff" field.
func PriceStaffNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldPriceStaff, v))
}

// PriceStaffIn applies the In predicate on the "price_staff" field.
func PriceStaffIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldPriceStaff, vs...))
}

// PriceStaffNotIn applies the NotIn predicate on the "price_staff" field.
func PriceStaffNotIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldPriceStaff, vs...))
}

// PriceStaffGT applies the GT predicate on the "price_staff" field.
func PriceStaffGT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldPriceStaff, v))
}

// PriceStaffGTE applies the GTE predicate on the "price_staff" field.
func PriceStaffGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldPriceStaff, v))
}

// PriceStaffLT applies the LT predicate on the "price_staff" field.
func PriceStaffLT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldPriceStaff, v))
}

// PriceStaffLTE applies the LTE predicate on the "price_staff" field.
func PriceStaffLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldPriceStaff, v))
}

// PriceStaffIsNil applies the IsNil predicate on the "price_staff" field.
func PriceStaffIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldPriceStaff))
}

// PriceStaffNotNil applies the NotNil predicate on the "price_staff" field.
func PriceStaffNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldPriceStaff))
}

// PriceGuestEQ applies the EQ predicate on the "price_guest" field.
func PriceGuestEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldPriceGuest, v))
}

// PriceGuestNEQ applies the NEQ predicate on the "price_guest" field.
func PriceGuestNEQ(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldPriceGuest, v))
}

// PriceGuestIn applies the In predicate on the "price_guest" field.
func PriceGuestIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldPriceGuest, vs...))
}

// PriceGuestNotIn applies the NotIn predicate on the "price_guest" field.
func PriceGuestNotIn(vs ...int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldPriceGuest, vs...))
}

// PriceGuestGT applies the GT predicate on the "price_guest" field.
func PriceGuestGT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldPriceGuest, v))
}

// PriceGuestGTE applies the GTE predicate on the "price_guest" field.
func PriceGuestGTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldPriceGuest, v))
}

// PriceGuestLT applies the LT predicate on the "price_guest" field.
func PriceGuestLT(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldPriceGuest, v))
}

// PriceGuestLTE applies the LTE predicate on the "price_guest" field.
func PriceGuestLTE(v int) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldPriceGuest, v))
}

// PriceGuestIsNil applies the IsNil predicate on the "price_guest" field.
func PriceGuestIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldPriceGuest))
}

// PriceGuestNotNil applies the NotNil predicate on the "price_guest" field.
func PriceGuestNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldPriceGuest))
}

// NotAvailableAfterEQ applies the EQ predicate on the "not_available_after" field.
func NotAvailableAfterEQ(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldEQ(FieldNotAvailableAfter, v))
}

// NotAvailableAfterNEQ applies the NEQ predicate on the "not_available_after" field.
func NotAvailableAfterNEQ(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNEQ(FieldNotAvailableAfter, v))
}

// NotAvailableAfterIn applies the In predicate on the "not_available_after" field.
func NotAvailableAfterIn(vs ...time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIn(FieldNotAvailableAfter, vs...))
}

// NotAvailableAfterNotIn applies the NotIn predicate on the "not_available_after" field.
func NotAvailableAfterNotIn(vs ...time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotIn(FieldNotAvailableAfter, vs...))
}

// NotAvailableAfterGT applies the GT predicate on the "not_available_after" field.
func NotAvailableAfterGT(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGT(FieldNotAvailableAfter, v))
}

// NotAvailableAfterGTE applies the GTE predicate on the "not_available_after" field.
func NotAvailableAfterGTE(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldGTE(FieldNotAvailableAfter, v))
}

// NotAvailableAfterLT applies the LT predicate on the "not_available_after" field.
func NotAvailableAfterLT(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLT(FieldNotAvailableAfter, v))
}

// NotAvailableAfterLTE applies the LTE predicate on the "not_available_after" field.
func NotAvailableAfterLTE(v time.Time) predicate.Occurrence {
	return predicate.Occurrence(sql.FieldLTE(FieldNotAvailableAfter, v))
}

// NotAvailableAfterIsNil applies the IsNil predicate on the "not_available_after" field.
func NotAvailableAfterIsNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldIsNull(FieldNotAvailableAfter))
}

// NotAvailableAfterNotNil applies the NotNil predicate on the "not_available_after" field.
func NotAvailableAfterNotNil() predicate.Occurrence {
	return predicate.Occurrence(sql.FieldNotNull(FieldNotAvailableAfter))
}

// HasLocation applies the HasEdge predicate on the "location" edge.
func HasLocation() predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LocationTable, LocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationWith applies the HasEdge predicate on the "location" edge with a given conditions (other predicates).
func HasLocationWith(preds ...predicate.Location) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := newLocationStep()
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
			sqlgraph.Edge(sqlgraph.M2O, true, DishTable, DishColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDishWith applies the HasEdge predicate on the "dish" edge with a given conditions (other predicates).
func HasDishWith(preds ...predicate.Dish) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := newDishStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := newTagsStep()
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
			sqlgraph.Edge(sqlgraph.M2M, false, SideDishesTable, SideDishesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSideDishesWith applies the HasEdge predicate on the "side_dishes" edge with a given conditions (other predicates).
func HasSideDishesWith(preds ...predicate.Dish) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := newSideDishesStep()
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
			sqlgraph.Edge(sqlgraph.O2M, false, ReviewsTable, ReviewsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReviewsWith applies the HasEdge predicate on the "reviews" edge with a given conditions (other predicates).
func HasReviewsWith(preds ...predicate.Review) predicate.Occurrence {
	return predicate.Occurrence(func(s *sql.Selector) {
		step := newReviewsStep()
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
