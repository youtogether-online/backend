package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"regexp"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Match(regexp.MustCompile(cfg.Username)),
		//field.String("admin").Unique().Immutable().Match(regexp.MustCompile(cfg.Email)),
		field.Enum("privacy").Values("PRIVATE", "FRIENDS", "PUBLIC").Default("PUBLIC"),
		field.String("password_hash").Optional().Sensitive(),
		field.Bool("has_chat").Default(true),
		field.String("description").Optional().MaxLen(140),
		field.String("avatar").Optional(),
		//TODO time accurately syncing
	}
}

// Edges of the Room
func (Room) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Room) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
