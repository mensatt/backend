// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent/dish"
	"github.com/mensatt/backend/internal/database/ent/dishalias"
	"github.com/mensatt/backend/internal/database/ent/occurrence"
	"github.com/mensatt/backend/internal/database/ent/predicate"
)

// DishUpdate is the builder for updating Dish entities.
type DishUpdate struct {
	config
	hooks    []Hook
	mutation *DishMutation
}

// Where appends a list predicates to the DishUpdate builder.
func (du *DishUpdate) Where(ps ...predicate.Dish) *DishUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetNameDe sets the "name_de" field.
func (du *DishUpdate) SetNameDe(s string) *DishUpdate {
	du.mutation.SetNameDe(s)
	return du
}

// SetNameEn sets the "name_en" field.
func (du *DishUpdate) SetNameEn(s string) *DishUpdate {
	du.mutation.SetNameEn(s)
	return du
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (du *DishUpdate) SetNillableNameEn(s *string) *DishUpdate {
	if s != nil {
		du.SetNameEn(*s)
	}
	return du
}

// ClearNameEn clears the value of the "name_en" field.
func (du *DishUpdate) ClearNameEn() *DishUpdate {
	du.mutation.ClearNameEn()
	return du
}

// AddDishOccurrenceIDs adds the "dish_occurrences" edge to the Occurrence entity by IDs.
func (du *DishUpdate) AddDishOccurrenceIDs(ids ...uuid.UUID) *DishUpdate {
	du.mutation.AddDishOccurrenceIDs(ids...)
	return du
}

// AddDishOccurrences adds the "dish_occurrences" edges to the Occurrence entity.
func (du *DishUpdate) AddDishOccurrences(o ...*Occurrence) *DishUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return du.AddDishOccurrenceIDs(ids...)
}

// AddAliasIDs adds the "aliases" edge to the DishAlias entity by IDs.
func (du *DishUpdate) AddAliasIDs(ids ...string) *DishUpdate {
	du.mutation.AddAliasIDs(ids...)
	return du
}

// AddAliases adds the "aliases" edges to the DishAlias entity.
func (du *DishUpdate) AddAliases(d ...*DishAlias) *DishUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddAliasIDs(ids...)
}

// AddSideDishOccurrenceIDs adds the "side_dish_occurrence" edge to the Occurrence entity by IDs.
func (du *DishUpdate) AddSideDishOccurrenceIDs(ids ...uuid.UUID) *DishUpdate {
	du.mutation.AddSideDishOccurrenceIDs(ids...)
	return du
}

// AddSideDishOccurrence adds the "side_dish_occurrence" edges to the Occurrence entity.
func (du *DishUpdate) AddSideDishOccurrence(o ...*Occurrence) *DishUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return du.AddSideDishOccurrenceIDs(ids...)
}

// Mutation returns the DishMutation object of the builder.
func (du *DishUpdate) Mutation() *DishMutation {
	return du.mutation
}

// ClearDishOccurrences clears all "dish_occurrences" edges to the Occurrence entity.
func (du *DishUpdate) ClearDishOccurrences() *DishUpdate {
	du.mutation.ClearDishOccurrences()
	return du
}

// RemoveDishOccurrenceIDs removes the "dish_occurrences" edge to Occurrence entities by IDs.
func (du *DishUpdate) RemoveDishOccurrenceIDs(ids ...uuid.UUID) *DishUpdate {
	du.mutation.RemoveDishOccurrenceIDs(ids...)
	return du
}

// RemoveDishOccurrences removes "dish_occurrences" edges to Occurrence entities.
func (du *DishUpdate) RemoveDishOccurrences(o ...*Occurrence) *DishUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return du.RemoveDishOccurrenceIDs(ids...)
}

// ClearAliases clears all "aliases" edges to the DishAlias entity.
func (du *DishUpdate) ClearAliases() *DishUpdate {
	du.mutation.ClearAliases()
	return du
}

// RemoveAliasIDs removes the "aliases" edge to DishAlias entities by IDs.
func (du *DishUpdate) RemoveAliasIDs(ids ...string) *DishUpdate {
	du.mutation.RemoveAliasIDs(ids...)
	return du
}

// RemoveAliases removes "aliases" edges to DishAlias entities.
func (du *DishUpdate) RemoveAliases(d ...*DishAlias) *DishUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveAliasIDs(ids...)
}

// ClearSideDishOccurrence clears all "side_dish_occurrence" edges to the Occurrence entity.
func (du *DishUpdate) ClearSideDishOccurrence() *DishUpdate {
	du.mutation.ClearSideDishOccurrence()
	return du
}

// RemoveSideDishOccurrenceIDs removes the "side_dish_occurrence" edge to Occurrence entities by IDs.
func (du *DishUpdate) RemoveSideDishOccurrenceIDs(ids ...uuid.UUID) *DishUpdate {
	du.mutation.RemoveSideDishOccurrenceIDs(ids...)
	return du
}

