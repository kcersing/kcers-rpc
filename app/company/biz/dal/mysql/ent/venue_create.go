// Code generated by ent, DO NOT EDIT.

package ent

import (
	"company/biz/dal/mysql/ent/entrylogs"
	"company/biz/dal/mysql/ent/venue"
	"company/biz/dal/mysql/ent/venueplace"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VenueCreate is the builder for creating a Venue entity.
type VenueCreate struct {
	config
	mutation *VenueMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (vc *VenueCreate) SetCreatedAt(t time.Time) *VenueCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VenueCreate) SetNillableCreatedAt(t *time.Time) *VenueCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VenueCreate) SetUpdatedAt(t time.Time) *VenueCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VenueCreate) SetNillableUpdatedAt(t *time.Time) *VenueCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetStatus sets the "status" field.
func (vc *VenueCreate) SetStatus(i int64) *VenueCreate {
	vc.mutation.SetStatus(i)
	return vc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (vc *VenueCreate) SetNillableStatus(i *int64) *VenueCreate {
	if i != nil {
		vc.SetStatus(*i)
	}
	return vc
}

// SetName sets the "name" field.
func (vc *VenueCreate) SetName(s string) *VenueCreate {
	vc.mutation.SetName(s)
	return vc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (vc *VenueCreate) SetNillableName(s *string) *VenueCreate {
	if s != nil {
		vc.SetName(*s)
	}
	return vc
}

// SetAddress sets the "address" field.
func (vc *VenueCreate) SetAddress(s string) *VenueCreate {
	vc.mutation.SetAddress(s)
	return vc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (vc *VenueCreate) SetNillableAddress(s *string) *VenueCreate {
	if s != nil {
		vc.SetAddress(*s)
	}
	return vc
}

// SetAddressDetail sets the "address_detail" field.
func (vc *VenueCreate) SetAddressDetail(s string) *VenueCreate {
	vc.mutation.SetAddressDetail(s)
	return vc
}

// SetNillableAddressDetail sets the "address_detail" field if the given value is not nil.
func (vc *VenueCreate) SetNillableAddressDetail(s *string) *VenueCreate {
	if s != nil {
		vc.SetAddressDetail(*s)
	}
	return vc
}

// SetLatitude sets the "latitude" field.
func (vc *VenueCreate) SetLatitude(s string) *VenueCreate {
	vc.mutation.SetLatitude(s)
	return vc
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (vc *VenueCreate) SetNillableLatitude(s *string) *VenueCreate {
	if s != nil {
		vc.SetLatitude(*s)
	}
	return vc
}

// SetLongitude sets the "longitude" field.
func (vc *VenueCreate) SetLongitude(s string) *VenueCreate {
	vc.mutation.SetLongitude(s)
	return vc
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (vc *VenueCreate) SetNillableLongitude(s *string) *VenueCreate {
	if s != nil {
		vc.SetLongitude(*s)
	}
	return vc
}

// SetMobile sets the "mobile" field.
func (vc *VenueCreate) SetMobile(s string) *VenueCreate {
	vc.mutation.SetMobile(s)
	return vc
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (vc *VenueCreate) SetNillableMobile(s *string) *VenueCreate {
	if s != nil {
		vc.SetMobile(*s)
	}
	return vc
}

// SetPic sets the "pic" field.
func (vc *VenueCreate) SetPic(s string) *VenueCreate {
	vc.mutation.SetPic(s)
	return vc
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (vc *VenueCreate) SetNillablePic(s *string) *VenueCreate {
	if s != nil {
		vc.SetPic(*s)
	}
	return vc
}

// SetInformation sets the "information" field.
func (vc *VenueCreate) SetInformation(s string) *VenueCreate {
	vc.mutation.SetInformation(s)
	return vc
}

// SetNillableInformation sets the "information" field if the given value is not nil.
func (vc *VenueCreate) SetNillableInformation(s *string) *VenueCreate {
	if s != nil {
		vc.SetInformation(*s)
	}
	return vc
}

// SetID sets the "id" field.
func (vc *VenueCreate) SetID(i int64) *VenueCreate {
	vc.mutation.SetID(i)
	return vc
}

// AddPlaceIDs adds the "places" edge to the VenuePlace entity by IDs.
func (vc *VenueCreate) AddPlaceIDs(ids ...int64) *VenueCreate {
	vc.mutation.AddPlaceIDs(ids...)
	return vc
}

// AddPlaces adds the "places" edges to the VenuePlace entity.
func (vc *VenueCreate) AddPlaces(v ...*VenuePlace) *VenueCreate {
	ids := make([]int64, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vc.AddPlaceIDs(ids...)
}

// AddVenueEntryIDs adds the "venue_entry" edge to the EntryLogs entity by IDs.
func (vc *VenueCreate) AddVenueEntryIDs(ids ...int64) *VenueCreate {
	vc.mutation.AddVenueEntryIDs(ids...)
	return vc
}

// AddVenueEntry adds the "venue_entry" edges to the EntryLogs entity.
func (vc *VenueCreate) AddVenueEntry(e ...*EntryLogs) *VenueCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return vc.AddVenueEntryIDs(ids...)
}

// Mutation returns the VenueMutation object of the builder.
func (vc *VenueCreate) Mutation() *VenueMutation {
	return vc.mutation
}

// Save creates the Venue in the database.
func (vc *VenueCreate) Save(ctx context.Context) (*Venue, error) {
	vc.defaults()
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VenueCreate) SaveX(ctx context.Context) *Venue {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VenueCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VenueCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VenueCreate) defaults() {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := venue.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		v := venue.DefaultUpdatedAt()
		vc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vc.mutation.Status(); !ok {
		v := venue.DefaultStatus
		vc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VenueCreate) check() error {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Venue.created_at"`)}
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Venue.updated_at"`)}
	}
	return nil
}

func (vc *VenueCreate) sqlSave(ctx context.Context) (*Venue, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VenueCreate) createSpec() (*Venue, *sqlgraph.CreateSpec) {
	var (
		_node = &Venue{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(venue.Table, sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64))
	)
	_spec.Schema = vc.schemaConfig.Venue
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.SetField(venue.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.SetField(venue.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := vc.mutation.Status(); ok {
		_spec.SetField(venue.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := vc.mutation.Name(); ok {
		_spec.SetField(venue.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := vc.mutation.Address(); ok {
		_spec.SetField(venue.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := vc.mutation.AddressDetail(); ok {
		_spec.SetField(venue.FieldAddressDetail, field.TypeString, value)
		_node.AddressDetail = value
	}
	if value, ok := vc.mutation.Latitude(); ok {
		_spec.SetField(venue.FieldLatitude, field.TypeString, value)
		_node.Latitude = value
	}
	if value, ok := vc.mutation.Longitude(); ok {
		_spec.SetField(venue.FieldLongitude, field.TypeString, value)
		_node.Longitude = value
	}
	if value, ok := vc.mutation.Mobile(); ok {
		_spec.SetField(venue.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := vc.mutation.Pic(); ok {
		_spec.SetField(venue.FieldPic, field.TypeString, value)
		_node.Pic = value
	}
	if value, ok := vc.mutation.Information(); ok {
		_spec.SetField(venue.FieldInformation, field.TypeString, value)
		_node.Information = value
	}
	if nodes := vc.mutation.PlacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   venue.PlacesTable,
			Columns: []string{venue.PlacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(venueplace.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = vc.schemaConfig.VenuePlace
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.VenueEntryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   venue.VenueEntryTable,
			Columns: []string{venue.VenueEntryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entrylogs.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = vc.schemaConfig.EntryLogs
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VenueCreateBulk is the builder for creating many Venue entities in bulk.
type VenueCreateBulk struct {
	config
	err      error
	builders []*VenueCreate
}

// Save creates the Venue entities in the database.
func (vcb *VenueCreateBulk) Save(ctx context.Context) ([]*Venue, error) {
	if vcb.err != nil {
		return nil, vcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Venue, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VenueMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VenueCreateBulk) SaveX(ctx context.Context) []*Venue {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VenueCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VenueCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
