package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	loc "github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/pkg/middleware/bind"
	"golang.org/x/crypto/bcrypt"

	"github.com/wtkeqrf0/you-together/ent/hook"
)

// User holds the schema definition for the User dto.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Match(bind.NameRegexp).Annotations(
			entsql.DefaultExpr("'user' || setval(pg_get_serial_sequence('users','id'),nextval(pg_get_serial_sequence('users','id'))-1)")).
			StructTag(`json:"name,omitempty" validate:"omitempty,gte=5,lte=20,name"`),

		field.String("email").Unique().Match(bind.EmailRegexp).
			StructTag(`json:"email,omitempty" validate:"required,email"`),

		field.Bool("is_email_verified").Default(false).
			StructTag(`json:"isEmailVerified,omitempty"`),

		field.Bytes("password_hash").Optional().Sensitive().Nillable(),

		field.Text("biography").Optional().MaxLen(512).Nillable().
			StructTag(`json:"biography,omitempty" validate:"omitempty,lte=512"`),

		field.String("role").Default("USER").
			StructTag(`json:"role,omitempty" validate:"omitempty,enum=USER*ADMIN"`),

		field.Strings("friends_ids").Optional().
			StructTag(`json:"friendsIds,omitempty"`),

		field.String("language").Default("EN").
			StructTag(`json:"language,omitempty" validate:"omitempty,enum=EN*RU"`),

		field.String("theme").Default("SYSTEM").
			StructTag(`json:"theme,omitempty" validate:"omitempty,enum=SYSTEM*WHITE*DARK"`),

		field.String("first_name").Optional().MinLen(3).MaxLen(32).Nillable().
			StructTag(`json:"firstName,omitempty" validate:"omitempty,gte=3,lte=32"`),

		field.String("last_name").Optional().MinLen(3).MaxLen(32).Nillable().
			StructTag(`json:"lastName,omitempty" validate:"omitempty,gte=3,lte=32"`),

		field.Strings("sessions").Optional().StructTag(`json:"-"`),
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
	}
}

func bcryptUserPassword(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *loc.UserMutation) (ent.Value, error) {
		password, _ := m.PasswordHash()

		hashedPassword, err := bcrypt.GenerateFromPassword(password, 12)
		if err != nil {
			return nil, err
		}

		m.SetPasswordHash(hashedPassword)

		return next.Mutate(ctx, m)
	})
}
