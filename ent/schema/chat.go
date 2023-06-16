package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	loc "github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/ent/hook"
	"golang.org/x/crypto/bcrypt"
)

// Chat holds the schema definition for the Chat entity.
type Chat struct {
	ent.Schema
}

// Fields of the Chat.
func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.String("message").MinLen(1).MaxLen(1024),
	}
}

// Edges of the Chat.
func (Chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Ref("chat").
			Required(),

		edge.To("user", User.Type).Required(),
	}
}

func (Chat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Chat) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(formatMessage,
			hook.And(
				hook.HasFields("message"),
				hook.HasOp(ent.OpUpdate|ent.OpUpdateOne|ent.OpCreate),
			),
		),
	}
}

func formatMessage(next ent.Mutator) ent.Mutator {
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
