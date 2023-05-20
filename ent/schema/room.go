package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"fmt"
	loc "github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/ent/hook"
	"github.com/wtkeqrf0/you-together/ent/room"
	"github.com/wtkeqrf0/you-together/pkg/bind"
	"golang.org/x/crypto/bcrypt"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StructTag(`json:"-"`).Annotations(
			entsql.DefaultExpr("nextval(pg_get_serial_sequence('users', 'id'))")),

		field.String("name").Unique().Match(bind.NameRegexp).Annotations(
			entsql.DefaultExpr("'room' || currval(pg_get_serial_sequence('users', 'id'))")).
			StructTag(`json:"name,omitempty"`),

		field.String("custom_name").Optional().MinLen(2).MaxLen(20).Nillable().
			StructTag(`json:"customName,omitempty"`),

		field.Int("owner_id").Unique().Positive().StructTag(`json:"-"`),

		field.Enum("privacy").Values("PRIVATE", "FRIENDS", "PUBLIC").
			Default("PUBLIC").StructTag(`json:"privacy,omitempty"`),

		field.Bytes("password_hash").Optional().Sensitive().Nillable(),

		field.Bool("has_chat").Default(true).
			StructTag(`json:"has_chat,omitempty"`),

		field.String("description").Optional().MaxLen(140).Nillable().
			StructTag(`json:"description,omitempty"`),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("rooms"),
	}
}

func (Room) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Room) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(bcryptRoomPassword,
			hook.And(
				hook.HasFields("password_hash"),
				hook.HasOp(ent.OpUpdate|ent.OpUpdateOne|ent.OpCreate),
			),
		),

		hook.If(userNameCheck,
			hook.And(
				hook.HasFields("name"),
				hook.HasOp(ent.OpUpdateOne|ent.OpUpdate|ent.OpCreate),
			),
		),
	}
}

func bcryptRoomPassword(next ent.Mutator) ent.Mutator {
	return hook.RoomFunc(func(ctx context.Context, m *loc.RoomMutation) (ent.Value, error) {
		password, ok := m.PasswordHash()
		if !ok {
			return nil, fmt.Errorf("password_hash is not set")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(password, 12)
		if err != nil {
			return nil, err
		}

		m.SetPasswordHash(hashedPassword)

		return next.Mutate(ctx, m)
	})
}

func userNameCheck(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *loc.UserMutation) (ent.Value, error) {
		username, ok := m.Name()
		if !ok {
			return nil, fmt.Errorf("roomname is not set")
		}

		if exist, err := m.Client().Room.Query().Where(room.Name(username)).Exist(ctx); err != nil {
			return nil, err

		} else if exist {
			return nil, fmt.Errorf("name already exist")
		}
		return next.Mutate(ctx, m)
	})
}
