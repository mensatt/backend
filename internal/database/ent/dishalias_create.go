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
)

// DishAliasCreate is the builder for creating a DishAlias entity.
type DishAliasCreate struct {
	config
	mutation *DishAliasMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNormalizedAliasName sets the "normalized_alias_name" field.
func (dac *DishAliasCreate) SetNormalizedAliasName(s string) *DishAliasCreate {
	dac.mutation.SetNormalizedAliasName(s)
	return dac
}

// SetID sets the "id" field.
func (dac *DishAliasCreate) SetID(s string) *DishAliasCreate {
	dac.mutation.SetID(s)
	return dac
}

// SetDishID sets the "dish" edge to the Dish entity by ID.
func (dac *DishAliasCreate) SetDishID(id uuid.UUID) *DishAliasCreate {
	dac.mutation.SetDishID(id)
	return dac
}

// SetDish sets the "dish" edge to the Dish entity.
func (dac *DishAliasCreate) SetDish(d *Dish) *DishAliasCreate {
	return dac.SetDishID(d.ID)
}

// Mutation returns the DishAliasMutation object of the builder.
func (dac *DishAliasCreate) Mutation() *DishAliasMutation {
	return dac.mutation
}

// Save creates the DishAlias in the database.
func (dac *DishAliasCreate) Save(ctx context.Context) (*DishAlias, error) {
	return withHooks(ctx, dac.sqlSave, dac.mutation, dac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dac *DishAliasCreate) SaveX(ctx context.Context) *DishAlias {
	v, err := dac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dac *DishAliasCreate) Exec(ctx context.Context) error {
	_, err := dac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dac *DishAliasCreate) ExecX(ctx context.Context) {
	if err := dac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dac *DishAliasCreate) check() error {
	if _, ok := dac.mutation.NormalizedAliasName(); !ok {
		return &ValidationError{Name: "normalized_alias_name", err: errors.New(`ent: missing required field "DishAlias.normalized_alias_name"`)}
	}
	if v, ok := dac.mutation.NormalizedAliasName(); ok {
		if err := dishalias.NormalizedAliasNameValidator(v); err != nil {
			return &ValidationError{Name: "normalized_alias_name", err: fmt.Errorf(`ent: validator failed for field "DishAlias.normalized_alias_name": %w`, err)}
		}
	}
	if _, ok := dac.mutation.DishID(); !ok {
		return &ValidationError{Name: "dish", err: errors.New(`ent: missing required edge "DishAlias.dish"`)}
	}
	return nil
}

func (dac *DishAliasCreate) sqlSave(ctx context.Context) (*DishAlias, error) {
	if err := dac.check(); err != nil {
		return nil, err
	}
	_node, _spec := dac.createSpec()
	if err := sqlgraph.CreateNode(ctx, dac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected DishAlias.ID type: %T", _spec.ID.Value)
		}
	}
	dac.mutation.id = &_node.ID
	dac.mutation.done = true
	return _node, nil
}

func (dac *DishAliasCreate) createSpec() (*DishAlias, *sqlgraph.CreateSpec) {
	var (
		_node = &DishAlias{config: dac.config}
		_spec = sqlgraph.NewCreateSpec(dishalias.Table, sqlgraph.NewFieldSpec(dishalias.FieldID, field.TypeString))
	)
	_spec.OnConflict = dac.conflict
	if id, ok := dac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dac.mutation.NormalizedAliasName(); ok {
		_spec.SetField(dishalias.FieldNormalizedAliasName, field.TypeString, value)
		_node.NormalizedAliasName = value
	}
	if nodes := dac.mutation.DishIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dishalias.DishTable,
			Columns: []string{dishalias.DishColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dish.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.dish = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DishAlias.Create().
//		SetNormalizedAliasName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DishAliasUpsert) {
//			SetNormalizedAliasName(v+v).
//		}).
//		Exec(ctx)
func (dac *DishAliasCreate) OnConflict(opts ...sql.ConflictOption) *DishAliasUpsertOne {
	dac.conflict = opts
	return &DishAliasUpsertOne{
		create: dac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DishAlias.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dac *DishAliasCreate) OnConflictColumns(columns ...string) *DishAliasUpsertOne {
	dac.conflict = append(dac.conflict, sql.ConflictColumns(columns...))
	return &DishAliasUpsertOne{
		create: dac,
	}
}

type (
	// DishAliasUpsertOne is the builder for "upsert"-ing
	//  one DishAlias node.
	DishAliasUpsertOne struct {
		create *DishAliasCreate
	}

	// DishAliasUpsert is the "OnConflict" setter.
	DishAliasUpsert struct {
		*sql.UpdateSet
	}
)

// SetNormalizedAliasName sets the "normalized_alias_name" field.
func (u *DishAliasUpsert) SetNormalizedAliasName(v string) *DishAliasUpsert {
	u.Set(dishalias.FieldNormalizedAliasName, v)
	return u
}

// UpdateNormalizedAliasName sets the "normalized_alias_name" field to the value that was provided on create.
func (u *DishAliasUpsert) UpdateNormalizedAliasName() *DishAliasUpsert {
	u.SetExcluded(dishalias.FieldNormalizedAliasName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DishAlias.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(dishalias.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DishAliasUpsertOne) UpdateNewValues() *DishAliasUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(dishalias.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DishAlias.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DishAliasUpsertOne) Ignore() *DishAliasUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DishAliasUpsertOne) DoNothing() *DishAliasUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DishAliasCreate.OnConflict
// documentation for more info.
func (u *DishAliasUpsertOne) Update(set func(*DishAliasUpsert)) *DishAliasUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DishAliasUpsert{UpdateSet: update})
	}))
	return u
}

