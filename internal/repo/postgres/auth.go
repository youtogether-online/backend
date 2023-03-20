package postgres

import (
	"context"
	"fmt"
	"github.com/wtkeqrf0/you_together/ent/user"
)

// UserExists return true if user Exists. Panic if error occurred
func (r *UserStorage) UserExists(ctx context.Context, username string) bool {
	return r.userClient.Query().Where(user.ID(username)).ExistX(ctx)
}

// CreateUserWithPassword and return it (only auth)
func (r *UserStorage) CreateUserWithPassword(ctx context.Context, email string, hashedPassword []byte) (string, error) {
	// TODO not fully returning user
	customer, err := r.userClient.Create().SetEmail(email).SetPasswordHash(hashedPassword).Save(ctx)
	return customer.ID, err
}

// CreateUserByEmail without password and return it (only auth)
func (r *UserStorage) CreateUserByEmail(ctx context.Context, email string) (string, error) {
	customer, err := r.userClient.Create().SetEmail(email).SetIsEmailVerified(true).Save(ctx)
	return customer.ID, err
}

// AuthUserByEmail returns the user's password hash and username with given email (only auth)
func (r *UserStorage) AuthUserByEmail(ctx context.Context, email string) ([]byte, string, error) {
	var res []struct {
		PasswordHash []byte `sql:"password_hash"`
		Username     string `sql:"username"`
	}
	err := r.userClient.Query().Where(
		user.EmailEQ(email),
	).Select(user.FieldPasswordHash).Scan(ctx, &res)
	if err != nil || len(res) != 1 {
		return nil, "", fmt.Errorf("cannot get user by email: %v", err)
	}
	return res[0].PasswordHash, res[0].Username, nil
}

// AuthUserWithInfo returns the user's "Is email verified" and username with given email (only auth)
func (r *UserStorage) AuthUserWithInfo(ctx context.Context, email string) (bool, string, error) {
	var res []struct {
		IsEmailVerified bool   `sql:"is_email_verified"`
		Username        string `sql:"username"`
	}
	err := r.userClient.Query().Where(
		user.EmailEQ(email),
	).Select(user.FieldPasswordHash).Scan(ctx, &res)
	if err != nil || len(res) != 1 {
		return false, "", fmt.Errorf("cannot get user by email: %v", err)
	}
	return res[0].IsEmailVerified, res[0].Username, nil
}

func (r *UserStorage) SetEmailVerified(ctx context.Context, email string) error {
	return r.userClient.Update().SetIsEmailVerified(true).Where(user.Email(email)).Exec(ctx)
}
