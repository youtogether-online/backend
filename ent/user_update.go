// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/wtkeqrf0/you-together/ent/predicate"
	"github.com/wtkeqrf0/you-together/ent/room"
	"github.com/wtkeqrf0/you-together/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdateTime sets the "update_time" field.
func (uu *UserUpdate) SetUpdateTime(t time.Time) *UserUpdate {
	uu.mutation.SetUpdateTime(t)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetIsEmailVerified sets the "is_email_verified" field.
func (uu *UserUpdate) SetIsEmailVerified(b bool) *UserUpdate {
	uu.mutation.SetIsEmailVerified(b)
	return uu
}

// SetNillableIsEmailVerified sets the "is_email_verified" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsEmailVerified(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsEmailVerified(*b)
	}
	return uu
}

// SetPasswordHash sets the "password_hash" field.
func (uu *UserUpdate) SetPasswordHash(b []byte) *UserUpdate {
	uu.mutation.SetPasswordHash(b)
	return uu
}

// ClearPasswordHash clears the value of the "password_hash" field.
func (uu *UserUpdate) ClearPasswordHash() *UserUpdate {
	uu.mutation.ClearPasswordHash()
	return uu
}

// SetBiography sets the "biography" field.
func (uu *UserUpdate) SetBiography(s string) *UserUpdate {
	uu.mutation.SetBiography(s)
	return uu
}

// SetNillableBiography sets the "biography" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBiography(s *string) *UserUpdate {
	if s != nil {
		uu.SetBiography(*s)
	}
	return uu
}

// ClearBiography clears the value of the "biography" field.
func (uu *UserUpdate) ClearBiography() *UserUpdate {
	uu.mutation.ClearBiography()
	return uu
}

// SetRole sets the "role" field.
func (uu *UserUpdate) SetRole(s string) *UserUpdate {
	uu.mutation.SetRole(s)
	return uu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(s *string) *UserUpdate {
	if s != nil {
		uu.SetRole(*s)
	}
	return uu
}

// SetFriendsIds sets the "friends_ids" field.
func (uu *UserUpdate) SetFriendsIds(s []string) *UserUpdate {
	uu.mutation.SetFriendsIds(s)
	return uu
}

// AppendFriendsIds appends s to the "friends_ids" field.
func (uu *UserUpdate) AppendFriendsIds(s []string) *UserUpdate {
	uu.mutation.AppendFriendsIds(s)
	return uu
}

// ClearFriendsIds clears the value of the "friends_ids" field.
func (uu *UserUpdate) ClearFriendsIds() *UserUpdate {
	uu.mutation.ClearFriendsIds()
	return uu
}

// SetLanguage sets the "language" field.
func (uu *UserUpdate) SetLanguage(s string) *UserUpdate {
	uu.mutation.SetLanguage(s)
	return uu
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLanguage(s *string) *UserUpdate {
	if s != nil {
		uu.SetLanguage(*s)
	}
	return uu
}

// SetTheme sets the "theme" field.
func (uu *UserUpdate) SetTheme(s string) *UserUpdate {
	uu.mutation.SetTheme(s)
	return uu
}

// SetNillableTheme sets the "theme" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTheme(s *string) *UserUpdate {
	if s != nil {
		uu.SetTheme(*s)
	}
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFirstName(s *string) *UserUpdate {
	if s != nil {
		uu.SetFirstName(*s)
	}
	return uu
}

// ClearFirstName clears the value of the "first_name" field.
func (uu *UserUpdate) ClearFirstName() *UserUpdate {
	uu.mutation.ClearFirstName()
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastName(s *string) *UserUpdate {
	if s != nil {
		uu.SetLastName(*s)
	}
	return uu
}

// ClearLastName clears the value of the "last_name" field.
func (uu *UserUpdate) ClearLastName() *UserUpdate {
	uu.mutation.ClearLastName()
	return uu
}

// SetSessions sets the "sessions" field.
func (uu *UserUpdate) SetSessions(s []string) *UserUpdate {
	uu.mutation.SetSessions(s)
	return uu
}

// AppendSessions appends s to the "sessions" field.
func (uu *UserUpdate) AppendSessions(s []string) *UserUpdate {
	uu.mutation.AppendSessions(s)
	return uu
}

// ClearSessions clears the value of the "sessions" field.
func (uu *UserUpdate) ClearSessions() *UserUpdate {
	uu.mutation.ClearSessions()
	return uu
}

// AddRoomIDs adds the "rooms" edge to the Room entity by IDs.
func (uu *UserUpdate) AddRoomIDs(ids ...int) *UserUpdate {
	uu.mutation.AddRoomIDs(ids...)
	return uu
}

// AddRooms adds the "rooms" edges to the Room entity.
func (uu *UserUpdate) AddRooms(r ...*Room) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddRoomIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearRooms clears all "rooms" edges to the Room entity.
func (uu *UserUpdate) ClearRooms() *UserUpdate {
	uu.mutation.ClearRooms()
	return uu
}

// RemoveRoomIDs removes the "rooms" edge to Room entities by IDs.
func (uu *UserUpdate) RemoveRoomIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveRoomIDs(ids...)
	return uu
}

// RemoveRooms removes "rooms" edges to Room entities.
func (uu *UserUpdate) RemoveRooms(r ...*Room) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveRoomIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdateTime(); !ok {
		v := user.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Biography(); ok {
		if err := user.BiographyValidator(v); err != nil {
			return &ValidationError{Name: "biography", err: fmt.Errorf(`ent: validator failed for field "User.biography": %w`, err)}
		}
	}
	if v, ok := uu.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.IsEmailVerified(); ok {
		_spec.SetField(user.FieldIsEmailVerified, field.TypeBool, value)
	}
	if value, ok := uu.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeBytes, value)
	}
	if uu.mutation.PasswordHashCleared() {
		_spec.ClearField(user.FieldPasswordHash, field.TypeBytes)
	}
	if value, ok := uu.mutation.Biography(); ok {
		_spec.SetField(user.FieldBiography, field.TypeString, value)
	}
	if uu.mutation.BiographyCleared() {
		_spec.ClearField(user.FieldBiography, field.TypeString)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeString, value)
	}
	if value, ok := uu.mutation.FriendsIds(); ok {
		_spec.SetField(user.FieldFriendsIds, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedFriendsIds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldFriendsIds, value)
		})
	}
	if uu.mutation.FriendsIdsCleared() {
		_spec.ClearField(user.FieldFriendsIds, field.TypeJSON)
	}
	if value, ok := uu.mutation.Language(); ok {
		_spec.SetField(user.FieldLanguage, field.TypeString, value)
	}
	if value, ok := uu.mutation.Theme(); ok {
		_spec.SetField(user.FieldTheme, field.TypeString, value)
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if uu.mutation.FirstNameCleared() {
		_spec.ClearField(user.FieldFirstName, field.TypeString)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if uu.mutation.LastNameCleared() {
		_spec.ClearField(user.FieldLastName, field.TypeString)
	}
	if value, ok := uu.mutation.Sessions(); ok {
		_spec.SetField(user.FieldSessions, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedSessions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldSessions, value)
		})
	}
	if uu.mutation.SessionsCleared() {
		_spec.ClearField(user.FieldSessions, field.TypeJSON)
	}
	if uu.mutation.RoomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RoomsTable,
			Columns: user.RoomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedRoomsIDs(); len(nodes) > 0 && !uu.mutation.RoomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RoomsTable,
			Columns: user.RoomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RoomsTable,
			Columns: user.RoomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdateTime sets the "update_time" field.
