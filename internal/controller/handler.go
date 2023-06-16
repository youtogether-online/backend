package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
)

type UserService interface {
	FindUserByUsername(username string) (*dao.User, error)
	FindUserByID(id int) (*ent.User, error)
	FindMe(id int) (*dao.Me, error)

	UpdateUser(customer dto.UpdateUser, id int) error
	UpdatePassword(newPassword []byte, id int) error
	UpdateEmail(email string, id int) error
	UpdateUsername(username string, id int) error
	UsernameExist(username string) (bool, error)
}

type RoomService interface {
	Create(rm dto.Room, creatorId int) (*ent.Room, error)
}

type AuthService interface {
	SetCodes(key string, value ...any) error
	EqualsPopCode(email string, code string) (bool, error)
	DelKeys(keys ...string)
	CompareHashAndPassword(old, new []byte) error

	CreateUserWithPassword(email string, password []byte, language *string) (*ent.User, error)
	CreateUserByEmail(email string, language *string) (*ent.User, error)
	AuthUserByEmail(email string) (*ent.User, error)
	SetEmailVerified(email string) error
}

type MailSender interface {
	SendEmail(subj, body, from string, to ...string) error
}

type Session interface {
	SetNewCookie(id int, c *gin.Context) error
	GenerateSecretCode() string
}

type Handler struct {
	user UserService
	room RoomService
	auth AuthService
	mail MailSender
	sess Session
}

func NewHandler(user UserService, room RoomService, auth AuthService, mail MailSender, sess Session) *Handler {
	return &Handler{user: user, room: room, auth: auth, mail: mail, sess: sess}
}
