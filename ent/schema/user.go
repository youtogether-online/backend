package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/wtkeqrf0/you_together/pkg/conf"
	"github.com/zhexuany/wordGenerator"
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
		field.String("id").StorageKey("username").Unique().Match(regexp.MustCompile(cfg.Username)).DefaultFunc(func() string {
			/*b := make([]rune, 6)
			for i := range b {
				b[i] = idRunes[rand.Intn(len(idRunes))]
			}
			return string(b)
			*/
			return wordGenerator.GetWord(rand.Intn(5) + 5)
		}).MinLen(5).MaxLen(20),
		field.String("email").Unique().Match(regexp.MustCompile(cfg.Email)),
		field.Bool("is_email_verified").Default(false),
		field.Bytes("password_hash").Optional().Sensitive(),
		field.Text("biography").Optional().MaxLen(512),
		field.Enum("role").Values("USER", "ADMIN").Default("USER"),
		field.Strings("friends_ids").Optional(),
		field.Enum("language").Values("EN", "RU").Default("EN"),
		field.Enum("theme").Values("WHITE", "DARK", "SYSTEM").Default("SYSTEM"),
		field.String("first_name").Optional().Match(regexp.MustCompile(cfg.Name)).MinLen(4).MaxLen(20),
		field.String("last_name").Optional().Match(regexp.MustCompile(cfg.Name)).MinLen(4).MaxLen(20),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rooms", Room.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