// SetNormalizedAliasName sets the "normalized_alias_name" field.
func (u *DishAliasUpsertOne) SetNormalizedAliasName(v string) *DishAliasUpsertOne {
	return u.Update(func(s *DishAliasUpsert) {
		s.SetNormalizedAliasName(v)
	})
}

// UpdateNormalizedAliasName sets the "normalized_alias_name" field to the value that was provided on create.
func (u *DishAliasUpsertOne) UpdateNormalizedAliasName() *DishAliasUpsertOne {
	return u.Update(func(s *DishAliasUpsert) {
		s.UpdateNormalizedAliasName()
	})
}

// Exec executes the query.
func (u *DishAliasUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DishAliasCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DishAliasUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DishAliasUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DishAliasUpsertOne.ID is not supported by MySQL driver. Use DishAliasUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DishAliasUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DishAliasCreateBulk is the builder for creating many DishAlias entities in bulk.
type DishAliasCreateBulk struct {
	config
	err      error
	builders []*DishAliasCreate
	conflict []sql.ConflictOption
}

// Save creates the DishAlias entities in the database.
func (dacb *DishAliasCreateBulk) Save(ctx context.Context) ([]*DishAlias, error) {
	if dacb.err != nil {
		return nil, dacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dacb.builders))
	nodes := make([]*DishAlias, len(dacb.builders))
	mutators := make([]Mutator, len(dacb.builders))
	for i := range dacb.builders {
		func(i int, root context.Context) {
			builder := dacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DishAliasMutation)
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
					_, err = mutators[i+1].Mutate(root, dacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dacb *DishAliasCreateBulk) SaveX(ctx context.Context) []*DishAlias {
	v, err := dacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dacb *DishAliasCreateBulk) Exec(ctx context.Context) error {
	_, err := dacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dacb *DishAliasCreateBulk) ExecX(ctx context.Context) {
	if err := dacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DishAlias.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DishAliasUpsert) {
//			SetNormalizedAliasName(v+v).
//		}).
//		Exec(ctx)
func (dacb *DishAliasCreateBulk) OnConflict(opts ...sql.ConflictOption) *DishAliasUpsertBulk {
	dacb.conflict = opts
	return &DishAliasUpsertBulk{
		create: dacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DishAlias.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dacb *DishAliasCreateBulk) OnConflictColumns(columns ...string) *DishAliasUpsertBulk {
	dacb.conflict = append(dacb.conflict, sql.ConflictColumns(columns...))
	return &DishAliasUpsertBulk{
		create: dacb,
	}
}

// DishAliasUpsertBulk is the builder for "upsert"-ing
// a bulk of DishAlias nodes.
type DishAliasUpsertBulk struct {
	create *DishAliasCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DishAlias.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(dishalias.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DishAliasUpsertBulk) UpdateNewValues() *DishAliasUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(dishalias.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DishAlias.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DishAliasUpsertBulk) Ignore() *DishAliasUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DishAliasUpsertBulk) DoNothing() *DishAliasUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DishAliasCreateBulk.OnConflict
// documentation for more info.
func (u *DishAliasUpsertBulk) Update(set func(*DishAliasUpsert)) *DishAliasUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DishAliasUpsert{UpdateSet: update})
	}))
	return u
}

// SetNormalizedAliasName sets the "normalized_alias_name" field.
func (u *DishAliasUpsertBulk) SetNormalizedAliasName(v string) *DishAliasUpsertBulk {
	return u.Update(func(s *DishAliasUpsert) {
		s.SetNormalizedAliasName(v)
	})
}

// UpdateNormalizedAliasName sets the "normalized_alias_name" field to the value that was provided on create.
func (u *DishAliasUpsertBulk) UpdateNormalizedAliasName() *DishAliasUpsertBulk {
	return u.Update(func(s *DishAliasUpsert) {
		s.UpdateNormalizedAliasName()
	})
}

// Exec executes the query.
func (u *DishAliasUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DishAliasCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DishAliasCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DishAliasUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
