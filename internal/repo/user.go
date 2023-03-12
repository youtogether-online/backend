package repo

import (
	"context"
	"fmt"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/ent/user"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
)

type UserStorage struct {
	userClient *ent.UserClient
}

// FindMe return detail information about user
func (r UserStorage) FindMe(ctx context.Context, id int) (dto.MyUserDTO, error) {
	var userDTO []dto.MyUserDTO

	err := r.userClient.Query().Where(user.ID(id)).Select("user_name", "email", "biography", "role", "avatar",
		"friends_ids", "language", "theme", "name").Scan(ctx, &userDTO)
	if err != nil || len(userDTO) != 1 {
		return dto.MyUserDTO{}, fmt.Errorf("user not found: %v", err)
	}
	return userDTO[0], err
}

// FindUserById return main information about user
func (r UserStorage) FindUserById(ctx context.Context, id int) (dto.UserDTO, error) {
	var userDTO []dto.UserDTO

	err := r.userClient.Query().Where(user.ID(id)).Select("user_name", "biography", "role",
		"avatar", "name").Scan(ctx, &userDTO)
	if err != nil || len(userDTO) != 1 {
		return dto.UserDTO{}, fmt.Errorf("user not found: %v", err)
	}
	return userDTO[0], nil
}

// FindAllUsers with given limit
func (r UserStorage) FindAllUsers(ctx context.Context, limit int) ([]*ent.User, error) {
	return r.userClient.Query().Limit(limit).All(ctx)
}

func (r UserStorage) UpdateUser(ctx context.Context, user *ent.User) error {
	return r.userClient.UpdateOne(user).Exec(ctx)
}

func (r UserStorage) DeleteUser(ctx context.Context, id int) error {
	return r.userClient.DeleteOneID(id).Exec(ctx)
}

func NewUserStorage(client *ent.UserClient) UserStorage {
	return UserStorage{userClient: client}
}
