package service

import (
	"context"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
)

type UserStorage interface {
	FindMe(ctx context.Context, id int) (dto.MyUserDTO, error)
	FindUserById(ctx context.Context, id int) (dto.UserDTO, error)
	FindAllUsers(ctx context.Context, limit int) ([]*ent.User, error)
	UpdateUser(ctx context.Context, user *ent.User) error
	DeleteUser(ctx context.Context, id int) error
}

type UserService struct {
	storage UserStorage
}

func NewUserService(storage UserStorage) *UserService {
	return &UserService{storage: storage}
}

func (u UserService) FindUserById(id int) (dto.UserDTO, error) {
	return u.storage.FindUserById(context.Background(), id)
}

func (u UserService) FindMe(id int) (dto.MyUserDTO, error) {
	user, err := u.storage.FindMe(context.Background(), id)
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
