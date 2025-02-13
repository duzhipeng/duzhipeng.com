// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bot/ent/station"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StationCreate is the builder for creating a Station entity.
type StationCreate struct {
	config
	mutation *StationMutation
	hooks    []Hook
}

// SetAgencyId sets the "agencyId" field.
func (sc *StationCreate) SetAgencyId(i int) *StationCreate {
	sc.mutation.SetAgencyId(i)
	return sc
}

// SetAgencyName sets the "agencyName" field.
func (sc *StationCreate) SetAgencyName(s string) *StationCreate {
	sc.mutation.SetAgencyName(s)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *StationCreate) SetCreatedAt(t time.Time) *StationCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StationCreate) SetNillableCreatedAt(t *time.Time) *StationCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StationCreate) SetUpdatedAt(t time.Time) *StationCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StationCreate) SetNillableUpdatedAt(t *time.Time) *StationCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// Mutation returns the StationMutation object of the builder.
func (sc *StationCreate) Mutation() *StationMutation {
	return sc.mutation
}

// Save creates the Station in the database.
func (sc *StationCreate) Save(ctx context.Context) (*Station, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StationCreate) SaveX(ctx context.Context) *Station {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StationCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StationCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StationCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := station.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := station.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StationCreate) check() error {
	if _, ok := sc.mutation.AgencyId(); !ok {
		return &ValidationError{Name: "agencyId", err: errors.New(`ent: missing required field "Station.agencyId"`)}
	}
	if _, ok := sc.mutation.AgencyName(); !ok {
		return &ValidationError{Name: "agencyName", err: errors.New(`ent: missing required field "Station.agencyName"`)}
	}
	if v, ok := sc.mutation.AgencyName(); ok {
		if err := station.AgencyNameValidator(v); err != nil {
			return &ValidationError{Name: "agencyName", err: fmt.Errorf(`ent: validator failed for field "Station.agencyName": %w`, err)}
		}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Station.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Station.updated_at"`)}
	}
	return nil
}

func (sc *StationCreate) sqlSave(ctx context.Context) (*Station, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StationCreate) createSpec() (*Station, *sqlgraph.CreateSpec) {
	var (
		_node = &Station{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(station.Table, sqlgraph.NewFieldSpec(station.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.AgencyId(); ok {
		_spec.SetField(station.FieldAgencyId, field.TypeInt, value)
		_node.AgencyId = value
	}
	if value, ok := sc.mutation.AgencyName(); ok {
		_spec.SetField(station.FieldAgencyName, field.TypeString, value)
		_node.AgencyName = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(station.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(station.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// StationCreateBulk is the builder for creating many Station entities in bulk.
type StationCreateBulk struct {
	config
	err      error
	builders []*StationCreate
}

// Save creates the Station entities in the database.
func (scb *StationCreateBulk) Save(ctx context.Context) ([]*Station, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Station, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StationMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StationCreateBulk) SaveX(ctx context.Context) []*Station {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StationCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StationCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
