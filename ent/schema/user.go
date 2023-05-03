package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"fmt"
	gen "github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/ent/hook"
	"github.com/wtkeqrf0/you-together/ent/room"
	"github.com/wtkeqrf0/you-together/pkg/bind"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// User holds the schema definition for the User dto.
type User struct {
	ent.Schema
}

const emailRegexp string = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Match(bind.NameRegexp).Annotations(
			entsql.DefaultExpr("'user' || currval(pg_get_serial_sequence('users','id'))")),
		field.String("email").Unique().Match(regexp.MustCompile(emailRegexp)),
		field.Bool("is_email_verified").Default(false),
		field.Bytes("password_hash").Optional().Sensitive().Nillable(),
		field.Text("biography").Optional().MaxLen(512).Nillable(),
		field.String("role").Default("USER"),
		field.Strings("friends_ids").Optional(),
		field.String("language").Default("EN"),
		field.String("theme").Default("SYSTEM"),
		field.String("first_name").Optional().MinLen(2).MaxLen(30).Nillable(),
		field.String("last_name").Optional().MinLen(2).MaxLen(30).Nillable(),
		field.Strings("sessions").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rooms", Room.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(bcryptUserPassword,
			hook.And(
				hook.HasFields("password_hash"),
				hook.HasOp(ent.OpUpdate|ent.OpUpdateOne|ent.OpCreate),
			),
		),
		hook.If(roomNameCheck,
			hook.And(
				hook.HasFields("name"),
				hook.HasOp(ent.OpUpdateOne|ent.OpUpdate|ent.OpCreate),
			),
		),
	}
}

func bcryptUserPassword(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
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

func roomNameCheck(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
		username, ok := m.Name()
		if !ok {
			return nil, fmt.Errorf("username is not set")
		}

		if exist, err := m.Client().Room.Query().Where(room.Name(username)).Exist(ctx); err != nil {
			return nil, err

		} else if exist {
			return nil, fmt.Errorf("name already exist")
		}
		return next.Mutate(ctx, m)
	})
}
