package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthPostgres interface {
	UserExists(ctx context.Context, username string) bool
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

// UserExists return true if user Exists. Panic if error occurred
func (a AuthService) UserExists(username string) bool {
	return a.postgres.UserExists(context.Background(), username)
}

// CreateUserWithPassword with or without password and return it (only auth)
func (a AuthService) CreateUserWithPassword(email, password string) ([]byte, string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, "", err
	}
	id, err := a.postgres.CreateUserWithPassword(context.Background(), email, hashedPassword)
	return hashedPassword, id, err
}

// CreateUserByEmail without password and return it (only auth)
func (a AuthService) CreateUserByEmail(email string) (string, error) {
	return a.postgres.CreateUserByEmail(context.Background(), email)
}

// AuthUserByEmail return the user with given email (only auth)
func (a AuthService) AuthUserByEmail(email string) ([]byte, string, error) {
	return a.postgres.AuthUserByEmail(context.Background(), email)
}

func (a AuthService) AuthUserWithInfo(email string) (bool, string, error) {
	return a.postgres.AuthUserWithInfo(context.Background(), email)
}

func (a AuthService) SetEmailVerified(email string) error {
	return a.postgres.SetEmailVerified(context.Background(), email)
}

type AuthRedis interface {
	SetSession(ctx context.Context, sessionId string, info map[string]string) error
	GetSession(ctx context.Context, sessionId string) (map[string]string, error)
	DelSession(ctx context.Context, sessionId string) error
	EqualsPopCode(ctx context.Context, email string, code string) bool
	SetCodes(ctx context.Context, key string, value ...any) error
	GetTTL(ctx context.Context, sessionId string) time.Duration
}

func (a AuthService) GetSession(sessionId string) (map[string]string, error) {
	return a.redis.GetSession(context.Background(), sessionId)
}

func (a AuthService) SetSession(sessionId string, info map[string]string) error {
	return a.redis.SetSession(context.Background(), sessionId, info)
}

func (a AuthService) DelSession(sessionId string) error {
	return a.redis.DelSession(context.Background(), sessionId)
}

func (a AuthService) EqualsPopCode(email string, code string) bool {
	return a.redis.EqualsPopCode(context.Background(), email, code)
}

func (a AuthService) SetCodes(key string, value ...any) error {
	return a.redis.SetCodes(context.Background(), key, value...)
}

func (a AuthService) GetTTL(sessionId string) time.Duration {
	return a.redis.GetTTL(context.Background(), sessionId)
}
