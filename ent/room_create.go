// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wtkeqrf0/you-together/ent/room"
	"github.com/wtkeqrf0/you-together/ent/user"
)

// RoomCreate is the builder for creating a Room entity.
type RoomCreate struct {
	config
	mutation *RoomMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (rc *RoomCreate) SetCreateTime(t time.Time) *RoomCreate {
	rc.mutation.SetCreateTime(t)
	return rc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rc *RoomCreate) SetNillableCreateTime(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetCreateTime(*t)
	}
	return rc
}

// SetUpdateTime sets the "update_time" field.
func (rc *RoomCreate) SetUpdateTime(t time.Time) *RoomCreate {
	rc.mutation.SetUpdateTime(t)
	return rc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rc *RoomCreate) SetNillableUpdateTime(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetUpdateTime(*t)
	}
	return rc
}

// SetName sets the "name" field.
func (rc *RoomCreate) SetName(s string) *RoomCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetCustomName sets the "custom_name" field.
func (rc *RoomCreate) SetCustomName(s string) *RoomCreate {
	rc.mutation.SetCustomName(s)
	return rc
}

// SetNillableCustomName sets the "custom_name" field if the given value is not nil.
func (rc *RoomCreate) SetNillableCustomName(s *string) *RoomCreate {
	if s != nil {
		rc.SetCustomName(*s)
	}
	return rc
}

// SetOwnerID sets the "owner_id" field.
func (rc *RoomCreate) SetOwnerID(i int) *RoomCreate {
	rc.mutation.SetOwnerID(i)
	return rc
}

// SetPrivacy sets the "privacy" field.
func (rc *RoomCreate) SetPrivacy(r room.Privacy) *RoomCreate {
	rc.mutation.SetPrivacy(r)
	return rc
}

// SetNillablePrivacy sets the "privacy" field if the given value is not nil.
func (rc *RoomCreate) SetNillablePrivacy(r *room.Privacy) *RoomCreate {
	if r != nil {
		rc.SetPrivacy(*r)
	}
	return rc
}

// SetPasswordHash sets the "password_hash" field.
func (rc *RoomCreate) SetPasswordHash(b []byte) *RoomCreate {
	rc.mutation.SetPasswordHash(b)
	return rc
}

// SetHasChat sets the "has_chat" field.
func (rc *RoomCreate) SetHasChat(b bool) *RoomCreate {
	rc.mutation.SetHasChat(b)
	return rc
}

// SetNillableHasChat sets the "has_chat" field if the given value is not nil.
func (rc *RoomCreate) SetNillableHasChat(b *bool) *RoomCreate {
	if b != nil {
		rc.SetHasChat(*b)
	}
	return rc
}

// SetDescription sets the "description" field.
func (rc *RoomCreate) SetDescription(s string) *RoomCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rc *RoomCreate) SetNillableDescription(s *string) *RoomCreate {
	if s != nil {
		rc.SetDescription(*s)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RoomCreate) SetID(i int) *RoomCreate {
	rc.mutation.SetID(i)
	return rc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (rc *RoomCreate) AddUserIDs(ids ...int) *RoomCreate {
	rc.mutation.AddUserIDs(ids...)
	return rc
}

// AddUsers adds the "users" edges to the User entity.
func (rc *RoomCreate) AddUsers(u ...*User) *RoomCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rc.AddUserIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (rc *RoomCreate) Mutation() *RoomMutation {
	return rc.mutation
}

// Save creates the Room in the database.
func (rc *RoomCreate) Save(ctx context.Context) (*Room, error) {
	if err := rc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Room, RoomMutation](ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoomCreate) SaveX(ctx context.Context) *Room {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoomCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoomCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoomCreate) defaults() error {
	if _, ok := rc.mutation.CreateTime(); !ok {
		if room.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized room.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := room.DefaultCreateTime()
		rc.mutation.SetCreateTime(v)
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		if room.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized room.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := room.DefaultUpdateTime()
		rc.mutation.SetUpdateTime(v)
	}
	if _, ok := rc.mutation.Privacy(); !ok {
		v := room.DefaultPrivacy
		rc.mutation.SetPrivacy(v)
	}
	if _, ok := rc.mutation.HasChat(); !ok {
		v := room.DefaultHasChat
		rc.mutation.SetHasChat(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoomCreate) check() error {
	if _, ok := rc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Room.create_time"`)}
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Room.update_time"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := room.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Room.name": %w`, err)}
		}
	}
	if v, ok := rc.mutation.CustomName(); ok {
		if err := room.CustomNameValidator(v); err != nil {
			return &ValidationError{Name: "custom_name", err: fmt.Errorf(`ent: validator failed for field "Room.custom_name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "Room.owner_id"`)}
	}
	if v, ok := rc.mutation.OwnerID(); ok {
		if err := room.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`ent: validator failed for field "Room.owner_id": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Privacy(); !ok {
		return &ValidationError{Name: "privacy", err: errors.New(`ent: missing required field "Room.privacy"`)}
	}
	if v, ok := rc.mutation.Privacy(); ok {
		if err := room.PrivacyValidator(v); err != nil {
			return &ValidationError{Name: "privacy", err: fmt.Errorf(`ent: validator failed for field "Room.privacy": %w`, err)}
		}
	}
	if _, ok := rc.mutation.HasChat(); !ok {
		return &ValidationError{Name: "has_chat", err: errors.New(`ent: missing required field "Room.has_chat"`)}
	}
	if v, ok := rc.mutation.Description(); ok {
		if err := room.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Room.description": %w`, err)}
		}
	}
	return nil
}

func (rc *RoomCreate) sqlSave(ctx context.Context) (*Room, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoomCreate) createSpec() (*Room, *sqlgraph.CreateSpec) {
	var (
		_node = &Room{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(room.Table, sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreateTime(); ok {
		_spec.SetField(room.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := rc.mutation.UpdateTime(); ok {
		_spec.SetField(room.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(room.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.CustomName(); ok {
		_spec.SetField(room.FieldCustomName, field.TypeString, value)
		_node.CustomName = &value
	}
	if value, ok := rc.mutation.OwnerID(); ok {
		_spec.SetField(room.FieldOwnerID, field.TypeInt, value)
		_node.OwnerID = value
	}
	if value, ok := rc.mutation.Privacy(); ok {
		_spec.SetField(room.FieldPrivacy, field.TypeEnum, value)
		_node.Privacy = value
	}
	if value, ok := rc.mutation.PasswordHash(); ok {
		_spec.SetField(room.FieldPasswordHash, field.TypeBytes, value)
		_node.PasswordHash = &value
	}
	if value, ok := rc.mutation.HasChat(); ok {
		_spec.SetField(room.FieldHasChat, field.TypeBool, value)
		_node.HasChat = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(room.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if nodes := rc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   room.UsersTable,
			Columns: room.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoomCreateBulk is the builder for creating many Room entities in bulk.
type RoomCreateBulk struct {
	config
	builders []*RoomCreate
}

// Save creates the Room entities in the database.
func (rcb *RoomCreateBulk) Save(ctx context.Context) ([]*Room, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Room, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoomMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoomCreateBulk) SaveX(ctx context.Context) []*Room {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoomCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoomCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
