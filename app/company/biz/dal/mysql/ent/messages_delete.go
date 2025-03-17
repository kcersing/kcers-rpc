// Code generated by ent, DO NOT EDIT.

package ent

import (
	"company/biz/dal/mysql/ent/internal"
	"company/biz/dal/mysql/ent/messages"
	"company/biz/dal/mysql/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessagesDelete is the builder for deleting a Messages entity.
type MessagesDelete struct {
	config
	hooks    []Hook
	mutation *MessagesMutation
}

// Where appends a list predicates to the MessagesDelete builder.
func (md *MessagesDelete) Where(ps ...predicate.Messages) *MessagesDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MessagesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MessagesDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MessagesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(messages.Table, sqlgraph.NewFieldSpec(messages.FieldID, field.TypeInt64))
	_spec.Node.Schema = md.schemaConfig.Messages
	ctx = internal.NewSchemaConfigContext(ctx, md.schemaConfig)
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	md.mutation.done = true
	return affected, err
}

// MessagesDeleteOne is the builder for deleting a single Messages entity.
type MessagesDeleteOne struct {
	md *MessagesDelete
}

// Where appends a list predicates to the MessagesDelete builder.
func (mdo *MessagesDeleteOne) Where(ps ...predicate.Messages) *MessagesDeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MessagesDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{messages.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MessagesDeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}
