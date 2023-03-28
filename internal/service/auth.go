package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

type AuthPostgres interface {
	UserExists(ctx context.Context, username string) bool
	UserExistsByEmail(ctx context.Context, email string) bool
	CreateUserWithPassword(ctx context.Context, email string, hashedPassword []byte) (string, error)
	CreateUserByEmail(ctx context.Context, email string) (string, error)
	AuthUserByEmail(ctx context.Context, email string) ([]byte, string, error)
	AuthUserWithInfo(ctx context.Context, email string) (bool, string, error)
	SetEmailVerified(ctx context.Context, email string) error
}

type AuthService struct {
	postgres AuthPostgres
	redis    AuthRedis
}

func NewAuthService(postgres AuthPostgres, redis AuthRedis) *AuthService {
	return &AuthService{postgres: postgres, redis: redis}
}

// UserExists returns true if user Exists. Panics if error occurred
func (a AuthService) UserExists(username string) bool {
	return a.postgres.UserExists(context.Background(), username)
}

// UserExistsByEmail returns true if user Exists. Panic if error occurred
func (a AuthService) UserExistsByEmail(email string) bool {
	return a.postgres.UserExistsByEmail(context.Background(), email)
}

// CreateUserWithPassword without verified email and returns it (only on registration)
func (a AuthService) CreateUserWithPassword(email, password string) ([]byte, string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, "", err
	}
	id, err := a.postgres.CreateUserWithPassword(context.Background(), email, hashedPassword)
	return hashedPassword, id, err
}

// CreateUserByEmail without password and returns it (only on registration)
func (a AuthService) CreateUserByEmail(email string) (string, error) {
	return a.postgres.CreateUserByEmail(context.Background(), email)
}

// AuthUserByEmail returns the user's password hash and username with given email (only on authorization)
func (a AuthService) AuthUserByEmail(email string) ([]byte, string, error) {
	return a.postgres.AuthUserByEmail(context.Background(), email)
}

// AuthUserWithInfo returns the user's "Is email verified" and username with given email (only auth)
func (a AuthService) AuthUserWithInfo(email string) (bool, string, error) {
	return a.postgres.AuthUserWithInfo(context.Background(), email)
}

// SetEmailVerified to true
func (a AuthService) SetEmailVerified(email string) error {
	return a.postgres.SetEmailVerified(context.Background(), email)
}

type AuthRedis interface {
	SetSession(ctx context.Context, sessionId string, info map[string]string) error
	GetSession(ctx context.Context, sessionId string) (map[string]string, error)
	ExpandExpireSession(ctx context.Context, sessionId string) (bool, error)
	FindSessionsByUsername(ctx context.Context, userName string) []map[string]string
	DelKeys(ctx context.Context, keys ...string)
	EqualsPopCode(ctx context.Context, email string, code string) (bool, error)
	SetCodes(ctx context.Context, key string, value ...any) error
}

// GetSession and all its parameters
func (a AuthService) GetSession(sessionId string) (map[string]string, error) {
	return a.redis.GetSession(context.Background(), sessionId)
}

// SetSession and all its parameters
func (a AuthService) SetSession(sessionId string, info map[string]string) error {
	return a.redis.SetSession(context.Background(), sessionId, info)
}

// ExpandExpireSession if key exists and have lesser than 15 days of expire
func (a AuthService) ExpandExpireSession(sessionId string) (bool, error) {
	return a.redis.ExpandExpireSession(context.Background(), sessionId)
}

// FindSessionsByUsername returns all existing sessions by username
func (a AuthService) FindSessionsByUsername(userName string) []map[string]string {
	return a.redis.FindSessionsByUsername(context.Background(), userName)
}

// DelKeys fully deletes session id
func (a AuthService) DelKeys(keys ...string) {
	a.redis.DelKeys(context.Background(), keys...)
}

// EqualsPopCode returns true if code is involved in email and deletes it
func (a AuthService) EqualsPopCode(email string, code string) (bool, error) {
	return a.redis.EqualsPopCode(context.Background(), email, code)
}

// SetCodes or add it to existing key
func (a AuthService) SetCodes(key string, value ...any) error {
	return a.redis.SetCodes(context.Background(), key, value...)
}
