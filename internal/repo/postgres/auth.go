package postgres

import (
	"context"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/ent/user"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
)

// IDExist returns true if username exists. Panics if error occurred
func (r *UserStorage) IDExist(ctx context.Context, id int) bool {
	return r.userClient.Query().Where(user.ID(id)).ExistX(ctx)
}

// UserExistsByEmail returns true if user Exists. Panic if error occurred
func (r *UserStorage) UserExistsByEmail(ctx context.Context, email string) bool {
	return r.userClient.Query().Where(user.Email(email)).ExistX(ctx)
}

// CreateUserWithPassword without verified email and returns it (only on registration)
func (r *UserStorage) CreateUserWithPassword(ctx context.Context, auth dto.EmailWithPassword) (*ent.User, error) {
	return r.userClient.Create().SetEmail(auth.Email).
		SetTheme(auth.Theme).SetLanguage(auth.Language).
		SetPasswordHash([]byte(auth.Password)).Save(ctx)
}

// CreateUserByEmail without password and returns it (only on registration)
func (r *UserStorage) CreateUserByEmail(ctx context.Context, auth dto.EmailWithCode) (*ent.User, error) {
	return r.userClient.Create().SetEmail(auth.Email).
		SetNillableTheme(&auth.Theme).SetNillableLanguage(&auth.Language).
		SetIsEmailVerified(true).Save(ctx)
}

// AuthUserByEmail returns the user's password hash and username with given email (only on sessions)
func (r *UserStorage) AuthUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.userClient.Query().Where(
		user.EmailEQ(email),
	).Only(ctx)
}

// SetEmailVerified to true
func (r *UserStorage) SetEmailVerified(ctx context.Context, email string) error {
	return r.userClient.Update().SetIsEmailVerified(true).Where(user.Email(email)).Exec(ctx)
}

func (r *UserStorage) AddSession(ctx context.Context, id int, sessions ...string) error {
	return r.userClient.Update().AppendSessions(sessions).Where(user.ID(id)).Exec(ctx)
}
