package repo

import (
	"context"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/ent/user"
	"golang.org/x/crypto/bcrypt"
)

// UserExists return true if user Exists. Panic if error occurred
func (r UserStorage) UserExists(ctx context.Context, id int) bool {
	return r.userClient.Query().Where(user.ID(id)).ExistX(ctx)
}

// CreateUser with or without password and return it (only auth)
func (r UserStorage) CreateUser(ctx context.Context, email, password string) (*ent.User, error) {
	res := r.userClient.Create().SetEmail(email)
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			return nil, err
		}
		res.SetPasswordHash(hashedPassword)
	}
	return res.Save(ctx)
}

// AuthUserByEmail return the user with given email (only auth)
func (r UserStorage) AuthUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.userClient.Query().Where(
		user.EmailEQ(email),
	).Only(ctx)
}
