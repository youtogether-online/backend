// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wtkeqrf0/you_together/ent/predicate"
	"github.com/wtkeqrf0/you_together/ent/room"
	"github.com/wtkeqrf0/you_together/ent/user"
)

// RoomUpdate is the builder for updating Room entities.
type RoomUpdate struct {
	config
	hooks    []Hook
	mutation *RoomMutation
}

// Where appends a list predicates to the RoomUpdate builder.
func (ru *RoomUpdate) Where(ps ...predicate.Room) *RoomUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdateTime sets the "update_time" field.
func (ru *RoomUpdate) SetUpdateTime(t time.Time) *RoomUpdate {
	ru.mutation.SetUpdateTime(t)
	return ru
}

// SetCustomName sets the "custom_name" field.
func (ru *RoomUpdate) SetCustomName(s string) *RoomUpdate {
	ru.mutation.SetCustomName(s)
	return ru
}

// SetNillableCustomName sets the "custom_name" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableCustomName(s *string) *RoomUpdate {
	if s != nil {
		ru.SetCustomName(*s)
	}
	return ru
}

// ClearCustomName clears the value of the "custom_name" field.
func (ru *RoomUpdate) ClearCustomName() *RoomUpdate {
	ru.mutation.ClearCustomName()
	return ru
}

// SetPrivacy sets the "privacy" field.
func (ru *RoomUpdate) SetPrivacy(r room.Privacy) *RoomUpdate {
	ru.mutation.SetPrivacy(r)
	return ru
}

// SetNillablePrivacy sets the "privacy" field if the given value is not nil.
func (ru *RoomUpdate) SetNillablePrivacy(r *room.Privacy) *RoomUpdate {
	if r != nil {
		ru.SetPrivacy(*r)
	}
	return ru
}

// SetPasswordHash sets the "password_hash" field.
func (ru *RoomUpdate) SetPasswordHash(s string) *RoomUpdate {
	ru.mutation.SetPasswordHash(s)
	return ru
}

// SetNillablePasswordHash sets the "password_hash" field if the given value is not nil.
func (ru *RoomUpdate) SetNillablePasswordHash(s *string) *RoomUpdate {
	if s != nil {
		ru.SetPasswordHash(*s)
	}
	return ru
}

// ClearPasswordHash clears the value of the "password_hash" field.
func (ru *RoomUpdate) ClearPasswordHash() *RoomUpdate {
	ru.mutation.ClearPasswordHash()
	return ru
}

// SetHasChat sets the "has_chat" field.
func (ru *RoomUpdate) SetHasChat(b bool) *RoomUpdate {
	ru.mutation.SetHasChat(b)
	return ru
}

// SetNillableHasChat sets the "has_chat" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableHasChat(b *bool) *RoomUpdate {
	if b != nil {
		ru.SetHasChat(*b)
	}
	return ru
}

