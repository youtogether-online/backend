package authorization

import (
	"context"
	"github.com/wtkeqrf0/you_together/ent"
)

type AuthRepo interface {
	UserExists(ctx context.Context, id int) bool
	CreateUser(ctx context.Context, mail, password string) (*ent.User, error)
	AuthUserByEmail(ctx context.Context, mail string) (*ent.User, error)
}

type Auth struct {
	auth AuthRepo
}

func NewAuth(auth AuthRepo) *Auth {
	return &Auth{auth: auth}
}

func (a Auth) UserExists(id int) bool {
	return a.auth.UserExists(context.Background(), id)
}

func (a Auth) CreateUser(email, password string) (*ent.User, error) {
	return a.auth.CreateUser(context.Background(), email, password)
}

func (a Auth) AuthUserByEmail(email string) (*ent.User, error) {
	return a.auth.AuthUserByEmail(context.Background(), email)
}