// RemoveSideDishOccurrence removes "side_dish_occurrence" edges to Occurrence entities.
func (du *DishUpdate) RemoveSideDishOccurrence(o ...*Occurrence) *DishUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return du.RemoveSideDishOccurrenceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DishUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DishMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DishUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DishUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DishUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DishUpdate) check() error {
	if v, ok := du.mutation.NameDe(); ok {
		if err := dish.NameDeValidator(v); err != nil {
			return &ValidationError{Name: "name_de", err: fmt.Errorf(`ent: validator failed for field "Dish.name_de": %w`, err)}
		}
	}
	if v, ok := du.mutation.NameEn(); ok {
		if err := dish.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf(`ent: validator failed for field "Dish.name_en": %w`, err)}
		}
	}
	return nil
}

func (du *DishUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dish.Table,
			Columns: dish.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: dish.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.NameDe(); ok {
		_spec.SetField(dish.FieldNameDe, field.TypeString, value)
	}
	if value, ok := du.mutation.NameEn(); ok {
		_spec.SetField(dish.FieldNameEn, field.TypeString, value)
	}
	if du.mutation.NameEnCleared() {
		_spec.ClearField(dish.FieldNameEn, field.TypeString)
	}
	if du.mutation.DishOccurrencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.DishOccurrencesTable,
			Columns: []string{dish.DishOccurrencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: occurrence.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedDishOccurrencesIDs(); len(nodes) > 0 && !du.mutation.DishOccurrencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.DishOccurrencesTable,
			Columns: []string{dish.DishOccurrencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: occurrence.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DishOccurrencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.DishOccurrencesTable,
			Columns: []string{dish.DishOccurrencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: occurrence.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.AliasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.AliasesTable,
			Columns: []string{dish.AliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dishalias.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedAliasesIDs(); len(nodes) > 0 && !du.mutation.AliasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.AliasesTable,
			Columns: []string{dish.AliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dishalias.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.AliasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.AliasesTable,
			Columns: []string{dish.AliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dishalias.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.SideDishOccurrenceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dish.SideDishOccurrenceTable,
			Columns: dish.SideDishOccurrencePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: occurrence.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedSideDishOccurrenceIDs(); len(nodes) > 0 && !du.mutation.SideDishOccurrenceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dish.SideDishOccurrenceTable,
			Columns: dish.SideDishOccurrencePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: occurrence.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.SideDishOccurrenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dish.SideDishOccurrenceTable,
			Columns: dish.SideDishOccurrencePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: occurrence.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dish.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DishUpdateOne is the builder for updating a single Dish entity.
type DishUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DishMutation
}

// SetNameDe sets the "name_de" field.
func (duo *DishUpdateOne) SetNameDe(s string) *DishUpdateOne {
	duo.mutation.SetNameDe(s)
	return duo
}

// SetNameEn sets the "name_en" field.
func (duo *DishUpdateOne) SetNameEn(s string) *DishUpdateOne {
	duo.mutation.SetNameEn(s)
	return duo
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (duo *DishUpdateOne) SetNillableNameEn(s *string) *DishUpdateOne {
	if s != nil {
		duo.SetNameEn(*s)
	}
	return duo
}

// ClearNameEn clears the value of the "name_en" field.
func (duo *DishUpdateOne) ClearNameEn() *DishUpdateOne {
	duo.mutation.ClearNameEn()
	return duo
}

// AddDishOccurrenceIDs adds the "dish_occurrences" edge to the Occurrence entity by IDs.
func (duo *DishUpdateOne) AddDishOccurrenceIDs(ids ...uuid.UUID) *DishUpdateOne {
	duo.mutation.AddDishOccurrenceIDs(ids...)
	return duo
}

// AddDishOccurrences adds the "dish_occurrences" edges to the Occurrence entity.
func (duo *DishUpdateOne) AddDishOccurrences(o ...*Occurrence) *DishUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return duo.AddDishOccurrenceIDs(ids...)
}

// AddAliasIDs adds the "aliases" edge to the DishAlias entity by IDs.
func (duo *DishUpdateOne) AddAliasIDs(ids ...string) *DishUpdateOne {
	duo.mutation.AddAliasIDs(ids...)
	return duo
}

// AddAliases adds the "aliases" edges to the DishAlias entity.
func (duo *DishUpdateOne) AddAliases(d ...*DishAlias) *DishUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddAliasIDs(ids...)
}

// AddSideDishOccurrenceIDs adds the "side_dish_occurrence" edge to the Occurrence entity by IDs.
func (duo *DishUpdateOne) AddSideDishOccurrenceIDs(ids ...uuid.UUID) *DishUpdateOne {
	duo.mutation.AddSideDishOccurrenceIDs(ids...)
	return duo
}

// AddSideDishOccurrence adds the "side_dish_occurrence" edges to the Occurrence entity.
func (duo *DishUpdateOne) AddSideDishOccurrence(o ...*Occurrence) *DishUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return duo.AddSideDishOccurrenceIDs(ids...)
}

// Mutation returns the DishMutation object of the builder.
func (duo *DishUpdateOne) Mutation() *DishMutation {
	return duo.mutation
}

// ClearDishOccurrences clears all "dish_occurrences" edges to the Occurrence entity.
func (duo *DishUpdateOne) ClearDishOccurrences() *DishUpdateOne {
	duo.mutation.ClearDishOccurrences()
	return duo
}

// RemoveDishOccurrenceIDs removes the "dish_occurrences" edge to Occurrence entities by IDs.
func (duo *DishUpdateOne) RemoveDishOccurrenceIDs(ids ...uuid.UUID) *DishUpdateOne {
	duo.mutation.RemoveDishOccurrenceIDs(ids...)
	return duo
}

// RemoveDishOccurrences removes "dish_occurrences" edges to Occurrence entities.
func (duo *DishUpdateOne) RemoveDishOccurrences(o ...*Occurrence) *DishUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return duo.RemoveDishOccurrenceIDs(ids...)
}

// ClearAliases clears all "aliases" edges to the DishAlias entity.
func (duo *DishUpdateOne) ClearAliases() *DishUpdateOne {
	duo.mutation.ClearAliases()
	return duo
}

// RemoveAliasIDs removes the "aliases" edge to DishAlias entities by IDs.
func (duo *DishUpdateOne) RemoveAliasIDs(ids ...string) *DishUpdateOne {
	duo.mutation.RemoveAliasIDs(ids...)
	return duo
}

// RemoveAliases removes "aliases" edges to DishAlias entities.
func (duo *DishUpdateOne) RemoveAliases(d ...*DishAlias) *DishUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveAliasIDs(ids...)
}

// ClearSideDishOccurrence clears all "side_dish_occurrence" edges to the Occurrence entity.
func (duo *DishUpdateOne) ClearSideDishOccurrence() *DishUpdateOne {
	duo.mutation.ClearSideDishOccurrence()
	return duo
}

// RemoveSideDishOccurrenceIDs removes the "side_dish_occurrence" edge to Occurrence entities by IDs.
func (duo *DishUpdateOne) RemoveSideDishOccurrenceIDs(ids ...uuid.UUID) *DishUpdateOne {
	duo.mutation.RemoveSideDishOccurrenceIDs(ids...)
	return duo
}

// RemoveSideDishOccurrence removes "side_dish_occurrence" edges to Occurrence entities.
func (duo *DishUpdateOne) RemoveSideDishOccurrence(o ...*Occurrence) *DishUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return duo.RemoveSideDishOccurrenceIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DishUpdateOne) Select(field string, fields ...string) *DishUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Dish entity.
func (duo *DishUpdateOne) Save(ctx context.Context) (*Dish, error) {
	var (
		err  error
		node *Dish
	)
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DishMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, duo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Dish)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DishMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DishUpdateOne) SaveX(ctx context.Context) *Dish {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DishUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DishUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DishUpdateOne) check() error {
	if v, ok := duo.mutation.NameDe(); ok {
		if err := dish.NameDeValidator(v); err != nil {
			return &ValidationError{Name: "name_de", err: fmt.Errorf(`ent: validator failed for field "Dish.name_de": %w`, err)}
		}
	}
	if v, ok := duo.mutation.NameEn(); ok {
		if err := dish.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf(`ent: validator failed for field "Dish.name_en": %w`, err)}
		}
	}
	return nil
}

