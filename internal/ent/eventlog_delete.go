// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TimeleapLabs/unchained/internal/ent/eventlog"
	"github.com/TimeleapLabs/unchained/internal/ent/predicate"
)

// EventLogDelete is the builder for deleting a EventLog entity.
type EventLogDelete struct {
	config
	hooks    []Hook
	mutation *EventLogMutation
}

// Where appends a list predicates to the EventLogDelete builder.
func (eld *EventLogDelete) Where(ps ...predicate.EventLog) *EventLogDelete {
	eld.mutation.Where(ps...)
	return eld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (eld *EventLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, eld.sqlExec, eld.mutation, eld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (eld *EventLogDelete) ExecX(ctx context.Context) int {
	n, err := eld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (eld *EventLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(eventlog.Table, sqlgraph.NewFieldSpec(eventlog.FieldID, field.TypeInt))
	if ps := eld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, eld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	eld.mutation.done = true
	return affected, err
}

// EventLogDeleteOne is the builder for deleting a single EventLog entity.
type EventLogDeleteOne struct {
	eld *EventLogDelete
}

// Where appends a list predicates to the EventLogDelete builder.
func (eldo *EventLogDeleteOne) Where(ps ...predicate.EventLog) *EventLogDeleteOne {
	eldo.eld.mutation.Where(ps...)
	return eldo
}

// Exec executes the deletion query.
func (eldo *EventLogDeleteOne) Exec(ctx context.Context) error {
	n, err := eldo.eld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{eventlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eldo *EventLogDeleteOne) ExecX(ctx context.Context) {
	if err := eldo.Exec(ctx); err != nil {
		panic(err)
	}
}
