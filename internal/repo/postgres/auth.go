package postgres

import (
	"context"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/ent/user"
)

// IDExist returns true if username exists. Panics if error occurred
func (r *UserStorage) IDExist(ctx context.Context, id int) (bool, error) {
	return r.userClient.Query().Where(user.ID(id)).Exist(ctx)
}

// UserExistsByEmail returns true if user Exists. Panic if error occurred
func (r *UserStorage) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	return r.userClient.Query().Where(user.Email(email)).Exist(ctx)
}

// CreateUserWithPassword without verified email and returns it (only on registration)
func (r *UserStorage) CreateUserWithPassword(ctx context.Context, email string, password []byte, language *string) (*ent.User, error) {
	return r.userClient.Create().SetEmail(email).SetNillableLanguage(language).
		SetPasswordHash(password).Save(ctx)
}

// CreateUserByEmail without password and returns it (only on registration)
func (r *UserStorage) CreateUserByEmail(ctx context.Context, email string, language *string) (*ent.User, error) {
	return r.userClient.Create().SetEmail(email).SetNillableLanguage(language).
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
