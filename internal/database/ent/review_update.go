// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent/image"
	"github.com/mensatt/backend/internal/database/ent/occurrence"
	"github.com/mensatt/backend/internal/database/ent/predicate"
	"github.com/mensatt/backend/internal/database/ent/review"
)

// ReviewUpdate is the builder for updating Review entities.
type ReviewUpdate struct {
	config
	hooks    []Hook
	mutation *ReviewMutation
}

// Where appends a list predicates to the ReviewUpdate builder.
func (ru *ReviewUpdate) Where(ps ...predicate.Review) *ReviewUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetDisplayName sets the "display_name" field.
func (ru *ReviewUpdate) SetDisplayName(s string) *ReviewUpdate {
	ru.mutation.SetDisplayName(s)
	return ru
}

// SetStars sets the "stars" field.
func (ru *ReviewUpdate) SetStars(i int) *ReviewUpdate {
	ru.mutation.ResetStars()
	ru.mutation.SetStars(i)
	return ru
}

// AddStars adds i to the "stars" field.
func (ru *ReviewUpdate) AddStars(i int) *ReviewUpdate {
	ru.mutation.AddStars(i)
	return ru
}

// SetText sets the "text" field.
func (ru *ReviewUpdate) SetText(s string) *ReviewUpdate {
	ru.mutation.SetText(s)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *ReviewUpdate) SetUpdatedAt(t time.Time) *ReviewUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetAcceptedAt sets the "accepted_at" field.
func (ru *ReviewUpdate) SetAcceptedAt(t time.Time) *ReviewUpdate {
	ru.mutation.SetAcceptedAt(t)
	return ru
}

// SetNillableAcceptedAt sets the "accepted_at" field if the given value is not nil.
func (ru *ReviewUpdate) SetNillableAcceptedAt(t *time.Time) *ReviewUpdate {
	if t != nil {
		ru.SetAcceptedAt(*t)
	}
	return ru
}

// ClearAcceptedAt clears the value of the "accepted_at" field.
func (ru *ReviewUpdate) ClearAcceptedAt() *ReviewUpdate {
	ru.mutation.ClearAcceptedAt()
	return ru
}

// SetOccurrenceID sets the "occurrence" edge to the Occurrence entity by ID.
func (ru *ReviewUpdate) SetOccurrenceID(id uuid.UUID) *ReviewUpdate {
	ru.mutation.SetOccurrenceID(id)
	return ru
}

// SetOccurrence sets the "occurrence" edge to the Occurrence entity.
func (ru *ReviewUpdate) SetOccurrence(o *Occurrence) *ReviewUpdate {
	return ru.SetOccurrenceID(o.ID)
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (ru *ReviewUpdate) AddImageIDs(ids ...uuid.UUID) *ReviewUpdate {
	ru.mutation.AddImageIDs(ids...)
	return ru
}

// AddImages adds the "images" edges to the Image entity.
func (ru *ReviewUpdate) AddImages(i ...*Image) *ReviewUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ru.AddImageIDs(ids...)
}

// Mutation returns the ReviewMutation object of the builder.
func (ru *ReviewUpdate) Mutation() *ReviewMutation {
	return ru.mutation
}

// ClearOccurrence clears the "occurrence" edge to the Occurrence entity.
func (ru *ReviewUpdate) ClearOccurrence() *ReviewUpdate {
	ru.mutation.ClearOccurrence()
	return ru
}

// ClearImages clears all "images" edges to the Image entity.
func (ru *ReviewUpdate) ClearImages() *ReviewUpdate {
	ru.mutation.ClearImages()
	return ru
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (ru *ReviewUpdate) RemoveImageIDs(ids ...uuid.UUID) *ReviewUpdate {
	ru.mutation.RemoveImageIDs(ids...)
	return ru
}

// RemoveImages removes "images" edges to Image entities.
func (ru *ReviewUpdate) RemoveImages(i ...*Image) *ReviewUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ru.RemoveImageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReviewUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReviewUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReviewUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReviewUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ReviewUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := review.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReviewUpdate) check() error {
	if v, ok := ru.mutation.DisplayName(); ok {
		if err := review.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`ent: validator failed for field "Review.display_name": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Stars(); ok {
		if err := review.StarsValidator(v); err != nil {
			return &ValidationError{Name: "stars", err: fmt.Errorf(`ent: validator failed for field "Review.stars": %w`, err)}
		}
	}
	if _, ok := ru.mutation.OccurrenceID(); ru.mutation.OccurrenceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Review.occurrence"`)
	}
	return nil
}

