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
	"github.com/mensatt/backend/internal/database/ent/occurrence"
	"github.com/mensatt/backend/internal/database/ent/predicate"
	"github.com/mensatt/backend/internal/database/ent/tag"
	"github.com/mensatt/backend/internal/database/schema"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	hooks    []Hook
	mutation *TagMutation
}

// Where appends a list predicates to the TagUpdate builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetKey sets the "key" field.
func (tu *TagUpdate) SetKey(s string) *TagUpdate {
	tu.mutation.SetKey(s)
	return tu
}

// SetName sets the "name" field.
func (tu *TagUpdate) SetName(s string) *TagUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TagUpdate) SetDescription(s string) *TagUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetShortName sets the "short_name" field.
func (tu *TagUpdate) SetShortName(s string) *TagUpdate {
	tu.mutation.SetShortName(s)
	return tu
}

// SetNillableShortName sets the "short_name" field if the given value is not nil.
func (tu *TagUpdate) SetNillableShortName(s *string) *TagUpdate {
	if s != nil {
		tu.SetShortName(*s)
	}
	return tu
}

// ClearShortName clears the value of the "short_name" field.
func (tu *TagUpdate) ClearShortName() *TagUpdate {
	tu.mutation.ClearShortName()
	return tu
}

// SetPriority sets the "priority" field.
func (tu *TagUpdate) SetPriority(sp schema.TagPriority) *TagUpdate {
	tu.mutation.SetPriority(sp)
	return tu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tu *TagUpdate) SetNillablePriority(sp *schema.TagPriority) *TagUpdate {
	if sp != nil {
		tu.SetPriority(*sp)
	}
	return tu
}

// SetIsAllergy sets the "is_allergy" field.
func (tu *TagUpdate) SetIsAllergy(b bool) *TagUpdate {
	tu.mutation.SetIsAllergy(b)
	return tu
}

// SetNillableIsAllergy sets the "is_allergy" field if the given value is not nil.
func (tu *TagUpdate) SetNillableIsAllergy(b *bool) *TagUpdate {
	if b != nil {
		tu.SetIsAllergy(*b)
	}
	return tu
}

// AddOccurrenceIDs adds the "occurrences" edge to the Occurrence entity by IDs.
func (tu *TagUpdate) AddOccurrenceIDs(ids ...uuid.UUID) *TagUpdate {
	tu.mutation.AddOccurrenceIDs(ids...)
	return tu
}

// AddOccurrences adds the "occurrences" edges to the Occurrence entity.
func (tu *TagUpdate) AddOccurrences(o ...*Occurrence) *TagUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return tu.AddOccurrenceIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tu *TagUpdate) Mutation() *TagMutation {
	return tu.mutation
}

// ClearOccurrences clears all "occurrences" edges to the Occurrence entity.
func (tu *TagUpdate) ClearOccurrences() *TagUpdate {
	tu.mutation.ClearOccurrences()
	return tu
}

// RemoveOccurrenceIDs removes the "occurrences" edge to Occurrence entities by IDs.
func (tu *TagUpdate) RemoveOccurrenceIDs(ids ...uuid.UUID) *TagUpdate {
	tu.mutation.RemoveOccurrenceIDs(ids...)
	return tu
}