func (uuo *UserUpdateOne) SetUpdateTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdateTime(t)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetIsEmailVerified sets the "is_email_verified" field.
func (uuo *UserUpdateOne) SetIsEmailVerified(b bool) *UserUpdateOne {
	uuo.mutation.SetIsEmailVerified(b)
	return uuo
}

// SetNillableIsEmailVerified sets the "is_email_verified" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsEmailVerified(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsEmailVerified(*b)
	}
	return uuo
}

// SetPasswordHash sets the "password_hash" field.
func (uuo *UserUpdateOne) SetPasswordHash(b []byte) *UserUpdateOne {
	uuo.mutation.SetPasswordHash(b)
	return uuo
}

// ClearPasswordHash clears the value of the "password_hash" field.
func (uuo *UserUpdateOne) ClearPasswordHash() *UserUpdateOne {
	uuo.mutation.ClearPasswordHash()
	return uuo
}

// SetBiography sets the "biography" field.
func (uuo *UserUpdateOne) SetBiography(s string) *UserUpdateOne {
	uuo.mutation.SetBiography(s)
	return uuo
}

// SetNillableBiography sets the "biography" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBiography(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetBiography(*s)
	}
	return uuo
}

// ClearBiography clears the value of the "biography" field.
func (uuo *UserUpdateOne) ClearBiography() *UserUpdateOne {
	uuo.mutation.ClearBiography()
	return uuo
}

// SetRole sets the "role" field.
func (uuo *UserUpdateOne) SetRole(s string) *UserUpdateOne {
	uuo.mutation.SetRole(s)
	return uuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetRole(*s)
	}
	return uuo
}

// SetFriendsIds sets the "friends_ids" field.
func (uuo *UserUpdateOne) SetFriendsIds(s []string) *UserUpdateOne {
	uuo.mutation.SetFriendsIds(s)
	return uuo
}

// AppendFriendsIds appends s to the "friends_ids" field.
func (uuo *UserUpdateOne) AppendFriendsIds(s []string) *UserUpdateOne {
	uuo.mutation.AppendFriendsIds(s)
	return uuo
}

