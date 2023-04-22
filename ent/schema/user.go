package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"math/rand"
	"regexp"
)

// User holds the schema definition for the User dto.
type User struct {
	ent.Schema
}

var (
	NameRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Match(regexp.MustCompile(name)).DefaultFunc(func() string {
			l := 6 + rand.Intn(4)
			b := make([]rune, l)

			b[0] = NameRunes[rand.Intn(52)]
			for i := 1; i < l; i++ {
				b[i] = NameRunes[rand.Intn(l)]
			}
			return string(b)
		}),
		field.String("email").Unique().Match(regexp.MustCompile(email)),
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
