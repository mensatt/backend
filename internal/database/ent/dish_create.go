// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent/dish"
	"github.com/mensatt/backend/internal/database/ent/dishalias"
	"github.com/mensatt/backend/internal/database/ent/occurrence"
)

// DishCreate is the builder for creating a Dish entity.
type DishCreate struct {
	config
	mutation *DishMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNameDe sets the "name_de" field.
func (dc *DishCreate) SetNameDe(s string) *DishCreate {
	dc.mutation.SetNameDe(s)
	return dc
}

// SetNameEn sets the "name_en" field.
func (dc *DishCreate) SetNameEn(s string) *DishCreate {
	dc.mutation.SetNameEn(s)
	return dc
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (dc *DishCreate) SetNillableNameEn(s *string) *DishCreate {
	if s != nil {
		dc.SetNameEn(*s)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DishCreate) SetID(u uuid.UUID) *DishCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DishCreate) SetNillableID(u *uuid.UUID) *DishCreate {
	if u != nil {
		dc.SetID(*u)
	}
	return dc
}

// AddDishOccurrenceIDs adds the "dish_occurrences" edge to the Occurrence entity by IDs.
func (dc *DishCreate) AddDishOccurrenceIDs(ids ...uuid.UUID) *DishCreate {
	dc.mutation.AddDishOccurrenceIDs(ids...)
	return dc
}

// AddDishOccurrences adds the "dish_occurrences" edges to the Occurrence entity.
func (dc *DishCreate) AddDishOccurrences(o ...*Occurrence) *DishCreate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return dc.AddDishOccurrenceIDs(ids...)
}

// AddAliasIDs adds the "aliases" edge to the DishAlias entity by IDs.
func (dc *DishCreate) AddAliasIDs(ids ...string) *DishCreate {
	dc.mutation.AddAliasIDs(ids...)
	return dc
}

// AddAliases adds the "aliases" edges to the DishAlias entity.
func (dc *DishCreate) AddAliases(d ...*DishAlias) *DishCreate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddAliasIDs(ids...)
}

// AddSideDishOccurrenceIDs adds the "side_dish_occurrence" edge to the Occurrence entity by IDs.
func (dc *DishCreate) AddSideDishOccurrenceIDs(ids ...uuid.UUID) *DishCreate {
	dc.mutation.AddSideDishOccurrenceIDs(ids...)
	return dc
}

// AddSideDishOccurrence adds the "side_dish_occurrence" edges to the Occurrence entity.
func (dc *DishCreate) AddSideDishOccurrence(o ...*Occurrence) *DishCreate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return dc.AddSideDishOccurrenceIDs(ids...)
}

// Mutation returns the DishMutation object of the builder.
func (dc *DishCreate) Mutation() *DishMutation {
	return dc.mutation
}

// Save creates the Dish in the database.
func (dc *DishCreate) Save(ctx context.Context) (*Dish, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DishCreate) SaveX(ctx context.Context) *Dish {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DishCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DishCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DishCreate) defaults() {
	if _, ok := dc.mutation.ID(); !ok {
		v := dish.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DishCreate) check() error {
	if _, ok := dc.mutation.NameDe(); !ok {
		return &ValidationError{Name: "name_de", err: errors.New(`ent: missing required field "Dish.name_de"`)}
	}
	if v, ok := dc.mutation.NameDe(); ok {
		if err := dish.NameDeValidator(v); err != nil {
			return &ValidationError{Name: "name_de", err: fmt.Errorf(`ent: validator failed for field "Dish.name_de": %w`, err)}
		}
	}
	if v, ok := dc.mutation.NameEn(); ok {
		if err := dish.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf(`ent: validator failed for field "Dish.name_en": %w`, err)}
		}
	}
	return nil
}

