package postgres

import (
	"context"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/ent/user"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/log"
)

type UserStorage struct {
	userClient *ent.UserClient
}

func NewUserStorage(userClient *ent.UserClient) *UserStorage {
	return &UserStorage{userClient: userClient}
}

// FindMe returns the detail information about user
func (r *UserStorage) FindMe(ctx context.Context, id int) (*dao.Me, error) {
	customer, err := r.userClient.Query().Where(user.ID(id)).
		Select(user.FieldName, user.FieldEmail, user.FieldIsEmailVerified,
			user.FieldBiography, user.FieldRole, user.FieldFriendsIds,
			user.FieldLanguage, user.FieldTheme, user.FieldFirstName,
			user.FieldLastName, user.FieldSessions, user.FieldCreateTime).Only(ctx)

	if err == nil {
		return dao.TransformToMe(customer), nil
	}

	return nil, err
}

// FindUserByUsername returns the main information about user
func (r *UserStorage) FindUserByUsername(ctx context.Context, username string) (*dao.User, error) {
	customer, err := r.userClient.Query().Where(user.Name(username)).
		Select(user.FieldName, user.FieldBiography, user.FieldRole,
			user.FieldFriendsIds, user.FieldFirstName, user.FieldLastName, user.FieldCreateTime).Only(ctx)

	if err == nil {
		return dao.TransformToUser(customer), nil
	}

	return nil, err
}

func (r *UserStorage) FindUserByID(ctx context.Context, id int) (*ent.User, error) {
	return r.userClient.Get(ctx, id)
}

func (r *UserStorage) UpdateUser(ctx context.Context, customer dto.UpdateUser, id int) error {
	res, err := r.userClient.UpdateOneID(id).
		SetNillableBiography(customer.Biography).
		SetNillableLanguage(customer.Language).
		SetNillableTheme(customer.Theme).
		SetNillableFirstName(customer.FirstName).
		SetNillableLastName(customer.LastName).Save(ctx)
	log.Debug(res)
	return err
}

func (r *UserStorage) UpdateEmail(ctx context.Context, email string, id int) error {
	res, err := r.userClient.UpdateOneID(id).SetEmail(email).
		SetIsEmailVerified(false).Save(ctx)
	log.Debug(res)
	return err
}

func (r *UserStorage) UpdatePassword(ctx context.Context, newPassword []byte, id int) error {
	res, err := r.userClient.UpdateOneID(id).SetPasswordHash(newPassword).
		SetIsEmailVerified(true).Save(ctx)
	log.Debug(res)
	return err
}

func (r *UserStorage) UpdateUsername(ctx context.Context, username string, id int) error {
	res, err := r.userClient.UpdateOneID(id).SetName(username).Save(ctx)
	log.Debug(res)
	return err
}

func (r *UserStorage) UsernameExist(ctx context.Context, username string) (bool, error) {
	return r.userClient.Query().Where(user.Name(username)).Exist(ctx)
}