func (ru *ReviewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: review.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.DisplayName(); ok {
		_spec.SetField(review.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Stars(); ok {
		_spec.SetField(review.FieldStars, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedStars(); ok {
		_spec.AddField(review.FieldStars, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Text(); ok {
		_spec.SetField(review.FieldText, field.TypeString, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(review.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.AcceptedAt(); ok {
		_spec.SetField(review.FieldAcceptedAt, field.TypeTime, value)
	}
	if ru.mutation.AcceptedAtCleared() {
		_spec.ClearField(review.FieldAcceptedAt, field.TypeTime)
	}
	if ru.mutation.OccurrenceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.OccurrenceTable,
			Columns: []string{review.OccurrenceColumn},
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
	if nodes := ru.mutation.OccurrenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.OccurrenceTable,
			Columns: []string{review.OccurrenceColumn},
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
	if ru.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.ImagesTable,
			Columns: []string{review.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedImagesIDs(); len(nodes) > 0 && !ru.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.ImagesTable,
			Columns: []string{review.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.ImagesTable,
			Columns: []string{review.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ReviewUpdateOne is the builder for updating a single Review entity.
type ReviewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReviewMutation
}

// SetDisplayName sets the "display_name" field.
func (ruo *ReviewUpdateOne) SetDisplayName(s string) *ReviewUpdateOne {
	ruo.mutation.SetDisplayName(s)
	return ruo
}

// SetStars sets the "stars" field.
func (ruo *ReviewUpdateOne) SetStars(i int) *ReviewUpdateOne {
	ruo.mutation.ResetStars()
	ruo.mutation.SetStars(i)
	return ruo
}

// AddStars adds i to the "stars" field.
func (ruo *ReviewUpdateOne) AddStars(i int) *ReviewUpdateOne {
	ruo.mutation.AddStars(i)
	return ruo
}

// SetText sets the "text" field.
func (ruo *ReviewUpdateOne) SetText(s string) *ReviewUpdateOne {
	ruo.mutation.SetText(s)
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *ReviewUpdateOne) SetUpdatedAt(t time.Time) *ReviewUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetAcceptedAt sets the "accepted_at" field.
func (ruo *ReviewUpdateOne) SetAcceptedAt(t time.Time) *ReviewUpdateOne {
	ruo.mutation.SetAcceptedAt(t)
	return ruo
}

// SetNillableAcceptedAt sets the "accepted_at" field if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableAcceptedAt(t *time.Time) *ReviewUpdateOne {
	if t != nil {
		ruo.SetAcceptedAt(*t)
	}
	return ruo
}

// ClearAcceptedAt clears the value of the "accepted_at" field.
func (ruo *ReviewUpdateOne) ClearAcceptedAt() *ReviewUpdateOne {
	ruo.mutation.ClearAcceptedAt()
	return ruo
}

// SetOccurrenceID sets the "occurrence" edge to the Occurrence entity by ID.
func (ruo *ReviewUpdateOne) SetOccurrenceID(id uuid.UUID) *ReviewUpdateOne {
	ruo.mutation.SetOccurrenceID(id)
	return ruo
}

// SetOccurrence sets the "occurrence" edge to the Occurrence entity.
func (ruo *ReviewUpdateOne) SetOccurrence(o *Occurrence) *ReviewUpdateOne {
	return ruo.SetOccurrenceID(o.ID)
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (ruo *ReviewUpdateOne) AddImageIDs(ids ...uuid.UUID) *ReviewUpdateOne {
	ruo.mutation.AddImageIDs(ids...)
	return ruo
}

// AddImages adds the "images" edges to the Image entity.
func (ruo *ReviewUpdateOne) AddImages(i ...*Image) *ReviewUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ruo.AddImageIDs(ids...)
}

// Mutation returns the ReviewMutation object of the builder.
func (ruo *ReviewUpdateOne) Mutation() *ReviewMutation {
	return ruo.mutation
}

// ClearOccurrence clears the "occurrence" edge to the Occurrence entity.
func (ruo *ReviewUpdateOne) ClearOccurrence() *ReviewUpdateOne {
	ruo.mutation.ClearOccurrence()
	return ruo
}

// ClearImages clears all "images" edges to the Image entity.
func (ruo *ReviewUpdateOne) ClearImages() *ReviewUpdateOne {
	ruo.mutation.ClearImages()
	return ruo
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (ruo *ReviewUpdateOne) RemoveImageIDs(ids ...uuid.UUID) *ReviewUpdateOne {
	ruo.mutation.RemoveImageIDs(ids...)
	return ruo
}

// RemoveImages removes "images" edges to Image entities.
func (ruo *ReviewUpdateOne) RemoveImages(i ...*Image) *ReviewUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ruo.RemoveImageIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReviewUpdateOne) Select(field string, fields ...string) *ReviewUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Review entity.
func (ruo *ReviewUpdateOne) Save(ctx context.Context) (*Review, error) {
	var (
		err  error
		node *Review
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Review)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ReviewMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReviewUpdateOne) SaveX(ctx context.Context) *Review {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReviewUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReviewUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ReviewUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := review.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReviewUpdateOne) check() error {
	if v, ok := ruo.mutation.DisplayName(); ok {
		if err := review.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`ent: validator failed for field "Review.display_name": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Stars(); ok {
		if err := review.StarsValidator(v); err != nil {
			return &ValidationError{Name: "stars", err: fmt.Errorf(`ent: validator failed for field "Review.stars": %w`, err)}
		}
	}
	if _, ok := ruo.mutation.OccurrenceID(); ruo.mutation.OccurrenceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Review.occurrence"`)
	}
	return nil
}

func (ruo *ReviewUpdateOne) sqlSave(ctx context.Context) (_node *Review, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: review.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Review.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, review.FieldID)
		for _, f := range fields {
			if !review.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != review.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.DisplayName(); ok {
		_spec.SetField(review.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Stars(); ok {
		_spec.SetField(review.FieldStars, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedStars(); ok {
		_spec.AddField(review.FieldStars, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Text(); ok {
		_spec.SetField(review.FieldText, field.TypeString, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(review.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.AcceptedAt(); ok {
		_spec.SetField(review.FieldAcceptedAt, field.TypeTime, value)
	}
	if ruo.mutation.AcceptedAtCleared() {
		_spec.ClearField(review.FieldAcceptedAt, field.TypeTime)
	}
	if ruo.mutation.OccurrenceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.OccurrenceTable,
			Columns: []string{review.OccurrenceColumn},
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
	if nodes := ruo.mutation.OccurrenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.OccurrenceTable,
			Columns: []string{review.OccurrenceColumn},
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
	if ruo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.ImagesTable,
			Columns: []string{review.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedImagesIDs(); len(nodes) > 0 && !ruo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.ImagesTable,
			Columns: []string{review.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.ImagesTable,
			Columns: []string{review.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Review{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
