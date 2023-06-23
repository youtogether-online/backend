package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	loc "github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/ent/hook"
	"golang.org/x/crypto/bcrypt"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{

		field.String("title").MinLen(3).MaxLen(32).
			Annotations(entsql.DefaultExpr("'room' || setval(pg_get_serial_sequence('rooms','id'),nextval(pg_get_serial_sequence('rooms','id'))-1)")).
			StructTag(`json:"title,omitempty" validate:"omitempty,gte=3,lte=32"`),

		field.String("privacy").Default("friends").
			StructTag(`json:"privacy,omitempty" validate:"omitempty,enum=public*private*friends"`),

		field.Bytes("password_hash").Optional().Sensitive().Nillable(),

		field.String("description").Optional().MaxLen(140).Nillable().
			StructTag(`json:"description,omitempty" validate:"omitempty,lte=140"`),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("room").Unique().Immutable().Required(),
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
	}
}

func bcryptRoomPassword(next ent.Mutator) ent.Mutator {
	return hook.RoomFunc(func(ctx context.Context, m *loc.RoomMutation) (ent.Value, error) {
		password, _ := m.PasswordHash()

		hashedPassword, err := bcrypt.GenerateFromPassword(password, 12)
		if err != nil {
			return nil, err
		}

		m.SetPasswordHash(hashedPassword)

		return next.Mutate(ctx, m)
	})
}
