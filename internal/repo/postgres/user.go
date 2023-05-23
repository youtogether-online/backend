package postgres

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/ent/user"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
)

type UserStorage struct {
	userClient *ent.UserClient
}

func NewUserStorage(userClient *ent.UserClient) *UserStorage {
	return &UserStorage{userClient: userClient}
}

// FindMe returns the detail information about user
func (r *UserStorage) FindMe(ctx context.Context, id int) (*dao.Me, error) {
	var me []*dao.Me

	err := r.userClient.Query().Where(user.ID(id)).
		Select(user.FieldName, user.FieldEmail, user.FieldIsEmailVerified,
			user.FieldBiography, user.FieldRole, user.FieldFriendsIds,
			user.FieldLanguage, user.FieldTheme, user.FieldFirstName,
			user.FieldLastName).Scan(ctx, &me)

	if me != nil {
		return me[0], nil
	}

	return nil, err
}

// FindUserByUsername returns the main information about user
func (r *UserStorage) FindUserByUsername(ctx context.Context, username string) (*dao.User, error) {
	var customer []*dao.User

	err := r.userClient.Query().Where(user.Name(username)).
		Select(user.FieldName, user.FieldBiography, user.FieldRole,
			user.FieldFriendsIds, user.FieldFirstName, user.FieldLastName).
		Scan(ctx, &customer)

	if customer != nil {
		return customer[0], nil
	}

	return nil, err
}

func (r *UserStorage) FindUserByID(ctx context.Context, id int) (*ent.User, error) {
	return r.userClient.Get(ctx, id)
}

func (r *UserStorage) UpdateUser(ctx context.Context, customer dto.UpdateUser, id int) error {
	updCustomer, err := r.userClient.Update().
		SetNillableBiography(customer.Biography).
		SetNillableLanguage(customer.Language).
		SetNillableTheme(customer.Theme).
		SetNillableFirstName(customer.FirstName).
		SetNillableLastName(customer.LastName).Where(user.ID(id)).Save(ctx)

	logrus.WithError(err).Debug(updCustomer)
	return err
}

func (r *UserStorage) UpdateEmail(ctx context.Context, email string, id int) error {
	return r.userClient.Update().SetEmail(email).
		SetIsEmailVerified(false).Where(user.ID(id)).Exec(ctx)
}

func (r *UserStorage) UpdatePassword(ctx context.Context, password string, id int) error {
	return r.userClient.Update().SetPasswordHash([]byte(password)).
		SetIsEmailVerified(true).Where(user.ID(id)).Exec(ctx)
}

func (r *UserStorage) UpdateUsername(ctx context.Context, username string, id int) error {
	return r.userClient.Update().SetName(username).Where(user.ID(id)).Exec(ctx)
}

func (r *UserStorage) UsernameExist(ctx context.Context, username string) (bool, error) {
	return r.userClient.Query().Where(user.Name(username)).Exist(ctx)
}
