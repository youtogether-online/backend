package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"github.com/wtkeqrf0/you-together/pkg/middleware/bind"
	"time"
)

var (
	chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	cfg   = conf.GetConfig()
)

// UserService interacts with the users table
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

type AuthMiddleware interface {
	RequireSession(c *gin.Context)
	GenerateSession(id int, ip, userAgent string) (string, error)
	SetNewCookie(id int, c *gin.Context)
	GetSession(c *gin.Context) *dao.Session
	PopCookie(c *gin.Context)
}

type Validator interface {
	Struct(s any) error
	Var(s any, tag string) error
}

type Handler struct {
	users UserService
	auth  AuthService
	mail  MailSender
	sess  AuthMiddleware
	v     *bind.Bind
}

func NewHandler(users UserService, auth AuthService, mail MailSender, sess AuthMiddleware, v *bind.Bind) *Handler {
	return &Handler{users: users, auth: auth, mail: mail, sess: sess, v: v}
}

func (h *Handler) InitRoutes(rg *gin.RouterGroup, mailSet bool) {

	auth := rg.Group("/auth")
	{
		bind.HandleData(h.signInByPassword, h.v)
		auth.POST("/password", bind.HandleData(h.signInByPassword, h.v))
		auth.POST("/email", bind.HandleData(h.signInByEmail, h.v))

		session := rg.Group("/session")
		{
			session.GET("", h.sess.RequireSession, bind.HandleData(h.getMe, h.v))
			session.DELETE("", h.signOut)
		}
	}

	user := rg.Group("/user")
	{
		user.GET("/:username", bind.HandleData(h.getUserByUsername, h.v))
		user.PATCH("", h.sess.RequireSession, h.updateUser)
		user.PATCH("/email", h.sess.RequireSession, h.updateEmail)
		user.PATCH("/password", h.sess.RequireSession, h.updatePassword)
		user.PATCH("/name", h.sess.RequireSession, h.updateUsername)
		user.GET("/check-name/:name", bind.HandleData(h.checkUsername, h.v))
	}

	if mailSet {
		email := rg.Group("/email")
		{
			email.POST("/send-code", bind.HandleData(h.sendCodeToEmail, h.v))
		}
	}
}

type ErrHandler interface {
	HandleErrors() gin.HandlerFunc
}

type QueryHandler interface {
	HandleQueries() gin.HandlerFunc
}

type Middlewares struct {
	erh ErrHandler
	qh  QueryHandler
}

func NewMiddleWares(erh ErrHandler, qh QueryHandler) *Middlewares {
	return &Middlewares{erh: erh, qh: qh}
}

func (m *Middlewares) InitGlobalMiddleWares(r *gin.Engine) {
	r.Use(m.qh.HandleQueries(), gin.Recovery(), m.erh.HandleErrors())
}
