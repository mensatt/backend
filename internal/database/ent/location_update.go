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
	"github.com/mensatt/backend/internal/database/ent/location"
	"github.com/mensatt/backend/internal/database/ent/occurrence"
	"github.com/mensatt/backend/internal/database/ent/predicate"
)

// LocationUpdate is the builder for updating Location entities.
type LocationUpdate struct {
	config
	hooks    []Hook
	mutation *LocationMutation
}

// Where appends a list predicates to the LocationUpdate builder.
func (lu *LocationUpdate) Where(ps ...predicate.Location) *LocationUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetExternalID sets the "external_id" field.
func (lu *LocationUpdate) SetExternalID(i int) *LocationUpdate {
	lu.mutation.ResetExternalID()
	lu.mutation.SetExternalID(i)
	return lu
}

// AddExternalID adds i to the "external_id" field.
func (lu *LocationUpdate) AddExternalID(i int) *LocationUpdate {
	lu.mutation.AddExternalID(i)
	return lu
}

// SetName sets the "name" field.
func (lu *LocationUpdate) SetName(s string) *LocationUpdate {
	lu.mutation.SetName(s)
	return lu
}

// AddOccurrenceIDs adds the "occurrences" edge to the Occurrence entity by IDs.
func (lu *LocationUpdate) AddOccurrenceIDs(ids ...uuid.UUID) *LocationUpdate {
	lu.mutation.AddOccurrenceIDs(ids...)
	return lu
}

// AddOccurrences adds the "occurrences" edges to the Occurrence entity.
func (lu *LocationUpdate) AddOccurrences(o ...*Occurrence) *LocationUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return lu.AddOccurrenceIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (lu *LocationUpdate) Mutation() *LocationMutation {
	return lu.mutation
}

// ClearOccurrences clears all "occurrences" edges to the Occurrence entity.
func (lu *LocationUpdate) ClearOccurrences() *LocationUpdate {
	lu.mutation.ClearOccurrences()
	return lu
}

// RemoveOccurrenceIDs removes the "occurrences" edge to Occurrence entities by IDs.
func (lu *LocationUpdate) RemoveOccurrenceIDs(ids ...uuid.UUID) *LocationUpdate {
	lu.mutation.RemoveOccurrenceIDs(ids...)
	return lu
}

// RemoveOccurrences removes "occurrences" edges to Occurrence entities.
func (lu *LocationUpdate) RemoveOccurrences(o ...*Occurrence) *LocationUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return lu.RemoveOccurrenceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LocationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		if err = lu.check(); err != nil {
			return 0, err
		}
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LocationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lu.check(); err != nil {
				return 0, err
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LocationUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LocationUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LocationUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LocationUpdate) check() error {
	if v, ok := lu.mutation.Name(); ok {
		if err := location.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Location.name": %w`, err)}
		}
	}
	return nil
}

func (lu *LocationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   location.Table,
			Columns: location.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: location.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.ExternalID(); ok {
		_spec.SetField(location.FieldExternalID, field.TypeInt, value)
	}
	if value, ok := lu.mutation.AddedExternalID(); ok {
		_spec.AddField(location.FieldExternalID, field.TypeInt, value)
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(location.FieldName, field.TypeString, value)
	}
	if lu.mutation.OccurrencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.OccurrencesTable,
			Columns: []string{location.OccurrencesColumn},
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
	if nodes := lu.mutation.RemovedOccurrencesIDs(); len(nodes) > 0 && !lu.mutation.OccurrencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.OccurrencesTable,
			Columns: []string{location.OccurrencesColumn},
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
	if nodes := lu.mutation.OccurrencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.OccurrencesTable,
			Columns: []string{location.OccurrencesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LocationUpdateOne is the builder for updating a single Location entity.
type LocationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LocationMutation
}

// SetExternalID sets the "external_id" field.
func (luo *LocationUpdateOne) SetExternalID(i int) *LocationUpdateOne {
	luo.mutation.ResetExternalID()
	luo.mutation.SetExternalID(i)
	return luo
}

// AddExternalID adds i to the "external_id" field.
func (luo *LocationUpdateOne) AddExternalID(i int) *LocationUpdateOne {
	luo.mutation.AddExternalID(i)
	return luo
}

// SetName sets the "name" field.
func (luo *LocationUpdateOne) SetName(s string) *LocationUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// AddOccurrenceIDs adds the "occurrences" edge to the Occurrence entity by IDs.
func (luo *LocationUpdateOne) AddOccurrenceIDs(ids ...uuid.UUID) *LocationUpdateOne {
	luo.mutation.AddOccurrenceIDs(ids...)
	return luo
}

// AddOccurrences adds the "occurrences" edges to the Occurrence entity.
func (luo *LocationUpdateOne) AddOccurrences(o ...*Occurrence) *LocationUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return luo.AddOccurrenceIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (luo *LocationUpdateOne) Mutation() *LocationMutation {
	return luo.mutation
}

// ClearOccurrences clears all "occurrences" edges to the Occurrence entity.
func (luo *LocationUpdateOne) ClearOccurrences() *LocationUpdateOne {
	luo.mutation.ClearOccurrences()
	return luo
}

// RemoveOccurrenceIDs removes the "occurrences" edge to Occurrence entities by IDs.
func (luo *LocationUpdateOne) RemoveOccurrenceIDs(ids ...uuid.UUID) *LocationUpdateOne {
	luo.mutation.RemoveOccurrenceIDs(ids...)
	return luo
}

// RemoveOccurrences removes "occurrences" edges to Occurrence entities.
func (luo *LocationUpdateOne) RemoveOccurrences(o ...*Occurrence) *LocationUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return luo.RemoveOccurrenceIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LocationUpdateOne) Select(field string, fields ...string) *LocationUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Location entity.
func (luo *LocationUpdateOne) Save(ctx context.Context) (*Location, error) {
	var (
		err  error
		node *Location
	)
	if len(luo.hooks) == 0 {
		if err = luo.check(); err != nil {
			return nil, err
		}
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LocationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luo.check(); err != nil {
				return nil, err
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, luo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Location)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LocationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LocationUpdateOne) SaveX(ctx context.Context) *Location {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LocationUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LocationUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LocationUpdateOne) check() error {
	if v, ok := luo.mutation.Name(); ok {
		if err := location.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Location.name": %w`, err)}
		}
	}
	return nil
}

func (luo *LocationUpdateOne) sqlSave(ctx context.Context) (_node *Location, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   location.Table,
			Columns: location.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: location.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Location.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, location.FieldID)
		for _, f := range fields {
			if !location.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != location.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.ExternalID(); ok {
		_spec.SetField(location.FieldExternalID, field.TypeInt, value)
	}
	if value, ok := luo.mutation.AddedExternalID(); ok {
		_spec.AddField(location.FieldExternalID, field.TypeInt, value)
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(location.FieldName, field.TypeString, value)
	}
	if luo.mutation.OccurrencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.OccurrencesTable,
			Columns: []string{location.OccurrencesColumn},
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
	if nodes := luo.mutation.RemovedOccurrencesIDs(); len(nodes) > 0 && !luo.mutation.OccurrencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.OccurrencesTable,
			Columns: []string{location.OccurrencesColumn},
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
	if nodes := luo.mutation.OccurrencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.OccurrencesTable,
			Columns: []string{location.OccurrencesColumn},
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
	_node = &Location{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}