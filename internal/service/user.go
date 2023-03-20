package service

import (
	"context"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"time"
)

type UserPostgres interface {
	FindMe(ctx context.Context, username string) (dto.MyUserDTO, error)
	FindUserByUsername(ctx context.Context, username string) (dto.UserDTO, error)
	FindAllUsers(ctx context.Context, limit int) ([]*ent.User, error)
	UpdateUser(ctx context.Context, user *ent.User) error
	DeleteUser(ctx context.Context, id int) error
}

type UserService struct {
	postgres UserPostgres
	redis    UserRedis
}

func NewUserService(postgres UserPostgres, redis UserRedis) *UserService {
	return &UserService{postgres: postgres, redis: redis}
}

// FindUserByUsername returns the main information about user
func (u UserService) FindUserByUsername(username string) (dto.UserDTO, error) {
	return u.postgres.FindUserByUsername(context.Background(), username)
}

// FindMe returns the detail information about user
func (u UserService) FindMe(username string) (dto.MyUserDTO, error) {
	user, err := u.postgres.FindMe(context.Background(), username)
	if err != nil {
		return dto.MyUserDTO{}, err
	}
	user.CutEmail()
	return user, nil
}

func (u UserService) FindAllUsers(limit int) ([]dto.UserDTO, error) {
	return nil, nil
}

func (u UserService) UpdateUser(user ent.User) error {
	return nil
}

func (u UserService) DeleteUser(id int) error {
	return nil
}

type UserRedis interface {
	ContainsKeys(ctx context.Context, keys ...string) (int64, error)
	SetVariable(ctx context.Context, key string, value any, exp time.Duration) error
}

// ContainsKeys of redis by key
func (u UserService) ContainsKeys(keys ...string) (int64, error) {
	return u.redis.ContainsKeys(context.Background(), keys...)
}

// SetVariable of redis by key, his value and exploration time
func (u UserService) SetVariable(key string, value any, exp time.Duration) error {
	return u.redis.SetVariable(context.Background(), key, value, exp)
}
