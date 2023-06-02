package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"time"
)

type UserService interface {
	FindUserByUsername(username string) (*dao.User, error)
	FindUserByID(id int) (*ent.User, error)
	FindMe(id int) (*dao.Me, error)

	UpdateUser(customer *dto.UpdateUser, id int) error
	UpdatePassword(newPassword []byte, id int) error
	UpdateEmail(email string, id int) error
	UpdateUsername(username string, id int) error
	UsernameExist(username string) (bool, error)

	SetVariable(key string, value any, exp time.Duration) error
	ContainsKeys(keys ...string) (int64, error)
}

type AuthService interface {
	SetCodes(key string, value ...any) error
	EqualsPopCode(email string, code string) (bool, error)
	GetSession(sessionId string) (*dao.Session, error)
	SetSession(sessionId string, info dao.Session) error
	DelKeys(keys ...string)

	CreateUserWithPassword(email string, password []byte, language *string) (*ent.User, error)
	UserExistsByEmail(email string) (bool, error)
	CreateUserByEmail(email string, language *string) (*ent.User, error)
	AuthUserByEmail(email string) (*ent.User, error)
	SetEmailVerified(email string) error
}

type MailSender interface {
	SendEmail(subj, body, from string, to ...string) error
}

type Sessions interface {
	GenerateSession(id int, ip, userAgent string) (string, error)
	SetNewCookie(id int, c *gin.Context)
	GetSession(c *gin.Context) *dao.Session
	GenerateSecretCode() string
}

type Handler struct {
	users UserService
	auth  AuthService
	mail  MailSender
	sess  Sessions
}

func NewHandler(users UserService, auth AuthService, mail MailSender, sess Sessions) *Handler {
	return &Handler{users: users, auth: auth, mail: mail, sess: sess}
}
