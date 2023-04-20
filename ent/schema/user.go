package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"math/rand"
	"regexp"
)

// User holds the schema definition for the User dto.
type User struct {
	ent.Schema
}

var (
	cfg     = conf.GetConfig().Regexp
	idRunes = []rune("abcdefghijklmnopqrstuvwxyz")
)

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Match(regexp.MustCompile(cfg.Username)).DefaultFunc(func() string {
			b := make([]rune, 6)
			for i := range b {
				b[i] = idRunes[rand.Intn(len(idRunes))]
			}
			return string(b)
		}),
		field.String("email").Unique().Match(regexp.MustCompile(cfg.Email)),
		field.Bool("is_email_verified").Default(false),
		field.Bytes("password_hash").Optional().Sensitive().Nillable(),
		field.Text("biography").Optional().MaxLen(512).Nillable(),
		field.String("role").Default("USER"),
		field.Strings("friends_ids").Optional(),
		field.String("language").Default("EN"),
		field.String("theme").Default("SYSTEM"),
		field.String("first_name").Optional().Match(regexp.MustCompile(cfg.Name)).MinLen(1).MaxLen(30).Nillable(),
		field.String("last_name").Optional().Match(regexp.MustCompile(cfg.Name)).MinLen(1).MaxLen(30).Nillable(),
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
