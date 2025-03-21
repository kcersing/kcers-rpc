// Code generated by ent, DO NOT EDIT.

package ent

import (
	"company/biz/dal/mysql/ent/contract"
	"company/biz/dal/mysql/ent/internal"
	"company/biz/dal/mysql/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContractUpdate is the builder for updating Contract entities.
type ContractUpdate struct {
	config
	hooks    []Hook
	mutation *ContractMutation
}

// Where appends a list predicates to the ContractUpdate builder.
func (cu *ContractUpdate) Where(ps ...predicate.Contract) *ContractUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ContractUpdate) SetUpdatedAt(t time.Time) *ContractUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetStatus sets the "status" field.
func (cu *ContractUpdate) SetStatus(i int64) *ContractUpdate {
	cu.mutation.ResetStatus()
	cu.mutation.SetStatus(i)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableStatus(i *int64) *ContractUpdate {
	if i != nil {
		cu.SetStatus(*i)
	}
	return cu
}

// AddStatus adds i to the "status" field.
func (cu *ContractUpdate) AddStatus(i int64) *ContractUpdate {
	cu.mutation.AddStatus(i)
	return cu
}

// ClearStatus clears the value of the "status" field.
func (cu *ContractUpdate) ClearStatus() *ContractUpdate {
	cu.mutation.ClearStatus()
	return cu
}

// SetName sets the "name" field.
func (cu *ContractUpdate) SetName(s string) *ContractUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableName(s *string) *ContractUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of the "name" field.
func (cu *ContractUpdate) ClearName() *ContractUpdate {
	cu.mutation.ClearName()
	return cu
}

// SetContent sets the "content" field.
func (cu *ContractUpdate) SetContent(s string) *ContractUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableContent(s *string) *ContractUpdate {
	if s != nil {
		cu.SetContent(*s)
	}
	return cu
}

// ClearContent clears the value of the "content" field.
func (cu *ContractUpdate) ClearContent() *ContractUpdate {
	cu.mutation.ClearContent()
	return cu
}

// Mutation returns the ContractMutation object of the builder.
func (cu *ContractUpdate) Mutation() *ContractMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ContractUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ContractUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ContractUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ContractUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ContractUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := contract.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *ContractUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(contract.Table, contract.Columns, sqlgraph.NewFieldSpec(contract.FieldID, field.TypeInt64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(contract.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(contract.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedStatus(); ok {
		_spec.AddField(contract.FieldStatus, field.TypeInt64, value)
	}
	if cu.mutation.StatusCleared() {
		_spec.ClearField(contract.FieldStatus, field.TypeInt64)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(contract.FieldName, field.TypeString, value)
	}
	if cu.mutation.NameCleared() {
		_spec.ClearField(contract.FieldName, field.TypeString)
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.SetField(contract.FieldContent, field.TypeString, value)
	}
	if cu.mutation.ContentCleared() {
		_spec.ClearField(contract.FieldContent, field.TypeString)
	}
	_spec.Node.Schema = cu.schemaConfig.Contract
	ctx = internal.NewSchemaConfigContext(ctx, cu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contract.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ContractUpdateOne is the builder for updating a single Contract entity.
type ContractUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContractMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ContractUpdateOne) SetUpdatedAt(t time.Time) *ContractUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *ContractUpdateOne) SetStatus(i int64) *ContractUpdateOne {
	cuo.mutation.ResetStatus()
	cuo.mutation.SetStatus(i)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableStatus(i *int64) *ContractUpdateOne {
	if i != nil {
		cuo.SetStatus(*i)
	}
	return cuo
}

// AddStatus adds i to the "status" field.
func (cuo *ContractUpdateOne) AddStatus(i int64) *ContractUpdateOne {
	cuo.mutation.AddStatus(i)
	return cuo
}

// ClearStatus clears the value of the "status" field.
func (cuo *ContractUpdateOne) ClearStatus() *ContractUpdateOne {
	cuo.mutation.ClearStatus()
	return cuo
}

// SetName sets the "name" field.
func (cuo *ContractUpdateOne) SetName(s string) *ContractUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableName(s *string) *ContractUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of the "name" field.
func (cuo *ContractUpdateOne) ClearName() *ContractUpdateOne {
	cuo.mutation.ClearName()
	return cuo
}

// SetContent sets the "content" field.
func (cuo *ContractUpdateOne) SetContent(s string) *ContractUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableContent(s *string) *ContractUpdateOne {
	if s != nil {
		cuo.SetContent(*s)
	}
	return cuo
}

// ClearContent clears the value of the "content" field.
func (cuo *ContractUpdateOne) ClearContent() *ContractUpdateOne {
	cuo.mutation.ClearContent()
	return cuo
}

// Mutation returns the ContractMutation object of the builder.
func (cuo *ContractUpdateOne) Mutation() *ContractMutation {
	return cuo.mutation
}

// Where appends a list predicates to the ContractUpdate builder.
func (cuo *ContractUpdateOne) Where(ps ...predicate.Contract) *ContractUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ContractUpdateOne) Select(field string, fields ...string) *ContractUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Contract entity.
func (cuo *ContractUpdateOne) Save(ctx context.Context) (*Contract, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ContractUpdateOne) SaveX(ctx context.Context) *Contract {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ContractUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ContractUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ContractUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := contract.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *ContractUpdateOne) sqlSave(ctx context.Context) (_node *Contract, err error) {
	_spec := sqlgraph.NewUpdateSpec(contract.Table, contract.Columns, sqlgraph.NewFieldSpec(contract.FieldID, field.TypeInt64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Contract.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contract.FieldID)
		for _, f := range fields {
			if !contract.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contract.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(contract.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(contract.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedStatus(); ok {
		_spec.AddField(contract.FieldStatus, field.TypeInt64, value)
	}
	if cuo.mutation.StatusCleared() {
		_spec.ClearField(contract.FieldStatus, field.TypeInt64)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(contract.FieldName, field.TypeString, value)
	}
	if cuo.mutation.NameCleared() {
		_spec.ClearField(contract.FieldName, field.TypeString)
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.SetField(contract.FieldContent, field.TypeString, value)
	}
	if cuo.mutation.ContentCleared() {
		_spec.ClearField(contract.FieldContent, field.TypeString)
	}
	_spec.Node.Schema = cuo.schemaConfig.Contract
	ctx = internal.NewSchemaConfigContext(ctx, cuo.schemaConfig)
	_node = &Contract{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contract.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
