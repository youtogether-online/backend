package postgres

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/ent/user"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/internal/middleware/exceptions"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
	userClient *ent.UserClient
}

func NewUserStorage(userClient *ent.UserClient) *UserStorage {
	return &UserStorage{userClient: userClient}
}

// FindMe returns the detail information about user
func (r *UserStorage) FindMe(ctx context.Context, id int) (*dto.MyUserDTO, error) {
	userDTO := new([]dto.MyUserDTO)

	err := r.userClient.Query().Where(user.ID(id)).
		Select(user.FieldName, user.FieldEmail, user.FieldIsEmailVerified,
			user.FieldBiography, user.FieldRole, user.FieldFriendsIds,
			user.FieldLanguage, user.FieldTheme, user.FieldFirstName,
			user.FieldLastName).Scan(ctx, userDTO)

	if len(*userDTO) < 1 {
		return nil, exceptions.NoSuchUser

	} else if err != nil {
		return nil, err
	}

	return &(*userDTO)[0], nil
}

// FindUserByUsername returns the main information about user
func (r *UserStorage) FindUserByUsername(ctx context.Context, username string) (*dto.UserDTO, error) {
	userDTO := new([]dto.UserDTO)

	err := r.userClient.Query().Where(user.Name(username)).
		Select(user.FieldName, user.FieldBiography, user.FieldRole,
			user.FieldFriendsIds, user.FieldFirstName, user.FieldLastName).
		Scan(ctx, userDTO)

	if len(*userDTO) < 1 {
		return nil, exceptions.NoSuchUser

	} else if err != nil {
		return nil, err
	}

	return &(*userDTO)[0], nil
}

func (r *UserStorage) FindUserByID(ctx context.Context, id int) (*ent.User, error) {
	return r.userClient.Get(ctx, id)
}

func (r *UserStorage) UpdateUser(ctx context.Context, customer dto.UpdateUserDTO, id int) error {
	updCustomer, err := r.userClient.Update().
		SetNillableBiography(customer.Biography).
		SetNillableLanguage(customer.Language).
		SetNillableTheme(customer.Theme).
		SetNillableFirstName(customer.FirstName).
		SetNillableLastName(customer.LastName).Where(user.ID(id)).Save(ctx)

	logrus.WithError(err).Info(updCustomer)
	return err
}

func (r *UserStorage) UpdateEmail(ctx context.Context, email string, id int) error {
	res, err := r.userClient.Update().SetEmail(email).
		SetIsEmailVerified(false).Where(user.ID(id)).Save(ctx)
	logrus.WithError(err).Info(res)
	return err
}

func (r *UserStorage) UpdatePassword(ctx context.Context, password string, id int) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	customer, err := r.userClient.Update().SetPasswordHash(hashedPassword).
		SetIsEmailVerified(true).Where(user.ID(id)).Save(ctx)

	logrus.WithError(err).Info(customer)
	return err
}

func (r *UserStorage) UpdateUsername(ctx context.Context, username string, id int) error {
	customer, err := r.userClient.Update().SetName(username).Where(user.ID(id)).Save(ctx)
	logrus.WithError(err).Info(customer)
	return err
}

func (r *UserStorage) UsernameExist(ctx context.Context, username string) bool {
	return r.userClient.Query().Where(user.Name(username)).ExistX(ctx)
}