// RemoveOccurrences removes "occurrences" edges to Occurrence entities.
func (tu *TagUpdate) RemoveOccurrences(o ...*Occurrence) *TagUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return tu.RemoveOccurrenceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TagUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := tag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tag.name": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Description(); ok {
		if err := tag.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Tag.description": %w`, err)}
		}
	}
	if v, ok := tu.mutation.ShortName(); ok {
		if err := tag.ShortNameValidator(v); err != nil {
			return &ValidationError{Name: "short_name", err: fmt.Errorf(`ent: validator failed for field "Tag.short_name": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Priority(); ok {
		if err := tag.PriorityValidator(v); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Tag.priority": %w`, err)}
		}
	}
	return nil
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tag.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Key(); ok {
		_spec.SetField(tag.FieldKey, field.TypeString, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(tag.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(tag.FieldDescription, field.TypeString, value)
	}
	if value, ok := tu.mutation.ShortName(); ok {
		_spec.SetField(tag.FieldShortName, field.TypeString, value)
	}
	if tu.mutation.ShortNameCleared() {
		_spec.ClearField(tag.FieldShortName, field.TypeString)
	}
	if value, ok := tu.mutation.Priority(); ok {
		_spec.SetField(tag.FieldPriority, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.IsAllergy(); ok {
		_spec.SetField(tag.FieldIsAllergy, field.TypeBool, value)
	}
	if tu.mutation.OccurrencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.OccurrencesTable,
			Columns: tag.OccurrencesPrimaryKey,
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
	if nodes := tu.mutation.RemovedOccurrencesIDs(); len(nodes) > 0 && !tu.mutation.OccurrencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.OccurrencesTable,
			Columns: tag.OccurrencesPrimaryKey,
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
	if nodes := tu.mutation.OccurrencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.OccurrencesTable,
			Columns: tag.OccurrencesPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagMutation
}

// SetKey sets the "key" field.
func (tuo *TagUpdateOne) SetKey(s string) *TagUpdateOne {
	tuo.mutation.SetKey(s)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TagUpdateOne) SetName(s string) *TagUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TagUpdateOne) SetDescription(s string) *TagUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetShortName sets the "short_name" field.
func (tuo *TagUpdateOne) SetShortName(s string) *TagUpdateOne {
	tuo.mutation.SetShortName(s)
	return tuo
}

// SetNillableShortName sets the "short_name" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableShortName(s *string) *TagUpdateOne {
	if s != nil {
		tuo.SetShortName(*s)
	}
	return tuo
}

// ClearShortName clears the value of the "short_name" field.
func (tuo *TagUpdateOne) ClearShortName() *TagUpdateOne {
	tuo.mutation.ClearShortName()
	return tuo
}

// SetPriority sets the "priority" field.
func (tuo *TagUpdateOne) SetPriority(sp schema.TagPriority) *TagUpdateOne {
	tuo.mutation.SetPriority(sp)
	return tuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillablePriority(sp *schema.TagPriority) *TagUpdateOne {
	if sp != nil {
		tuo.SetPriority(*sp)
	}
	return tuo
}

// SetIsAllergy sets the "is_allergy" field.
func (tuo *TagUpdateOne) SetIsAllergy(b bool) *TagUpdateOne {
	tuo.mutation.SetIsAllergy(b)
	return tuo
}

// SetNillableIsAllergy sets the "is_allergy" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableIsAllergy(b *bool) *TagUpdateOne {
	if b != nil {
		tuo.SetIsAllergy(*b)
	}
	return tuo
}

// AddOccurrenceIDs adds the "occurrences" edge to the Occurrence entity by IDs.
func (tuo *TagUpdateOne) AddOccurrenceIDs(ids ...uuid.UUID) *TagUpdateOne {
	tuo.mutation.AddOccurrenceIDs(ids...)
	return tuo
}

// AddOccurrences adds the "occurrences" edges to the Occurrence entity.
func (tuo *TagUpdateOne) AddOccurrences(o ...*Occurrence) *TagUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return tuo.AddOccurrenceIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tuo *TagUpdateOne) Mutation() *TagMutation {
	return tuo.mutation
}

// ClearOccurrences clears all "occurrences" edges to the Occurrence entity.
func (tuo *TagUpdateOne) ClearOccurrences() *TagUpdateOne {
	tuo.mutation.ClearOccurrences()
	return tuo
}

// RemoveOccurrenceIDs removes the "occurrences" edge to Occurrence entities by IDs.
func (tuo *TagUpdateOne) RemoveOccurrenceIDs(ids ...uuid.UUID) *TagUpdateOne {
	tuo.mutation.RemoveOccurrenceIDs(ids...)
	return tuo
}

// RemoveOccurrences removes "occurrences" edges to Occurrence entities.
func (tuo *TagUpdateOne) RemoveOccurrences(o ...*Occurrence) *TagUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return tuo.RemoveOccurrenceIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagUpdateOne) Select(field string, fields ...string) *TagUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tag entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	var (
		err  error
		node *Tag
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Tag)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TagMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TagUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := tag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tag.name": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Description(); ok {
		if err := tag.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Tag.description": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.ShortName(); ok {
		if err := tag.ShortNameValidator(v); err != nil {
			return &ValidationError{Name: "short_name", err: fmt.Errorf(`ent: validator failed for field "Tag.short_name": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Priority(); ok {
		if err := tag.PriorityValidator(v); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Tag.priority": %w`, err)}
		}
	}
	return nil
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (_node *Tag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tag.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for _, f := range fields {
			if !tag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Key(); ok {
		_spec.SetField(tag.FieldKey, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(tag.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(tag.FieldDescription, field.TypeString, value)
	}
	if value, ok := tuo.mutation.ShortName(); ok {
		_spec.SetField(tag.FieldShortName, field.TypeString, value)
	}
	if tuo.mutation.ShortNameCleared() {
		_spec.ClearField(tag.FieldShortName, field.TypeString)
	}
	if value, ok := tuo.mutation.Priority(); ok {
		_spec.SetField(tag.FieldPriority, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.IsAllergy(); ok {
		_spec.SetField(tag.FieldIsAllergy, field.TypeBool, value)
	}
	if tuo.mutation.OccurrencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.OccurrencesTable,
			Columns: tag.OccurrencesPrimaryKey,
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
	if nodes := tuo.mutation.RemovedOccurrencesIDs(); len(nodes) > 0 && !tuo.mutation.OccurrencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.OccurrencesTable,
			Columns: tag.OccurrencesPrimaryKey,
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
	if nodes := tuo.mutation.OccurrencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.OccurrencesTable,
			Columns: tag.OccurrencesPrimaryKey,
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
	_node = &Tag{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
