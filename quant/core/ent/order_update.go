// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"core/ent/order"
	"core/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetMaintOrderNo sets the "maintOrderNo" field.
func (ou *OrderUpdate) SetMaintOrderNo(s string) *OrderUpdate {
	ou.mutation.SetMaintOrderNo(s)
	return ou
}

// SetNillableMaintOrderNo sets the "maintOrderNo" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableMaintOrderNo(s *string) *OrderUpdate {
	if s != nil {
		ou.SetMaintOrderNo(*s)
	}
	return ou
}

// SetVehiclePlateNo sets the "vehiclePlateNo" field.
func (ou *OrderUpdate) SetVehiclePlateNo(s string) *OrderUpdate {
	ou.mutation.SetVehiclePlateNo(s)
	return ou
}

// SetNillableVehiclePlateNo sets the "vehiclePlateNo" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableVehiclePlateNo(s *string) *OrderUpdate {
	if s != nil {
		ou.SetVehiclePlateNo(*s)
	}
	return ou
}

// SetVehicleTeamName sets the "vehicleTeamName" field.
func (ou *OrderUpdate) SetVehicleTeamName(s string) *OrderUpdate {
	ou.mutation.SetVehicleTeamName(s)
	return ou
}

// SetNillableVehicleTeamName sets the "vehicleTeamName" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableVehicleTeamName(s *string) *OrderUpdate {
	if s != nil {
		ou.SetVehicleTeamName(*s)
	}
	return ou
}

// SetMaintRequestType sets the "maintRequestType" field.
func (ou *OrderUpdate) SetMaintRequestType(s string) *OrderUpdate {
	ou.mutation.SetMaintRequestType(s)
	return ou
}

// SetNillableMaintRequestType sets the "maintRequestType" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableMaintRequestType(s *string) *OrderUpdate {
	if s != nil {
		ou.SetMaintRequestType(*s)
	}
	return ou
}

// SetDispatchedStationId sets the "dispatchedStationId" field.
func (ou *OrderUpdate) SetDispatchedStationId(i int) *OrderUpdate {
	ou.mutation.ResetDispatchedStationId()
	ou.mutation.SetDispatchedStationId(i)
	return ou
}

// SetNillableDispatchedStationId sets the "dispatchedStationId" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableDispatchedStationId(i *int) *OrderUpdate {
	if i != nil {
		ou.SetDispatchedStationId(*i)
	}
	return ou
}

// AddDispatchedStationId adds i to the "dispatchedStationId" field.
func (ou *OrderUpdate) AddDispatchedStationId(i int) *OrderUpdate {
	ou.mutation.AddDispatchedStationId(i)
	return ou
}

// SetStationName sets the "stationName" field.
func (ou *OrderUpdate) SetStationName(s string) *OrderUpdate {
	ou.mutation.SetStationName(s)
	return ou
}

// SetNillableStationName sets the "stationName" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableStationName(s *string) *OrderUpdate {
	if s != nil {
		ou.SetStationName(*s)
	}
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OrderUpdate) SetUpdatedAt(t time.Time) *OrderUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	ou.defaults()
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrderUpdate) defaults() {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.MaintOrderNo(); ok {
		if err := order.MaintOrderNoValidator(v); err != nil {
			return &ValidationError{Name: "maintOrderNo", err: fmt.Errorf(`ent: validator failed for field "Order.maintOrderNo": %w`, err)}
		}
	}
	if v, ok := ou.mutation.VehiclePlateNo(); ok {
		if err := order.VehiclePlateNoValidator(v); err != nil {
			return &ValidationError{Name: "vehiclePlateNo", err: fmt.Errorf(`ent: validator failed for field "Order.vehiclePlateNo": %w`, err)}
		}
	}
	if v, ok := ou.mutation.VehicleTeamName(); ok {
		if err := order.VehicleTeamNameValidator(v); err != nil {
			return &ValidationError{Name: "vehicleTeamName", err: fmt.Errorf(`ent: validator failed for field "Order.vehicleTeamName": %w`, err)}
		}
	}
	if v, ok := ou.mutation.MaintRequestType(); ok {
		if err := order.MaintRequestTypeValidator(v); err != nil {
			return &ValidationError{Name: "maintRequestType", err: fmt.Errorf(`ent: validator failed for field "Order.maintRequestType": %w`, err)}
		}
	}
	if v, ok := ou.mutation.StationName(); ok {
		if err := order.StationNameValidator(v); err != nil {
			return &ValidationError{Name: "stationName", err: fmt.Errorf(`ent: validator failed for field "Order.stationName": %w`, err)}
		}
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.MaintOrderNo(); ok {
		_spec.SetField(order.FieldMaintOrderNo, field.TypeString, value)
	}
	if value, ok := ou.mutation.VehiclePlateNo(); ok {
		_spec.SetField(order.FieldVehiclePlateNo, field.TypeString, value)
	}
	if value, ok := ou.mutation.VehicleTeamName(); ok {
		_spec.SetField(order.FieldVehicleTeamName, field.TypeString, value)
	}
	if value, ok := ou.mutation.MaintRequestType(); ok {
		_spec.SetField(order.FieldMaintRequestType, field.TypeString, value)
	}
	if value, ok := ou.mutation.DispatchedStationId(); ok {
		_spec.SetField(order.FieldDispatchedStationId, field.TypeInt, value)
	}
	if value, ok := ou.mutation.AddedDispatchedStationId(); ok {
		_spec.AddField(order.FieldDispatchedStationId, field.TypeInt, value)
	}
	if value, ok := ou.mutation.StationName(); ok {
		_spec.SetField(order.FieldStationName, field.TypeString, value)
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetMaintOrderNo sets the "maintOrderNo" field.
func (ouo *OrderUpdateOne) SetMaintOrderNo(s string) *OrderUpdateOne {
	ouo.mutation.SetMaintOrderNo(s)
	return ouo
}

// SetNillableMaintOrderNo sets the "maintOrderNo" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableMaintOrderNo(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetMaintOrderNo(*s)
	}
	return ouo
}

// SetVehiclePlateNo sets the "vehiclePlateNo" field.
func (ouo *OrderUpdateOne) SetVehiclePlateNo(s string) *OrderUpdateOne {
	ouo.mutation.SetVehiclePlateNo(s)
	return ouo
}

// SetNillableVehiclePlateNo sets the "vehiclePlateNo" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableVehiclePlateNo(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetVehiclePlateNo(*s)
	}
	return ouo
}

