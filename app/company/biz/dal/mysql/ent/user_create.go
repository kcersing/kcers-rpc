// Code generated by ent, DO NOT EDIT.

package ent

import (
	"company/biz/dal/mysql/ent/entrylogs"
	"company/biz/dal/mysql/ent/face"
	"company/biz/dal/mysql/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(i int64) *UserCreate {
	uc.mutation.SetStatus(i)
	return uc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uc *UserCreate) SetNillableStatus(i *int64) *UserCreate {
	if i != nil {
		uc.SetStatus(*i)
	}
	return uc
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNickname sets the "nickname" field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.mutation.SetNickname(s)
	return uc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uc *UserCreate) SetNillableNickname(s *string) *UserCreate {
	if s != nil {
		uc.SetNickname(*s)
	}
	return uc
}

// SetSideMode sets the "side_mode" field.
func (uc *UserCreate) SetSideMode(s string) *UserCreate {
	uc.mutation.SetSideMode(s)
	return uc
}

// SetNillableSideMode sets the "side_mode" field if the given value is not nil.
func (uc *UserCreate) SetNillableSideMode(s *string) *UserCreate {
	if s != nil {
		uc.SetSideMode(*s)
	}
	return uc
}

// SetBaseColor sets the "base_color" field.
func (uc *UserCreate) SetBaseColor(s string) *UserCreate {
	uc.mutation.SetBaseColor(s)
	return uc
}

// SetNillableBaseColor sets the "base_color" field if the given value is not nil.
func (uc *UserCreate) SetNillableBaseColor(s *string) *UserCreate {
	if s != nil {
		uc.SetBaseColor(*s)
	}
	return uc
}

// SetActiveColor sets the "active_color" field.
func (uc *UserCreate) SetActiveColor(s string) *UserCreate {
	uc.mutation.SetActiveColor(s)
	return uc
}

// SetNillableActiveColor sets the "active_color" field if the given value is not nil.
func (uc *UserCreate) SetNillableActiveColor(s *string) *UserCreate {
	if s != nil {
		uc.SetActiveColor(*s)
	}
	return uc
}