// SetDescription sets the "description" field.
func (ru *RoomUpdate) SetDescription(s string) *RoomUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableDescription(s *string) *RoomUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// ClearDescription clears the value of the "description" field.
func (ru *RoomUpdate) ClearDescription() *RoomUpdate {
	ru.mutation.ClearDescription()
	return ru
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (ru *RoomUpdate) AddUserIDs(ids ...string) *RoomUpdate {
	ru.mutation.AddUserIDs(ids...)
	return ru
}

// AddUsers adds the "users" edges to the User entity.
func (ru *RoomUpdate) AddUsers(u ...*User) *RoomUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ru.AddUserIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (ru *RoomUpdate) Mutation() *RoomMutation {
	return ru.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (ru *RoomUpdate) ClearUsers() *RoomUpdate {
	ru.mutation.ClearUsers()
	return ru
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (ru *RoomUpdate) RemoveUserIDs(ids ...string) *RoomUpdate {
	ru.mutation.RemoveUserIDs(ids...)
	return ru
}

// RemoveUsers removes "users" edges to User entities.
func (ru *RoomUpdate) RemoveUsers(u ...*User) *RoomUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ru.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoomUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks[int, RoomMutation](ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoomUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoomUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoomUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RoomUpdate) defaults() {
	if _, ok := ru.mutation.UpdateTime(); !ok {
		v := room.UpdateDefaultUpdateTime()
		ru.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoomUpdate) check() error {
	if v, ok := ru.mutation.CustomName(); ok {
		if err := room.CustomNameValidator(v); err != nil {
			return &ValidationError{Name: "custom_name", err: fmt.Errorf(`ent: validator failed for field "Room.custom_name": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Privacy(); ok {
		if err := room.PrivacyValidator(v); err != nil {
			return &ValidationError{Name: "privacy", err: fmt.Errorf(`ent: validator failed for field "Room.privacy": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Description(); ok {
		if err := room.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Room.description": %w`, err)}
		}
	}
	return nil
}

func (ru *RoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(room.Table, room.Columns, sqlgraph.NewFieldSpec(room.FieldID, field.TypeString))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdateTime(); ok {
		_spec.SetField(room.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ru.mutation.CustomName(); ok {
		_spec.SetField(room.FieldCustomName, field.TypeString, value)
	}
	if ru.mutation.CustomNameCleared() {
		_spec.ClearField(room.FieldCustomName, field.TypeString)
	}
	if value, ok := ru.mutation.Privacy(); ok {
		_spec.SetField(room.FieldPrivacy, field.TypeEnum, value)
	}
	if value, ok := ru.mutation.PasswordHash(); ok {
		_spec.SetField(room.FieldPasswordHash, field.TypeString, value)
	}
	if ru.mutation.PasswordHashCleared() {
		_spec.ClearField(room.FieldPasswordHash, field.TypeString)
	}
	if value, ok := ru.mutation.HasChat(); ok {
		_spec.SetField(room.FieldHasChat, field.TypeBool, value)
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.SetField(room.FieldDescription, field.TypeString, value)
	}
	if ru.mutation.DescriptionCleared() {
		_spec.ClearField(room.FieldDescription, field.TypeString)
	}
	if ru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   room.UsersTable,
			Columns: room.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedUsersIDs(); len(nodes) > 0 && !ru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   room.UsersTable,
			Columns: room.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   room.UsersTable,
			Columns: room.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoomUpdateOne is the builder for updating a single Room entity.
type RoomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoomMutation
}

// SetUpdateTime sets the "update_time" field.
func (ruo *RoomUpdateOne) SetUpdateTime(t time.Time) *RoomUpdateOne {
	ruo.mutation.SetUpdateTime(t)
	return ruo
}

// SetCustomName sets the "custom_name" field.
func (ruo *RoomUpdateOne) SetCustomName(s string) *RoomUpdateOne {
	ruo.mutation.SetCustomName(s)
	return ruo
}

// SetNillableCustomName sets the "custom_name" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableCustomName(s *string) *RoomUpdateOne {
	if s != nil {
		ruo.SetCustomName(*s)
	}
	return ruo
}

// ClearCustomName clears the value of the "custom_name" field.
func (ruo *RoomUpdateOne) ClearCustomName() *RoomUpdateOne {
	ruo.mutation.ClearCustomName()
	return ruo
}

// SetPrivacy sets the "privacy" field.
func (ruo *RoomUpdateOne) SetPrivacy(r room.Privacy) *RoomUpdateOne {
	ruo.mutation.SetPrivacy(r)
	return ruo
}

// SetNillablePrivacy sets the "privacy" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillablePrivacy(r *room.Privacy) *RoomUpdateOne {
	if r != nil {
		ruo.SetPrivacy(*r)
	}
	return ruo
}

// SetPasswordHash sets the "password_hash" field.
func (ruo *RoomUpdateOne) SetPasswordHash(s string) *RoomUpdateOne {
	ruo.mutation.SetPasswordHash(s)
	return ruo
}

// SetNillablePasswordHash sets the "password_hash" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillablePasswordHash(s *string) *RoomUpdateOne {
	if s != nil {
		ruo.SetPasswordHash(*s)
	}
	return ruo
}

// ClearPasswordHash clears the value of the "password_hash" field.
func (ruo *RoomUpdateOne) ClearPasswordHash() *RoomUpdateOne {
	ruo.mutation.ClearPasswordHash()
	return ruo
}

// SetHasChat sets the "has_chat" field.
func (ruo *RoomUpdateOne) SetHasChat(b bool) *RoomUpdateOne {
	ruo.mutation.SetHasChat(b)
	return ruo
}

// SetNillableHasChat sets the "has_chat" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableHasChat(b *bool) *RoomUpdateOne {
	if b != nil {
		ruo.SetHasChat(*b)
	}
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RoomUpdateOne) SetDescription(s string) *RoomUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableDescription(s *string) *RoomUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// ClearDescription clears the value of the "description" field.
func (ruo *RoomUpdateOne) ClearDescription() *RoomUpdateOne {
	ruo.mutation.ClearDescription()
	return ruo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (ruo *RoomUpdateOne) AddUserIDs(ids ...string) *RoomUpdateOne {
	ruo.mutation.AddUserIDs(ids...)
	return ruo
}

// AddUsers adds the "users" edges to the User entity.
func (ruo *RoomUpdateOne) AddUsers(u ...*User) *RoomUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ruo.AddUserIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (ruo *RoomUpdateOne) Mutation() *RoomMutation {
	return ruo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (ruo *RoomUpdateOne) ClearUsers() *RoomUpdateOne {
	ruo.mutation.ClearUsers()
	return ruo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (ruo *RoomUpdateOne) RemoveUserIDs(ids ...string) *RoomUpdateOne {
	ruo.mutation.RemoveUserIDs(ids...)
	return ruo
}

// RemoveUsers removes "users" edges to User entities.
func (ruo *RoomUpdateOne) RemoveUsers(u ...*User) *RoomUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ruo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the RoomUpdate builder.
func (ruo *RoomUpdateOne) Where(ps ...predicate.Room) *RoomUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoomUpdateOne) Select(field string, fields ...string) *RoomUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Room entity.
func (ruo *RoomUpdateOne) Save(ctx context.Context) (*Room, error) {
	ruo.defaults()
	return withHooks[*Room, RoomMutation](ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoomUpdateOne) SaveX(ctx context.Context) *Room {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoomUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoomUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RoomUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateTime(); !ok {
		v := room.UpdateDefaultUpdateTime()
		ruo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoomUpdateOne) check() error {
	if v, ok := ruo.mutation.CustomName(); ok {
		if err := room.CustomNameValidator(v); err != nil {
			return &ValidationError{Name: "custom_name", err: fmt.Errorf(`ent: validator failed for field "Room.custom_name": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Privacy(); ok {
		if err := room.PrivacyValidator(v); err != nil {
			return &ValidationError{Name: "privacy", err: fmt.Errorf(`ent: validator failed for field "Room.privacy": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Description(); ok {
		if err := room.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Room.description": %w`, err)}
		}
	}
	return nil
}

func (ruo *RoomUpdateOne) sqlSave(ctx context.Context) (_node *Room, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(room.Table, room.Columns, sqlgraph.NewFieldSpec(room.FieldID, field.TypeString))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Room.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, room.FieldID)
		for _, f := range fields {
			if !room.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != room.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdateTime(); ok {
		_spec.SetField(room.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.CustomName(); ok {
		_spec.SetField(room.FieldCustomName, field.TypeString, value)
	}
	if ruo.mutation.CustomNameCleared() {
		_spec.ClearField(room.FieldCustomName, field.TypeString)
	}
	if value, ok := ruo.mutation.Privacy(); ok {
		_spec.SetField(room.FieldPrivacy, field.TypeEnum, value)
	}
	if value, ok := ruo.mutation.PasswordHash(); ok {
		_spec.SetField(room.FieldPasswordHash, field.TypeString, value)
	}
	if ruo.mutation.PasswordHashCleared() {
		_spec.ClearField(room.FieldPasswordHash, field.TypeString)
	}
	if value, ok := ruo.mutation.HasChat(); ok {
		_spec.SetField(room.FieldHasChat, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.SetField(room.FieldDescription, field.TypeString, value)
	}
	if ruo.mutation.DescriptionCleared() {
		_spec.ClearField(room.FieldDescription, field.TypeString)
	}
	if ruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   room.UsersTable,
			Columns: room.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !ruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   room.UsersTable,
			Columns: room.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   room.UsersTable,
			Columns: room.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Room{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
