package service

import (
	"context"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
)

type AuthPostgres interface {
	IDExist(ctx context.Context, id int) (bool, error)
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	CreateUserWithPassword(ctx context.Context, email string, password []byte, language *string) (*ent.User, error)
	CreateUserByEmail(ctx context.Context, email string, language *string) (*ent.User, error)
	AuthUserByEmail(ctx context.Context, email string) (*ent.User, error)
	SetEmailVerified(ctx context.Context, email string) error
	AddSession(ctx context.Context, id int, sessions ...string) error
}

type AuthService struct {
	postgres AuthPostgres
	redis    AuthRedis
}

func NewAuthService(postgres AuthPostgres, redis AuthRedis) *AuthService {
	return &AuthService{postgres: postgres, redis: redis}
}

// IDExist returns true if user Exists. Panics if error occurred
func (a *AuthService) IDExist(id int) (bool, error) {
	return a.postgres.IDExist(context.Background(), id)
}

// UserExistsByEmail returns true if user Exists. Panic if error occurred
func (a *AuthService) UserExistsByEmail(email string) (bool, error) {
	return a.postgres.UserExistsByEmail(context.Background(), email)
}

// CreateUserWithPassword without verified email and returns it (only on registration)
func (a *AuthService) CreateUserWithPassword(email string, password []byte, language *string) (*ent.User, error) {
	return a.postgres.CreateUserWithPassword(context.Background(), email, password, language)
}

// CreateUserByEmail without password and returns it (only on registration)
func (a *AuthService) CreateUserByEmail(email string, language *string) (*ent.User, error) {
	return a.postgres.CreateUserByEmail(context.Background(), email, language)
}

// AuthUserByEmail returns the user's password hash and username with given email (only on sessions)
func (a *AuthService) AuthUserByEmail(email string) (*ent.User, error) {
	return a.postgres.AuthUserByEmail(context.Background(), email)
}

// SetEmailVerified to true
func (a *AuthService) SetEmailVerified(email string) error {
	return a.postgres.SetEmailVerified(context.Background(), email)
}

func (a *AuthService) AddSession(id int, sessions ...string) error {
	return a.postgres.AddSession(context.Background(), id, sessions...)
}

type AuthRedis interface {
	SetSession(ctx context.Context, sessionId string, info dao.Session) error
	GetSession(ctx context.Context, sessionId string) (*dao.Session, error)
	ExpandExpireSession(ctx context.Context, sessionId string) (bool, error)
	DelKeys(ctx context.Context, keys ...string)
	EqualsPopCode(ctx context.Context, email string, code string) (bool, error)
	SetCodes(ctx context.Context, key string, value ...any) error
}

// GetSession and all its parameters
func (a *AuthService) GetSession(sessionId string) (*dao.Session, error) {
	return a.redis.GetSession(context.Background(), sessionId)
}

// SetSession and all its parameters
func (a *AuthService) SetSession(sessionId string, info dao.Session) error {
	return a.redis.SetSession(context.Background(), sessionId, info)
}

// ExpandExpireSession if key exists and have lesser than 15 days of expire
func (a *AuthService) ExpandExpireSession(sessionId string) (bool, error) {
	return a.redis.ExpandExpireSession(context.Background(), sessionId)
}

// DelKeys fully deletes session id
func (a *AuthService) DelKeys(keys ...string) {
	a.redis.DelKeys(context.Background(), keys...)
}

// EqualsPopCode returns true if code is involved in email and deletes it
func (a *AuthService) EqualsPopCode(email string, code string) (bool, error) {
	return a.redis.EqualsPopCode(context.Background(), email, code)
}

// SetCodes or add it to existing key
func (a *AuthService) SetCodes(key string, value ...any) error {
	return a.redis.SetCodes(context.Background(), key, value...)
}
