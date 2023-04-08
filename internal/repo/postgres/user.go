package postgres

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/ent/user"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"github.com/wtkeqrf0/you_together/internal/middleware/exceptions"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
	userClient *ent.UserClient
}

func NewUserStorage(userClient *ent.UserClient) *UserStorage {
	return &UserStorage{userClient: userClient}
}

// FindMe returns the detail information about user
func (r *UserStorage) FindMe(ctx context.Context, id string) (dto.MyUserDTO, error) {
	var userDTO []dto.MyUserDTO

	err := r.userClient.Query().Where(user.ID(id)).
		Select(user.FieldID, user.FieldEmail, user.FieldIsEmailVerified,
			user.FieldBiography, user.FieldRole, user.FieldFriendsIds,
			user.FieldLanguage, user.FieldTheme, user.FieldFirstName,
			user.FieldLastName).Scan(ctx, &userDTO)

	if len(userDTO) != 1 {
		return dto.MyUserDTO{}, ent.MaskNotFound(err)
	} else if err != nil {
		return dto.MyUserDTO{}, err
	}
	return userDTO[0], nil
}

// FindUserByUsername returns the main information about user
func (r *UserStorage) FindUserByUsername(ctx context.Context, username string) (dto.UserDTO, error) {
	var userDTO []dto.UserDTO

	err := r.userClient.Query().Where(user.Name(username)).
		Select(user.FieldID, user.FieldBiography,
			user.FieldRole, user.FieldFirstName, user.FieldLastName).
		Scan(ctx, &userDTO)

	if len(userDTO) != 1 {
		return dto.UserDTO{}, exceptions.NoSuchUser

	} else if err != nil {
		return dto.UserDTO{}, err
	}

	return userDTO[0], nil
}

func (r *UserStorage) FindUserByID(ctx context.Context, id string) (*ent.User, error) {
	return r.userClient.Get(ctx, id)
}

func (r *UserStorage) UpdateUser(ctx context.Context, customer dto.UpdateUserDTO, id string) error {
	updCustomer, err := r.userClient.Update().
		SetNillableBiography(&customer.Biography).
		SetNillableLanguage(&customer.Language).
		SetNillableTheme(&customer.Theme).
		SetNillableFirstName(&customer.FirstName).
		SetNillableLastName(&customer.LastName).Where(user.Name(id)).Save(ctx)

	logrus.WithError(err).Info(updCustomer)
	return err
}

func (r *UserStorage) UpdateEmail(ctx context.Context, email, id string) error {
	res, err := r.userClient.Update().SetEmail(email).
		SetIsEmailVerified(false).Where(user.Name(id)).Save(ctx)
	logrus.WithError(err).Info(res)
	return err
}

func (r *UserStorage) UpdatePassword(ctx context.Context, password, id string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	customer, err := r.userClient.Update().SetPasswordHash(hashedPassword).
		SetIsEmailVerified(true).Where(user.Name(id)).Save(ctx)

	logrus.WithError(err).Info(customer)
	return err
}

func (r *UserStorage) UpdateUsername(ctx context.Context, username, id string) error {
	customer, err := r.userClient.Update().SetName(username).Where(user.Name(id)).Save(ctx)
	logrus.WithError(err).Info(customer)
	return err
}

func (r *UserStorage) UsernameExist(ctx context.Context, username string) bool {
	return r.userClient.Query().Where(user.Name(username)).ExistX(ctx)
}