// ClearFriendsIds clears the value of the "friends_ids" field.
func (uuo *UserUpdateOne) ClearFriendsIds() *UserUpdateOne {
	uuo.mutation.ClearFriendsIds()
	return uuo
}

// SetLanguage sets the "language" field.
func (uuo *UserUpdateOne) SetLanguage(s string) *UserUpdateOne {
	uuo.mutation.SetLanguage(s)
	return uuo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLanguage(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLanguage(*s)
	}
	return uuo
}

// SetTheme sets the "theme" field.
func (uuo *UserUpdateOne) SetTheme(s string) *UserUpdateOne {
	uuo.mutation.SetTheme(s)
	return uuo
}

// SetNillableTheme sets the "theme" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTheme(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTheme(*s)
	}
	return uuo
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFirstName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFirstName(*s)
	}
	return uuo
}

// ClearFirstName clears the value of the "first_name" field.
func (uuo *UserUpdateOne) ClearFirstName() *UserUpdateOne {
	uuo.mutation.ClearFirstName()
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLastName(*s)
	}
	return uuo
}

// ClearLastName clears the value of the "last_name" field.
func (uuo *UserUpdateOne) ClearLastName() *UserUpdateOne {
	uuo.mutation.ClearLastName()
	return uuo
}

// SetSessions sets the "sessions" field.
func (uuo *UserUpdateOne) SetSessions(s []string) *UserUpdateOne {
	uuo.mutation.SetSessions(s)
	return uuo
}

// AppendSessions appends s to the "sessions" field.
func (uuo *UserUpdateOne) AppendSessions(s []string) *UserUpdateOne {
	uuo.mutation.AppendSessions(s)
	return uuo
}

// ClearSessions clears the value of the "sessions" field.
func (uuo *UserUpdateOne) ClearSessions() *UserUpdateOne {
	uuo.mutation.ClearSessions()
	return uuo
}

// AddRoomIDs adds the "rooms" edge to the Room entity by IDs.
func (uuo *UserUpdateOne) AddRoomIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddRoomIDs(ids...)
	return uuo
}

// AddRooms adds the "rooms" edges to the Room entity.
func (uuo *UserUpdateOne) AddRooms(r ...*Room) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddRoomIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearRooms clears all "rooms" edges to the Room entity.
func (uuo *UserUpdateOne) ClearRooms() *UserUpdateOne {
	uuo.mutation.ClearRooms()
	return uuo
}

// RemoveRoomIDs removes the "rooms" edge to Room entities by IDs.
func (uuo *UserUpdateOne) RemoveRoomIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveRoomIDs(ids...)
	return uuo
}

// RemoveRooms removes "rooms" edges to Room entities.
func (uuo *UserUpdateOne) RemoveRooms(r ...*Room) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveRoomIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdateTime(); !ok {
		v := user.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Biography(); ok {
		if err := user.BiographyValidator(v); err != nil {
			return &ValidationError{Name: "biography", err: fmt.Errorf(`ent: validator failed for field "User.biography": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.IsEmailVerified(); ok {
		_spec.SetField(user.FieldIsEmailVerified, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeBytes, value)
	}
	if uuo.mutation.PasswordHashCleared() {
		_spec.ClearField(user.FieldPasswordHash, field.TypeBytes)
	}
	if value, ok := uuo.mutation.Biography(); ok {
		_spec.SetField(user.FieldBiography, field.TypeString, value)
	}
	if uuo.mutation.BiographyCleared() {
		_spec.ClearField(user.FieldBiography, field.TypeString)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeString, value)
	}
	if value, ok := uuo.mutation.FriendsIds(); ok {
		_spec.SetField(user.FieldFriendsIds, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedFriendsIds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldFriendsIds, value)
		})
	}
	if uuo.mutation.FriendsIdsCleared() {
		_spec.ClearField(user.FieldFriendsIds, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Language(); ok {
		_spec.SetField(user.FieldLanguage, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Theme(); ok {
		_spec.SetField(user.FieldTheme, field.TypeString, value)
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if uuo.mutation.FirstNameCleared() {
		_spec.ClearField(user.FieldFirstName, field.TypeString)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if uuo.mutation.LastNameCleared() {
		_spec.ClearField(user.FieldLastName, field.TypeString)
	}
	if value, ok := uuo.mutation.Sessions(); ok {
		_spec.SetField(user.FieldSessions, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedSessions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldSessions, value)
		})
	}
	if uuo.mutation.SessionsCleared() {
		_spec.ClearField(user.FieldSessions, field.TypeJSON)
	}
	if uuo.mutation.RoomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RoomsTable,
			Columns: user.RoomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedRoomsIDs(); len(nodes) > 0 && !uuo.mutation.RoomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RoomsTable,
			Columns: user.RoomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RoomsTable,
			Columns: user.RoomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