// SetRoleID sets the "role_id" field.
func (uc *UserCreate) SetRoleID(i int64) *UserCreate {
	uc.mutation.SetRoleID(i)
	return uc
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableRoleID(i *int64) *UserCreate {
	if i != nil {
		uc.SetRoleID(*i)
	}
	return uc
}

// SetMobile sets the "mobile" field.
func (uc *UserCreate) SetMobile(s string) *UserCreate {
	uc.mutation.SetMobile(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetWecom sets the "wecom" field.
func (uc *UserCreate) SetWecom(s string) *UserCreate {
	uc.mutation.SetWecom(s)
	return uc
}

// SetNillableWecom sets the "wecom" field if the given value is not nil.
func (uc *UserCreate) SetNillableWecom(s *string) *UserCreate {
	if s != nil {
		uc.SetWecom(*s)
	}
	return uc
}

// SetJob sets the "job" field.
func (uc *UserCreate) SetJob(s string) *UserCreate {
	uc.mutation.SetJob(s)
	return uc
}

// SetNillableJob sets the "job" field if the given value is not nil.
func (uc *UserCreate) SetNillableJob(s *string) *UserCreate {
	if s != nil {
		uc.SetJob(*s)
	}
	return uc
}

// SetOrganization sets the "organization" field.
func (uc *UserCreate) SetOrganization(s string) *UserCreate {
	uc.mutation.SetOrganization(s)
	return uc
}

// SetNillableOrganization sets the "organization" field if the given value is not nil.
func (uc *UserCreate) SetNillableOrganization(s *string) *UserCreate {
	if s != nil {
		uc.SetOrganization(*s)
	}
	return uc
}

// SetDefaultVenueID sets the "default_venue_id" field.
func (uc *UserCreate) SetDefaultVenueID(i int64) *UserCreate {
	uc.mutation.SetDefaultVenueID(i)
	return uc
}

// SetNillableDefaultVenueID sets the "default_venue_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableDefaultVenueID(i *int64) *UserCreate {
	if i != nil {
		uc.SetDefaultVenueID(*i)
	}
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UserCreate) SetAvatar(s string) *UserCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatar(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatar(*s)
	}
	return uc
}

// SetGender sets the "gender" field.
func (uc *UserCreate) SetGender(i int64) *UserCreate {
	uc.mutation.SetGender(i)
	return uc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uc *UserCreate) SetNillableGender(i *int64) *UserCreate {
	if i != nil {
		uc.SetGender(*i)
	}
	return uc
}

// SetBirthday sets the "birthday" field.
func (uc *UserCreate) SetBirthday(t time.Time) *UserCreate {
	uc.mutation.SetBirthday(t)
	return uc
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (uc *UserCreate) SetNillableBirthday(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetBirthday(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int64) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// AddUserEntryIDs adds the "user_entry" edge to the EntryLogs entity by IDs.
func (uc *UserCreate) AddUserEntryIDs(ids ...int64) *UserCreate {
	uc.mutation.AddUserEntryIDs(ids...)
	return uc
}

// AddUserEntry adds the "user_entry" edges to the EntryLogs entity.
func (uc *UserCreate) AddUserEntry(e ...*EntryLogs) *UserCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uc.AddUserEntryIDs(ids...)
}

// AddUserFaceIDs adds the "user_face" edge to the Face entity by IDs.
func (uc *UserCreate) AddUserFaceIDs(ids ...int64) *UserCreate {
	uc.mutation.AddUserFaceIDs(ids...)
	return uc
}

// AddUserFace adds the "user_face" edges to the Face entity.
func (uc *UserCreate) AddUserFace(f ...*Face) *UserCreate {
	ids := make([]int64, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uc.AddUserFaceIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.Status(); !ok {
		v := user.DefaultStatus
		uc.mutation.SetStatus(v)
	}
	if _, ok := uc.mutation.SideMode(); !ok {
		v := user.DefaultSideMode
		uc.mutation.SetSideMode(v)
	}
	if _, ok := uc.mutation.BaseColor(); !ok {
		v := user.DefaultBaseColor
		uc.mutation.SetBaseColor(v)
	}
	if _, ok := uc.mutation.ActiveColor(); !ok {
		v := user.DefaultActiveColor
		uc.mutation.SetActiveColor(v)
	}
	if _, ok := uc.mutation.RoleID(); !ok {
		v := user.DefaultRoleID
		uc.mutation.SetRoleID(v)
	}
	if _, ok := uc.mutation.Gender(); !ok {
		v := user.DefaultGender
		uc.mutation.SetGender(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if _, ok := uc.mutation.Mobile(); !ok {
		return &ValidationError{Name: "mobile", err: errors.New(`ent: missing required field "User.mobile"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	)
	_spec.Schema = uc.schemaConfig.User
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := uc.mutation.SideMode(); ok {
		_spec.SetField(user.FieldSideMode, field.TypeString, value)
		_node.SideMode = value
	}
	if value, ok := uc.mutation.BaseColor(); ok {
		_spec.SetField(user.FieldBaseColor, field.TypeString, value)
		_node.BaseColor = value
	}
	if value, ok := uc.mutation.ActiveColor(); ok {
		_spec.SetField(user.FieldActiveColor, field.TypeString, value)
		_node.ActiveColor = value
	}
	if value, ok := uc.mutation.RoleID(); ok {
		_spec.SetField(user.FieldRoleID, field.TypeInt64, value)
		_node.RoleID = value
	}
	if value, ok := uc.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Wecom(); ok {
		_spec.SetField(user.FieldWecom, field.TypeString, value)
		_node.Wecom = value
	}
	if value, ok := uc.mutation.Job(); ok {
		_spec.SetField(user.FieldJob, field.TypeString, value)
		_node.Job = value
	}
	if value, ok := uc.mutation.Organization(); ok {
		_spec.SetField(user.FieldOrganization, field.TypeString, value)
		_node.Organization = value
	}
	if value, ok := uc.mutation.DefaultVenueID(); ok {
		_spec.SetField(user.FieldDefaultVenueID, field.TypeInt64, value)
		_node.DefaultVenueID = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := uc.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeInt64, value)
		_node.Gender = value
	}
	if value, ok := uc.mutation.Birthday(); ok {
		_spec.SetField(user.FieldBirthday, field.TypeTime, value)
		_node.Birthday = value
	}
	if nodes := uc.mutation.UserEntryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserEntryTable,
			Columns: []string{user.UserEntryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entrylogs.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = uc.schemaConfig.EntryLogs
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserFaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserFaceTable,
			Columns: []string{user.UserFaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(face.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = uc.schemaConfig.Face
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
