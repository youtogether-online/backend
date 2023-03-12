package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/wtkeqrf0/you_together/pkg/conf"
	"regexp"
)

// User holds the schema definition for the User dto.
type User struct {
	ent.Schema
}

var cfg = conf.GetConfig().Regexp

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").Optional().Match(regexp.MustCompile(cfg.UserName)).Unique(),
		field.String("email").Unique().Immutable().Match(regexp.MustCompile(cfg.Email)),
		field.Bytes("password_hash").Optional().Sensitive().Nillable(),
		field.Text("biography").Optional().MaxLen(140),
		field.Enum("role").Values("USER", "ADMIN").Default("USER"),
		field.String("avatar").Optional().Nillable(),
		field.Ints("friends_ids").Optional(),
		field.Enum("language").Values("EN", "RU").Default("EN"),
		field.Enum("theme").Values("WHITE", "DARK", "SYSTEM").Default("SYSTEM"),
		field.String("name").Optional().Match(regexp.MustCompile(cfg.Name)).Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
