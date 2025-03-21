// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"schedule/biz/dal/mysql/ent/predicate"
	"schedule/biz/dal/mysql/ent/schedule"
	"schedule/biz/dal/mysql/ent/schedulecoach"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScheduleCoachUpdate is the builder for updating ScheduleCoach entities.
type ScheduleCoachUpdate struct {
	config
	hooks    []Hook
	mutation *ScheduleCoachMutation
}

// Where appends a list predicates to the ScheduleCoachUpdate builder.
func (scu *ScheduleCoachUpdate) Where(ps ...predicate.ScheduleCoach) *ScheduleCoachUpdate {
	scu.mutation.Where(ps...)
	return scu
}

// SetUpdatedAt sets the "updated_at" field.
func (scu *ScheduleCoachUpdate) SetUpdatedAt(t time.Time) *ScheduleCoachUpdate {
	scu.mutation.SetUpdatedAt(t)
	return scu
}

// SetStatus sets the "status" field.
func (scu *ScheduleCoachUpdate) SetStatus(i int64) *ScheduleCoachUpdate {
	scu.mutation.ResetStatus()
	scu.mutation.SetStatus(i)
	return scu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (scu *ScheduleCoachUpdate) SetNillableStatus(i *int64) *ScheduleCoachUpdate {
	if i != nil {
		scu.SetStatus(*i)
	}
	return scu
}

// AddStatus adds i to the "status" field.
func (scu *ScheduleCoachUpdate) AddStatus(i int64) *ScheduleCoachUpdate {
	scu.mutation.AddStatus(i)
	return scu
}

// ClearStatus clears the value of the "status" field.
func (scu *ScheduleCoachUpdate) ClearStatus() *ScheduleCoachUpdate {
	scu.mutation.ClearStatus()
	return scu
}

// SetVenueID sets the "venue_id" field.
func (scu *ScheduleCoachUpdate) SetVenueID(i int64) *ScheduleCoachUpdate {
	scu.mutation.ResetVenueID()
	scu.mutation.SetVenueID(i)
	return scu
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (scu *ScheduleCoachUpdate) SetNillableVenueID(i *int64) *ScheduleCoachUpdate {
	if i != nil {
		scu.SetVenueID(*i)
	}
	return scu
}

// AddVenueID adds i to the "venue_id" field.
func (scu *ScheduleCoachUpdate) AddVenueID(i int64) *ScheduleCoachUpdate {
	scu.mutation.AddVenueID(i)
	return scu
}

// ClearVenueID clears the value of the "venue_id" field.
func (scu *ScheduleCoachUpdate) ClearVenueID() *ScheduleCoachUpdate {
	scu.mutation.ClearVenueID()
	return scu
}

// SetCoachID sets the "coach_id" field.
func (scu *ScheduleCoachUpdate) SetCoachID(i int64) *ScheduleCoachUpdate {
	scu.mutation.ResetCoachID()
	scu.mutation.SetCoachID(i)
	return scu
}

// SetNillableCoachID sets the "coach_id" field if the given value is not nil.
func (scu *ScheduleCoachUpdate) SetNillableCoachID(i *int64) *ScheduleCoachUpdate {
	if i != nil {
		scu.SetCoachID(*i)
	}
	return scu
}

// AddCoachID adds i to the "coach_id" field.
func (scu *ScheduleCoachUpdate) AddCoachID(i int64) *ScheduleCoachUpdate {
	scu.mutation.AddCoachID(i)
	return scu
}

// ClearCoachID clears the value of the "coach_id" field.
func (scu *ScheduleCoachUpdate) ClearCoachID() *ScheduleCoachUpdate {
	scu.mutation.ClearCoachID()
	return scu
}

// SetScheduleID sets the "schedule_id" field.
func (scu *ScheduleCoachUpdate) SetScheduleID(i int64) *ScheduleCoachUpdate {
	scu.mutation.SetScheduleID(i)
	return scu
}

// SetNillableScheduleID sets the "schedule_id" field if the given value is not nil.
func (scu *ScheduleCoachUpdate) SetNillableScheduleID(i *int64) *ScheduleCoachUpdate {
	if i != nil {
		scu.SetScheduleID(*i)
	}
	return scu
}

// ClearScheduleID clears the value of the "schedule_id" field.
func (scu *ScheduleCoachUpdate) ClearScheduleID() *ScheduleCoachUpdate {
	scu.mutation.ClearScheduleID()
	return scu
}

// SetScheduleName sets the "schedule_name" field.
func (scu *ScheduleCoachUpdate) SetScheduleName(s string) *ScheduleCoachUpdate {
	scu.mutation.SetScheduleName(s)
	return scu
}

// SetNillableScheduleName sets the "schedule_name" field if the given value is not nil.
func (scu *ScheduleCoachUpdate) SetNillableScheduleName(s *string) *ScheduleCoachUpdate {
	if s != nil {
		scu.SetScheduleName(*s)
	}
	return scu
}

// ClearScheduleName clears the value of the "schedule_name" field.
func (scu *ScheduleCoachUpdate) ClearScheduleName() *ScheduleCoachUpdate {
	scu.mutation.ClearScheduleName()
	return scu
}

// SetType sets the "type" field.
func (scu *ScheduleCoachUpdate) SetType(s string) *ScheduleCoachUpdate {
	scu.mutation.SetType(s)
	return scu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (scu *ScheduleCoachUpdate) SetNillableType(s *string) *ScheduleCoachUpdate {
	if s != nil {
		scu.SetType(*s)
	}
	return scu
}

// ClearType clears the value of the "type" field.
func (scu *ScheduleCoachUpdate) ClearType() *ScheduleCoachUpdate {
	scu.mutation.ClearType()
	return scu
}

// SetStartTime sets the "start_time" field.
func (scu *ScheduleCoachUpdate) SetStartTime(t time.Time) *ScheduleCoachUpdate {
	scu.mutation.SetStartTime(t)
	return scu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (scu *ScheduleCoachUpdate) SetNillableStartTime(t *time.Time) *ScheduleCoachUpdate {
	if t != nil {
		scu.SetStartTime(*t)
	}
	return scu
}

// ClearStartTime clears the value of the "start_time" field.
func (scu *ScheduleCoachUpdate) ClearStartTime() *ScheduleCoachUpdate {
	scu.mutation.ClearStartTime()
	return scu
}

// SetEndTime sets the "end_time" field.
func (scu *ScheduleCoachUpdate) SetEndTime(t time.Time) *ScheduleCoachUpdate {
	scu.mutation.SetEndTime(t)
	return scu
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (scu *ScheduleCoachUpdate) SetNillableEndTime(t *time.Time) *ScheduleCoachUpdate {
	if t != nil {
		scu.SetEndTime(*t)
	}
	return scu
}

// ClearEndTime clears the value of the "end_time" field.
func (scu *ScheduleCoachUpdate) ClearEndTime() *ScheduleCoachUpdate {
	scu.mutation.ClearEndTime()
	return scu
}

// SetSignStartTime sets the "sign_start_time" field.
func (scu *ScheduleCoachUpdate) SetSignStartTime(t time.Time) *ScheduleCoachUpdate {
	scu.mutation.SetSignStartTime(t)
	return scu
}

// SetNillableSignStartTime sets the "sign_start_time" field if the given value is not nil.
func (scu *ScheduleCoachUpdate) SetNillableSignStartTime(t *time.Time) *ScheduleCoachUpdate {
	if t != nil {
		scu.SetSignStartTime(*t)
	}
	return scu
}

// ClearSignStartTime clears the value of the "sign_start_time" field.
func (scu *ScheduleCoachUpdate) ClearSignStartTime() *ScheduleCoachUpdate {
	scu.mutation.ClearSignStartTime()
	return scu
}

// SetSignEndTime sets the "sign_end_time" field.
func (scu *ScheduleCoachUpdate) SetSignEndTime(t time.Time) *ScheduleCoachUpdate {
	scu.mutation.SetSignEndTime(t)
	return scu
}

// SetNillableSignEndTime sets the "sign_end_time" field if the given value is not nil.
func (scu *ScheduleCoachUpdate) SetNillableSignEndTime(t *time.Time) *ScheduleCoachUpdate {
	if t != nil {
		scu.SetSignEndTime(*t)
	}
	return scu
}

// ClearSignEndTime clears the value of the "sign_end_time" field.
func (scu *ScheduleCoachUpdate) ClearSignEndTime() *ScheduleCoachUpdate {
	scu.mutation.ClearSignEndTime()
	return scu
}

// SetCoachName sets the "coach_name" field.
func (scu *ScheduleCoachUpdate) SetCoachName(s string) *ScheduleCoachUpdate {
	scu.mutation.SetCoachName(s)
	return scu
}

// SetNillableCoachName sets the "coach_name" field if the given value is not nil.
func (scu *ScheduleCoachUpdate) SetNillableCoachName(s *string) *ScheduleCoachUpdate {
	if s != nil {
		scu.SetCoachName(*s)
	}
	return scu
}

// ClearCoachName clears the value of the "coach_name" field.
func (scu *ScheduleCoachUpdate) ClearCoachName() *ScheduleCoachUpdate {
	scu.mutation.ClearCoachName()
	return scu
}

// SetSchedule sets the "schedule" edge to the Schedule entity.
func (scu *ScheduleCoachUpdate) SetSchedule(s *Schedule) *ScheduleCoachUpdate {
	return scu.SetScheduleID(s.ID)
}

// Mutation returns the ScheduleCoachMutation object of the builder.
func (scu *ScheduleCoachUpdate) Mutation() *ScheduleCoachMutation {
	return scu.mutation
}

// ClearSchedule clears the "schedule" edge to the Schedule entity.
func (scu *ScheduleCoachUpdate) ClearSchedule() *ScheduleCoachUpdate {
	scu.mutation.ClearSchedule()
	return scu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (scu *ScheduleCoachUpdate) Save(ctx context.Context) (int, error) {
	scu.defaults()
	return withHooks(ctx, scu.sqlSave, scu.mutation, scu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scu *ScheduleCoachUpdate) SaveX(ctx context.Context) int {
	affected, err := scu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (scu *ScheduleCoachUpdate) Exec(ctx context.Context) error {
	_, err := scu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scu *ScheduleCoachUpdate) ExecX(ctx context.Context) {
	if err := scu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scu *ScheduleCoachUpdate) defaults() {
	if _, ok := scu.mutation.UpdatedAt(); !ok {
		v := schedulecoach.UpdateDefaultUpdatedAt()
		scu.mutation.SetUpdatedAt(v)
	}
}

func (scu *ScheduleCoachUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(schedulecoach.Table, schedulecoach.Columns, sqlgraph.NewFieldSpec(schedulecoach.FieldID, field.TypeInt64))
	if ps := scu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scu.mutation.UpdatedAt(); ok {
		_spec.SetField(schedulecoach.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := scu.mutation.Status(); ok {
		_spec.SetField(schedulecoach.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := scu.mutation.AddedStatus(); ok {
		_spec.AddField(schedulecoach.FieldStatus, field.TypeInt64, value)
	}
	if scu.mutation.StatusCleared() {
		_spec.ClearField(schedulecoach.FieldStatus, field.TypeInt64)
	}
	if value, ok := scu.mutation.VenueID(); ok {
		_spec.SetField(schedulecoach.FieldVenueID, field.TypeInt64, value)
	}
	if value, ok := scu.mutation.AddedVenueID(); ok {
		_spec.AddField(schedulecoach.FieldVenueID, field.TypeInt64, value)
	}
	if scu.mutation.VenueIDCleared() {
		_spec.ClearField(schedulecoach.FieldVenueID, field.TypeInt64)
	}
	if value, ok := scu.mutation.CoachID(); ok {
		_spec.SetField(schedulecoach.FieldCoachID, field.TypeInt64, value)
	}
	if value, ok := scu.mutation.AddedCoachID(); ok {
		_spec.AddField(schedulecoach.FieldCoachID, field.TypeInt64, value)
	}
	if scu.mutation.CoachIDCleared() {
		_spec.ClearField(schedulecoach.FieldCoachID, field.TypeInt64)
	}
	if value, ok := scu.mutation.ScheduleName(); ok {
		_spec.SetField(schedulecoach.FieldScheduleName, field.TypeString, value)
	}
	if scu.mutation.ScheduleNameCleared() {
		_spec.ClearField(schedulecoach.FieldScheduleName, field.TypeString)
	}
	if value, ok := scu.mutation.GetType(); ok {
		_spec.SetField(schedulecoach.FieldType, field.TypeString, value)
	}
	if scu.mutation.TypeCleared() {
		_spec.ClearField(schedulecoach.FieldType, field.TypeString)
	}
	if value, ok := scu.mutation.StartTime(); ok {
		_spec.SetField(schedulecoach.FieldStartTime, field.TypeTime, value)
	}
	if scu.mutation.StartTimeCleared() {
		_spec.ClearField(schedulecoach.FieldStartTime, field.TypeTime)
	}
	if value, ok := scu.mutation.EndTime(); ok {
		_spec.SetField(schedulecoach.FieldEndTime, field.TypeTime, value)
	}
	if scu.mutation.EndTimeCleared() {
		_spec.ClearField(schedulecoach.FieldEndTime, field.TypeTime)
	}
	if value, ok := scu.mutation.SignStartTime(); ok {
		_spec.SetField(schedulecoach.FieldSignStartTime, field.TypeTime, value)
	}
	if scu.mutation.SignStartTimeCleared() {
		_spec.ClearField(schedulecoach.FieldSignStartTime, field.TypeTime)
	}
	if value, ok := scu.mutation.SignEndTime(); ok {
		_spec.SetField(schedulecoach.FieldSignEndTime, field.TypeTime, value)
	}
	if scu.mutation.SignEndTimeCleared() {
		_spec.ClearField(schedulecoach.FieldSignEndTime, field.TypeTime)
	}
	if value, ok := scu.mutation.CoachName(); ok {
		_spec.SetField(schedulecoach.FieldCoachName, field.TypeString, value)
	}
	if scu.mutation.CoachNameCleared() {
		_spec.ClearField(schedulecoach.FieldCoachName, field.TypeString)
	}
	if scu.mutation.ScheduleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedulecoach.ScheduleTable,
			Columns: []string{schedulecoach.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schedule.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedulecoach.ScheduleTable,
			Columns: []string{schedulecoach.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schedule.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, scu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schedulecoach.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	scu.mutation.done = true
	return n, nil
}

// ScheduleCoachUpdateOne is the builder for updating a single ScheduleCoach entity.
type ScheduleCoachUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ScheduleCoachMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (scuo *ScheduleCoachUpdateOne) SetUpdatedAt(t time.Time) *ScheduleCoachUpdateOne {
	scuo.mutation.SetUpdatedAt(t)
	return scuo
}

// SetStatus sets the "status" field.
func (scuo *ScheduleCoachUpdateOne) SetStatus(i int64) *ScheduleCoachUpdateOne {
	scuo.mutation.ResetStatus()
	scuo.mutation.SetStatus(i)
	return scuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (scuo *ScheduleCoachUpdateOne) SetNillableStatus(i *int64) *ScheduleCoachUpdateOne {
	if i != nil {
		scuo.SetStatus(*i)
	}
	return scuo
}

// AddStatus adds i to the "status" field.
func (scuo *ScheduleCoachUpdateOne) AddStatus(i int64) *ScheduleCoachUpdateOne {
	scuo.mutation.AddStatus(i)
	return scuo
}

// ClearStatus clears the value of the "status" field.
func (scuo *ScheduleCoachUpdateOne) ClearStatus() *ScheduleCoachUpdateOne {
	scuo.mutation.ClearStatus()
	return scuo
}

// SetVenueID sets the "venue_id" field.
func (scuo *ScheduleCoachUpdateOne) SetVenueID(i int64) *ScheduleCoachUpdateOne {
	scuo.mutation.ResetVenueID()
	scuo.mutation.SetVenueID(i)
	return scuo
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (scuo *ScheduleCoachUpdateOne) SetNillableVenueID(i *int64) *ScheduleCoachUpdateOne {
	if i != nil {
		scuo.SetVenueID(*i)
	}
	return scuo
}

// AddVenueID adds i to the "venue_id" field.
func (scuo *ScheduleCoachUpdateOne) AddVenueID(i int64) *ScheduleCoachUpdateOne {
	scuo.mutation.AddVenueID(i)
	return scuo
}

// ClearVenueID clears the value of the "venue_id" field.
func (scuo *ScheduleCoachUpdateOne) ClearVenueID() *ScheduleCoachUpdateOne {
	scuo.mutation.ClearVenueID()
	return scuo
}

// SetCoachID sets the "coach_id" field.
func (scuo *ScheduleCoachUpdateOne) SetCoachID(i int64) *ScheduleCoachUpdateOne {
	scuo.mutation.ResetCoachID()
	scuo.mutation.SetCoachID(i)
	return scuo
}

// SetNillableCoachID sets the "coach_id" field if the given value is not nil.
func (scuo *ScheduleCoachUpdateOne) SetNillableCoachID(i *int64) *ScheduleCoachUpdateOne {
	if i != nil {
		scuo.SetCoachID(*i)
	}
	return scuo
}

// AddCoachID adds i to the "coach_id" field.
func (scuo *ScheduleCoachUpdateOne) AddCoachID(i int64) *ScheduleCoachUpdateOne {
	scuo.mutation.AddCoachID(i)
	return scuo
}

// ClearCoachID clears the value of the "coach_id" field.
func (scuo *ScheduleCoachUpdateOne) ClearCoachID() *ScheduleCoachUpdateOne {
	scuo.mutation.ClearCoachID()
	return scuo
}

// SetScheduleID sets the "schedule_id" field.
func (scuo *ScheduleCoachUpdateOne) SetScheduleID(i int64) *ScheduleCoachUpdateOne {
	scuo.mutation.SetScheduleID(i)
	return scuo
}

// SetNillableScheduleID sets the "schedule_id" field if the given value is not nil.
func (scuo *ScheduleCoachUpdateOne) SetNillableScheduleID(i *int64) *ScheduleCoachUpdateOne {
	if i != nil {
		scuo.SetScheduleID(*i)
	}
	return scuo
}

// ClearScheduleID clears the value of the "schedule_id" field.
func (scuo *ScheduleCoachUpdateOne) ClearScheduleID() *ScheduleCoachUpdateOne {
	scuo.mutation.ClearScheduleID()
	return scuo
}

// SetScheduleName sets the "schedule_name" field.
func (scuo *ScheduleCoachUpdateOne) SetScheduleName(s string) *ScheduleCoachUpdateOne {
	scuo.mutation.SetScheduleName(s)
	return scuo
}

// SetNillableScheduleName sets the "schedule_name" field if the given value is not nil.
func (scuo *ScheduleCoachUpdateOne) SetNillableScheduleName(s *string) *ScheduleCoachUpdateOne {
	if s != nil {
		scuo.SetScheduleName(*s)
	}
	return scuo
}

// ClearScheduleName clears the value of the "schedule_name" field.
func (scuo *ScheduleCoachUpdateOne) ClearScheduleName() *ScheduleCoachUpdateOne {
	scuo.mutation.ClearScheduleName()
	return scuo
}

// SetType sets the "type" field.
func (scuo *ScheduleCoachUpdateOne) SetType(s string) *ScheduleCoachUpdateOne {
	scuo.mutation.SetType(s)
	return scuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (scuo *ScheduleCoachUpdateOne) SetNillableType(s *string) *ScheduleCoachUpdateOne {
	if s != nil {
		scuo.SetType(*s)
	}
	return scuo
}

// ClearType clears the value of the "type" field.
func (scuo *ScheduleCoachUpdateOne) ClearType() *ScheduleCoachUpdateOne {
	scuo.mutation.ClearType()
	return scuo
}

// SetStartTime sets the "start_time" field.
func (scuo *ScheduleCoachUpdateOne) SetStartTime(t time.Time) *ScheduleCoachUpdateOne {
	scuo.mutation.SetStartTime(t)
	return scuo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (scuo *ScheduleCoachUpdateOne) SetNillableStartTime(t *time.Time) *ScheduleCoachUpdateOne {
	if t != nil {
		scuo.SetStartTime(*t)
	}
	return scuo
}

// ClearStartTime clears the value of the "start_time" field.
func (scuo *ScheduleCoachUpdateOne) ClearStartTime() *ScheduleCoachUpdateOne {
	scuo.mutation.ClearStartTime()
	return scuo
}

// SetEndTime sets the "end_time" field.
func (scuo *ScheduleCoachUpdateOne) SetEndTime(t time.Time) *ScheduleCoachUpdateOne {
	scuo.mutation.SetEndTime(t)
	return scuo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (scuo *ScheduleCoachUpdateOne) SetNillableEndTime(t *time.Time) *ScheduleCoachUpdateOne {
	if t != nil {
		scuo.SetEndTime(*t)
	}
	return scuo
}

// ClearEndTime clears the value of the "end_time" field.
func (scuo *ScheduleCoachUpdateOne) ClearEndTime() *ScheduleCoachUpdateOne {
	scuo.mutation.ClearEndTime()
	return scuo
}

// SetSignStartTime sets the "sign_start_time" field.
func (scuo *ScheduleCoachUpdateOne) SetSignStartTime(t time.Time) *ScheduleCoachUpdateOne {
	scuo.mutation.SetSignStartTime(t)
	return scuo
}

// SetNillableSignStartTime sets the "sign_start_time" field if the given value is not nil.
func (scuo *ScheduleCoachUpdateOne) SetNillableSignStartTime(t *time.Time) *ScheduleCoachUpdateOne {
	if t != nil {
		scuo.SetSignStartTime(*t)
	}
	return scuo
}

// ClearSignStartTime clears the value of the "sign_start_time" field.
func (scuo *ScheduleCoachUpdateOne) ClearSignStartTime() *ScheduleCoachUpdateOne {
	scuo.mutation.ClearSignStartTime()
	return scuo
}

// SetSignEndTime sets the "sign_end_time" field.
func (scuo *ScheduleCoachUpdateOne) SetSignEndTime(t time.Time) *ScheduleCoachUpdateOne {
	scuo.mutation.SetSignEndTime(t)
	return scuo
}

// SetNillableSignEndTime sets the "sign_end_time" field if the given value is not nil.
func (scuo *ScheduleCoachUpdateOne) SetNillableSignEndTime(t *time.Time) *ScheduleCoachUpdateOne {
	if t != nil {
		scuo.SetSignEndTime(*t)
	}
	return scuo
}

// ClearSignEndTime clears the value of the "sign_end_time" field.
func (scuo *ScheduleCoachUpdateOne) ClearSignEndTime() *ScheduleCoachUpdateOne {
	scuo.mutation.ClearSignEndTime()
	return scuo
}

// SetCoachName sets the "coach_name" field.
func (scuo *ScheduleCoachUpdateOne) SetCoachName(s string) *ScheduleCoachUpdateOne {
	scuo.mutation.SetCoachName(s)
	return scuo
}

// SetNillableCoachName sets the "coach_name" field if the given value is not nil.
func (scuo *ScheduleCoachUpdateOne) SetNillableCoachName(s *string) *ScheduleCoachUpdateOne {
	if s != nil {
		scuo.SetCoachName(*s)
	}
	return scuo
}

// ClearCoachName clears the value of the "coach_name" field.
func (scuo *ScheduleCoachUpdateOne) ClearCoachName() *ScheduleCoachUpdateOne {
	scuo.mutation.ClearCoachName()
	return scuo
}

// SetSchedule sets the "schedule" edge to the Schedule entity.
func (scuo *ScheduleCoachUpdateOne) SetSchedule(s *Schedule) *ScheduleCoachUpdateOne {
	return scuo.SetScheduleID(s.ID)
}

// Mutation returns the ScheduleCoachMutation object of the builder.
func (scuo *ScheduleCoachUpdateOne) Mutation() *ScheduleCoachMutation {
	return scuo.mutation
}

// ClearSchedule clears the "schedule" edge to the Schedule entity.
func (scuo *ScheduleCoachUpdateOne) ClearSchedule() *ScheduleCoachUpdateOne {
	scuo.mutation.ClearSchedule()
	return scuo
}

// Where appends a list predicates to the ScheduleCoachUpdate builder.
func (scuo *ScheduleCoachUpdateOne) Where(ps ...predicate.ScheduleCoach) *ScheduleCoachUpdateOne {
	scuo.mutation.Where(ps...)
	return scuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (scuo *ScheduleCoachUpdateOne) Select(field string, fields ...string) *ScheduleCoachUpdateOne {
	scuo.fields = append([]string{field}, fields...)
	return scuo
}

// Save executes the query and returns the updated ScheduleCoach entity.
func (scuo *ScheduleCoachUpdateOne) Save(ctx context.Context) (*ScheduleCoach, error) {
	scuo.defaults()
	return withHooks(ctx, scuo.sqlSave, scuo.mutation, scuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scuo *ScheduleCoachUpdateOne) SaveX(ctx context.Context) *ScheduleCoach {
	node, err := scuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (scuo *ScheduleCoachUpdateOne) Exec(ctx context.Context) error {
	_, err := scuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scuo *ScheduleCoachUpdateOne) ExecX(ctx context.Context) {
	if err := scuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scuo *ScheduleCoachUpdateOne) defaults() {
	if _, ok := scuo.mutation.UpdatedAt(); !ok {
		v := schedulecoach.UpdateDefaultUpdatedAt()
		scuo.mutation.SetUpdatedAt(v)
	}
}

func (scuo *ScheduleCoachUpdateOne) sqlSave(ctx context.Context) (_node *ScheduleCoach, err error) {
	_spec := sqlgraph.NewUpdateSpec(schedulecoach.Table, schedulecoach.Columns, sqlgraph.NewFieldSpec(schedulecoach.FieldID, field.TypeInt64))
	id, ok := scuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ScheduleCoach.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := scuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, schedulecoach.FieldID)
		for _, f := range fields {
			if !schedulecoach.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != schedulecoach.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := scuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scuo.mutation.UpdatedAt(); ok {
		_spec.SetField(schedulecoach.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := scuo.mutation.Status(); ok {
		_spec.SetField(schedulecoach.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := scuo.mutation.AddedStatus(); ok {
		_spec.AddField(schedulecoach.FieldStatus, field.TypeInt64, value)
	}
	if scuo.mutation.StatusCleared() {
		_spec.ClearField(schedulecoach.FieldStatus, field.TypeInt64)
	}
	if value, ok := scuo.mutation.VenueID(); ok {
		_spec.SetField(schedulecoach.FieldVenueID, field.TypeInt64, value)
	}
	if value, ok := scuo.mutation.AddedVenueID(); ok {
		_spec.AddField(schedulecoach.FieldVenueID, field.TypeInt64, value)
	}
	if scuo.mutation.VenueIDCleared() {
		_spec.ClearField(schedulecoach.FieldVenueID, field.TypeInt64)
	}
	if value, ok := scuo.mutation.CoachID(); ok {
		_spec.SetField(schedulecoach.FieldCoachID, field.TypeInt64, value)
	}
	if value, ok := scuo.mutation.AddedCoachID(); ok {
		_spec.AddField(schedulecoach.FieldCoachID, field.TypeInt64, value)
	}
	if scuo.mutation.CoachIDCleared() {
		_spec.ClearField(schedulecoach.FieldCoachID, field.TypeInt64)
	}
	if value, ok := scuo.mutation.ScheduleName(); ok {
		_spec.SetField(schedulecoach.FieldScheduleName, field.TypeString, value)
	}
	if scuo.mutation.ScheduleNameCleared() {
		_spec.ClearField(schedulecoach.FieldScheduleName, field.TypeString)
	}
	if value, ok := scuo.mutation.GetType(); ok {
		_spec.SetField(schedulecoach.FieldType, field.TypeString, value)
	}
	if scuo.mutation.TypeCleared() {
		_spec.ClearField(schedulecoach.FieldType, field.TypeString)
	}
	if value, ok := scuo.mutation.StartTime(); ok {
		_spec.SetField(schedulecoach.FieldStartTime, field.TypeTime, value)
	}
	if scuo.mutation.StartTimeCleared() {
		_spec.ClearField(schedulecoach.FieldStartTime, field.TypeTime)
	}
	if value, ok := scuo.mutation.EndTime(); ok {
		_spec.SetField(schedulecoach.FieldEndTime, field.TypeTime, value)
	}
	if scuo.mutation.EndTimeCleared() {
		_spec.ClearField(schedulecoach.FieldEndTime, field.TypeTime)
	}
	if value, ok := scuo.mutation.SignStartTime(); ok {
		_spec.SetField(schedulecoach.FieldSignStartTime, field.TypeTime, value)
	}
	if scuo.mutation.SignStartTimeCleared() {
		_spec.ClearField(schedulecoach.FieldSignStartTime, field.TypeTime)
	}
	if value, ok := scuo.mutation.SignEndTime(); ok {
		_spec.SetField(schedulecoach.FieldSignEndTime, field.TypeTime, value)
	}
	if scuo.mutation.SignEndTimeCleared() {
		_spec.ClearField(schedulecoach.FieldSignEndTime, field.TypeTime)
	}
	if value, ok := scuo.mutation.CoachName(); ok {
		_spec.SetField(schedulecoach.FieldCoachName, field.TypeString, value)
	}
	if scuo.mutation.CoachNameCleared() {
		_spec.ClearField(schedulecoach.FieldCoachName, field.TypeString)
	}
	if scuo.mutation.ScheduleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedulecoach.ScheduleTable,
			Columns: []string{schedulecoach.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schedule.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedulecoach.ScheduleTable,
			Columns: []string{schedulecoach.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schedule.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ScheduleCoach{config: scuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, scuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schedulecoach.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	scuo.mutation.done = true
	return _node, nil
}
