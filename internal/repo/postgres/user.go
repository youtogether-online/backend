package postgres

import (
	"context"
	"fmt"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/ent/user"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"strconv"
)

type UserStorage struct {
	userClient *ent.UserClient
}

func NewUserStorage(userClient *ent.UserClient) *UserStorage {
	return &UserStorage{userClient: userClient}
}

// FindMe returns the detail information about user
func (r *UserStorage) FindMe(ctx context.Context, username string) (dto.MyUserDTO, error) {
	var userDTO []dto.MyUserDTO

	err := r.userClient.Query().Where(user.ID(username)).
		Select(user.FieldID, user.FieldEmail, user.FieldIsEmailVerified,
			user.FieldBiography, user.FieldRole, user.FieldFriendsIds,
			user.FieldLanguage, user.FieldTheme, user.FieldFirstName,
			user.FieldLastName).Scan(ctx, &userDTO)

	if err != nil || len(userDTO) != 1 {
		return dto.MyUserDTO{}, fmt.Errorf("user not found: %v", err)
	}
	return userDTO[0], nil
}

// FindUserByUsername returns the main information about user
func (r *UserStorage) FindUserByUsername(ctx context.Context, username string) (dto.UserDTO, error) {
	var userDTO []dto.UserDTO

	err := r.userClient.Query().Where(user.ID(username)).
		Select(user.FieldID, user.FieldBiography,
			user.FieldRole, user.FieldFirstName, user.FieldLastName).
		Scan(ctx, &userDTO)

	if err != nil || len(userDTO) != 1 {
		return dto.UserDTO{}, fmt.Errorf("user not found: %v", err)
	}
	return userDTO[0], nil
}

// FindAllUsers with given limit
func (r *UserStorage) FindAllUsers(ctx context.Context, limit int) ([]*ent.User, error) {
	return r.userClient.Query().Limit(limit).All(ctx)
}

func (r *UserStorage) UpdateUser(ctx context.Context, user *ent.User) error {
	return r.userClient.UpdateOne(user).Exec(ctx)
}

func (r *UserStorage) DeleteUser(ctx context.Context, id int) error {
	return r.userClient.DeleteOneID(strconv.Itoa(id)).Exec(ctx)
}
