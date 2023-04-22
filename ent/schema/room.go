package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"math/rand"
	"regexp"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

const (
	name  string = "."
	email string = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Match(regexp.MustCompile(name)).DefaultFunc(func() string {
			l := 6 + rand.Intn(4)
			b := make([]rune, l)

			b[0] = NameRunes[rand.Intn(52)]
			for i := 1; i < l; i++ {
				b[i] = NameRunes[rand.Intn(l)]
			}
			return string(b)
		}).MinLen(5).MaxLen(20),
		field.String("custom_name").Optional().MinLen(2).MaxLen(20).Nillable(),
		field.Int("owner_id").Unique().Positive(),
		field.Enum("privacy").Values("PRIVATE", "FRIENDS", "PUBLIC").Default("PUBLIC"),
		field.String("password_hash").Optional().Sensitive().Nillable(),
		field.Bool("has_chat").Default(true),
		field.String("description").Optional().MaxLen(140).Nillable(),
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