func (duo *DishUpdateOne) sqlSave(ctx context.Context) (_node *Dish, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dish.Table,
			Columns: dish.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: dish.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Dish.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dish.FieldID)
		for _, f := range fields {
			if !dish.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dish.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.NameDe(); ok {
		_spec.SetField(dish.FieldNameDe, field.TypeString, value)
	}
	if value, ok := duo.mutation.NameEn(); ok {
		_spec.SetField(dish.FieldNameEn, field.TypeString, value)
	}
	if duo.mutation.NameEnCleared() {
		_spec.ClearField(dish.FieldNameEn, field.TypeString)
	}
	if duo.mutation.DishOccurrencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.DishOccurrencesTable,
			Columns: []string{dish.DishOccurrencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: occurrence.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedDishOccurrencesIDs(); len(nodes) > 0 && !duo.mutation.DishOccurrencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.DishOccurrencesTable,
			Columns: []string{dish.DishOccurrencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: occurrence.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DishOccurrencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.DishOccurrencesTable,
			Columns: []string{dish.DishOccurrencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: occurrence.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.AliasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.AliasesTable,
			Columns: []string{dish.AliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dishalias.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedAliasesIDs(); len(nodes) > 0 && !duo.mutation.AliasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.AliasesTable,
			Columns: []string{dish.AliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dishalias.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.AliasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.AliasesTable,
			Columns: []string{dish.AliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dishalias.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.SideDishOccurrenceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dish.SideDishOccurrenceTable,
			Columns: dish.SideDishOccurrencePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: occurrence.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedSideDishOccurrenceIDs(); len(nodes) > 0 && !duo.mutation.SideDishOccurrenceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dish.SideDishOccurrenceTable,
			Columns: dish.SideDishOccurrencePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: occurrence.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.SideDishOccurrenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dish.SideDishOccurrenceTable,
			Columns: dish.SideDishOccurrencePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: occurrence.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Dish{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dish.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