func (dc *DishCreate) sqlSave(ctx context.Context) (*Dish, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DishCreate) createSpec() (*Dish, *sqlgraph.CreateSpec) {
	var (
		_node = &Dish{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(dish.Table, sqlgraph.NewFieldSpec(dish.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = dc.conflict
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.NameDe(); ok {
		_spec.SetField(dish.FieldNameDe, field.TypeString, value)
		_node.NameDe = value
	}
	if value, ok := dc.mutation.NameEn(); ok {
		_spec.SetField(dish.FieldNameEn, field.TypeString, value)
		_node.NameEn = &value
	}
	if nodes := dc.mutation.DishOccurrencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.DishOccurrencesTable,
			Columns: []string{dish.DishOccurrencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(occurrence.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.AliasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dish.AliasesTable,
			Columns: []string{dish.AliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dishalias.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.SideDishOccurrenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dish.SideDishOccurrenceTable,
			Columns: dish.SideDishOccurrencePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(occurrence.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Dish.Create().
//		SetNameDe(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DishUpsert) {
//			SetNameDe(v+v).
//		}).
//		Exec(ctx)
func (dc *DishCreate) OnConflict(opts ...sql.ConflictOption) *DishUpsertOne {
	dc.conflict = opts
	return &DishUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Dish.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dc *DishCreate) OnConflictColumns(columns ...string) *DishUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DishUpsertOne{
		create: dc,
	}
}

type (
	// DishUpsertOne is the builder for "upsert"-ing
	//  one Dish node.
	DishUpsertOne struct {
		create *DishCreate
	}

	// DishUpsert is the "OnConflict" setter.
	DishUpsert struct {
		*sql.UpdateSet
	}
)

// SetNameDe sets the "name_de" field.
func (u *DishUpsert) SetNameDe(v string) *DishUpsert {
	u.Set(dish.FieldNameDe, v)
	return u
}

// UpdateNameDe sets the "name_de" field to the value that was provided on create.
func (u *DishUpsert) UpdateNameDe() *DishUpsert {
	u.SetExcluded(dish.FieldNameDe)
	return u
}

// SetNameEn sets the "name_en" field.
func (u *DishUpsert) SetNameEn(v string) *DishUpsert {
	u.Set(dish.FieldNameEn, v)
	return u
}

// UpdateNameEn sets the "name_en" field to the value that was provided on create.
func (u *DishUpsert) UpdateNameEn() *DishUpsert {
	u.SetExcluded(dish.FieldNameEn)
	return u
}

// ClearNameEn clears the value of the "name_en" field.
func (u *DishUpsert) ClearNameEn() *DishUpsert {
	u.SetNull(dish.FieldNameEn)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Dish.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(dish.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DishUpsertOne) UpdateNewValues() *DishUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(dish.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Dish.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DishUpsertOne) Ignore() *DishUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DishUpsertOne) DoNothing() *DishUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DishCreate.OnConflict
// documentation for more info.
func (u *DishUpsertOne) Update(set func(*DishUpsert)) *DishUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DishUpsert{UpdateSet: update})
	}))
	return u
}

// SetNameDe sets the "name_de" field.
func (u *DishUpsertOne) SetNameDe(v string) *DishUpsertOne {
	return u.Update(func(s *DishUpsert) {
		s.SetNameDe(v)
	})
}

// UpdateNameDe sets the "name_de" field to the value that was provided on create.
func (u *DishUpsertOne) UpdateNameDe() *DishUpsertOne {
	return u.Update(func(s *DishUpsert) {
		s.UpdateNameDe()
	})
}

// SetNameEn sets the "name_en" field.
func (u *DishUpsertOne) SetNameEn(v string) *DishUpsertOne {
	return u.Update(func(s *DishUpsert) {
		s.SetNameEn(v)
	})
}

// UpdateNameEn sets the "name_en" field to the value that was provided on create.
func (u *DishUpsertOne) UpdateNameEn() *DishUpsertOne {
	return u.Update(func(s *DishUpsert) {
		s.UpdateNameEn()
	})
}

// ClearNameEn clears the value of the "name_en" field.
func (u *DishUpsertOne) ClearNameEn() *DishUpsertOne {
	return u.Update(func(s *DishUpsert) {
		s.ClearNameEn()
	})
}

// Exec executes the query.
func (u *DishUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DishCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DishUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DishUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DishUpsertOne.ID is not supported by MySQL driver. Use DishUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DishUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DishCreateBulk is the builder for creating many Dish entities in bulk.
type DishCreateBulk struct {
	config
	builders []*DishCreate
	conflict []sql.ConflictOption
}

// Save creates the Dish entities in the database.
func (dcb *DishCreateBulk) Save(ctx context.Context) ([]*Dish, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dish, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DishMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DishCreateBulk) SaveX(ctx context.Context) []*Dish {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DishCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DishCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Dish.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DishUpsert) {
//			SetNameDe(v+v).
//		}).
//		Exec(ctx)
func (dcb *DishCreateBulk) OnConflict(opts ...sql.ConflictOption) *DishUpsertBulk {
	dcb.conflict = opts
	return &DishUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Dish.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dcb *DishCreateBulk) OnConflictColumns(columns ...string) *DishUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DishUpsertBulk{
		create: dcb,
	}
}

// DishUpsertBulk is the builder for "upsert"-ing
// a bulk of Dish nodes.
type DishUpsertBulk struct {
	create *DishCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Dish.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(dish.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DishUpsertBulk) UpdateNewValues() *DishUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(dish.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Dish.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DishUpsertBulk) Ignore() *DishUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DishUpsertBulk) DoNothing() *DishUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DishCreateBulk.OnConflict
// documentation for more info.
func (u *DishUpsertBulk) Update(set func(*DishUpsert)) *DishUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DishUpsert{UpdateSet: update})
	}))
	return u
}

// SetNameDe sets the "name_de" field.
func (u *DishUpsertBulk) SetNameDe(v string) *DishUpsertBulk {
	return u.Update(func(s *DishUpsert) {
		s.SetNameDe(v)
	})
}

// UpdateNameDe sets the "name_de" field to the value that was provided on create.
func (u *DishUpsertBulk) UpdateNameDe() *DishUpsertBulk {
	return u.Update(func(s *DishUpsert) {
		s.UpdateNameDe()
	})
}

// SetNameEn sets the "name_en" field.
func (u *DishUpsertBulk) SetNameEn(v string) *DishUpsertBulk {
	return u.Update(func(s *DishUpsert) {
		s.SetNameEn(v)
	})
}

// UpdateNameEn sets the "name_en" field to the value that was provided on create.
func (u *DishUpsertBulk) UpdateNameEn() *DishUpsertBulk {
	return u.Update(func(s *DishUpsert) {
		s.UpdateNameEn()
	})
}

// ClearNameEn clears the value of the "name_en" field.
func (u *DishUpsertBulk) ClearNameEn() *DishUpsertBulk {
	return u.Update(func(s *DishUpsert) {
		s.ClearNameEn()
	})
}

// Exec executes the query.
func (u *DishUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DishCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DishCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DishUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
