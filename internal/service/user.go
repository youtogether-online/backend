package service

import (
	"context"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"strings"
	"time"
)

type UserPostgres interface {
	FindMe(ctx context.Context, id string) (dto.MyUserDTO, error)
	FindUserByUsername(ctx context.Context, username string) (dto.UserDTO, error)
	FindUserByID(ctx context.Context, id string) (*ent.User, error)
	UpdateUser(ctx context.Context, customer dto.UpdateUserDTO, username string) error
	UpdatePassword(ctx context.Context, password, username string) error
	UpdateEmail(ctx context.Context, email, username string) error
	UpdateUsername(ctx context.Context, newUsername, username string) error
	UsernameExist(ctx context.Context, username string) bool
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

func (u UserService) FindUserByID(id string) (*ent.User, error) {
	return u.postgres.FindUserByID(context.Background(), id)
}

// FindMe returns the detail information about user
func (u UserService) FindMe(id string) (dto.MyUserDTO, error) {
	user, err := u.postgres.FindMe(context.Background(), id)
	if err != nil {
		return dto.MyUserDTO{}, err
	}
	user.Email = user.Email[:1] + "**" + user.Email[strings.Index(user.Email, "@")-1:]
	return user, nil
}

func (u UserService) UpdateUser(customer dto.UpdateUserDTO, id string) error {
	return u.postgres.UpdateUser(context.Background(), customer, id)
}

func (u UserService) UpdatePassword(password, id string) error {
	return u.postgres.UpdatePassword(context.Background(), id, password)
}

func (u UserService) UpdateEmail(email, id string) error {
	return u.postgres.UpdateEmail(context.Background(), email, id)
}

func (u UserService) UpdateUsername(username, id string) error {
	return u.postgres.UpdateUsername(context.Background(), username, id)
}

func (u UserService) UsernameExist(username string) bool {
	return u.postgres.UsernameExist(context.Background(), username)
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
