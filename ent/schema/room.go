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

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Match(regexp.MustCompile(cfg.Username)).DefaultFunc(func() string {
			b := make([]rune, 6)
			for i := range b {
				b[i] = idRunes[rand.Intn(len(idRunes))]
			}
			return string(b)
		}).MinLen(5).MaxLen(20),
		field.String("custom_name").Optional().MinLen(3).MaxLen(20).Nillable(),
		field.String("owner").Unique().Immutable().Match(regexp.MustCompile(cfg.Email)),
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