// SetVehicleTeamName sets the "vehicleTeamName" field.
func (ouo *OrderUpdateOne) SetVehicleTeamName(s string) *OrderUpdateOne {
	ouo.mutation.SetVehicleTeamName(s)
	return ouo
}

// SetNillableVehicleTeamName sets the "vehicleTeamName" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableVehicleTeamName(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetVehicleTeamName(*s)
	}
	return ouo
}

// SetMaintRequestType sets the "maintRequestType" field.
func (ouo *OrderUpdateOne) SetMaintRequestType(s string) *OrderUpdateOne {
	ouo.mutation.SetMaintRequestType(s)
	return ouo
}

// SetNillableMaintRequestType sets the "maintRequestType" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableMaintRequestType(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetMaintRequestType(*s)
	}
	return ouo
}

// SetDispatchedStationId sets the "dispatchedStationId" field.
func (ouo *OrderUpdateOne) SetDispatchedStationId(i int) *OrderUpdateOne {
	ouo.mutation.ResetDispatchedStationId()
	ouo.mutation.SetDispatchedStationId(i)
	return ouo
}

// SetNillableDispatchedStationId sets the "dispatchedStationId" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDispatchedStationId(i *int) *OrderUpdateOne {
	if i != nil {
		ouo.SetDispatchedStationId(*i)
	}
	return ouo
}

// AddDispatchedStationId adds i to the "dispatchedStationId" field.
func (ouo *OrderUpdateOne) AddDispatchedStationId(i int) *OrderUpdateOne {
	ouo.mutation.AddDispatchedStationId(i)
	return ouo
}

// SetStationName sets the "stationName" field.
func (ouo *OrderUpdateOne) SetStationName(s string) *OrderUpdateOne {
	ouo.mutation.SetStationName(s)
	return ouo
}

// SetNillableStationName sets the "stationName" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableStationName(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetStationName(*s)
	}
	return ouo
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OrderUpdateOne) SetUpdatedAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ouo *OrderUpdateOne) Where(ps ...predicate.Order) *OrderUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	ouo.defaults()
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrderUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.MaintOrderNo(); ok {
		if err := order.MaintOrderNoValidator(v); err != nil {
			return &ValidationError{Name: "maintOrderNo", err: fmt.Errorf(`ent: validator failed for field "Order.maintOrderNo": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.VehiclePlateNo(); ok {
		if err := order.VehiclePlateNoValidator(v); err != nil {
			return &ValidationError{Name: "vehiclePlateNo", err: fmt.Errorf(`ent: validator failed for field "Order.vehiclePlateNo": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.VehicleTeamName(); ok {
		if err := order.VehicleTeamNameValidator(v); err != nil {
			return &ValidationError{Name: "vehicleTeamName", err: fmt.Errorf(`ent: validator failed for field "Order.vehicleTeamName": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.MaintRequestType(); ok {
		if err := order.MaintRequestTypeValidator(v); err != nil {
			return &ValidationError{Name: "maintRequestType", err: fmt.Errorf(`ent: validator failed for field "Order.maintRequestType": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.StationName(); ok {
		if err := order.StationNameValidator(v); err != nil {
			return &ValidationError{Name: "stationName", err: fmt.Errorf(`ent: validator failed for field "Order.stationName": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.MaintOrderNo(); ok {
		_spec.SetField(order.FieldMaintOrderNo, field.TypeString, value)
	}
	if value, ok := ouo.mutation.VehiclePlateNo(); ok {
		_spec.SetField(order.FieldVehiclePlateNo, field.TypeString, value)
	}
	if value, ok := ouo.mutation.VehicleTeamName(); ok {
		_spec.SetField(order.FieldVehicleTeamName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.MaintRequestType(); ok {
		_spec.SetField(order.FieldMaintRequestType, field.TypeString, value)
	}
	if value, ok := ouo.mutation.DispatchedStationId(); ok {
		_spec.SetField(order.FieldDispatchedStationId, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.AddedDispatchedStationId(); ok {
		_spec.AddField(order.FieldDispatchedStationId, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.StationName(); ok {
		_spec.SetField(order.FieldStationName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
