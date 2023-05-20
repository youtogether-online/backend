package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/conf"
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
	UpdateUser(customer dto.UpdateUser, id int) error
	UpdatePassword(password string, id int) error
	UpdateEmail(email string, id int) error
	UpdateUsername(username string, id int) error
	UsernameExist(username string) bool

	SetVariable(key string, value any, exp time.Duration) error
	ContainsKeys(keys ...string) (int64, error)
}

type AuthService interface {
	SetCodes(key string, value ...any) error
	EqualsPopCode(email string, code string) (bool, error)
	GetSession(sessionId string) (*dao.Session, error)
	SetSession(sessionId string, info dao.Session) error
	DelKeys(keys ...string)

	CreateUserWithPassword(auth dto.EmailWithPassword) (*ent.User, error)
	UserExistsByEmail(email string) bool
	CreateUserByEmail(auth dto.EmailWithCode) (*ent.User, error)
	AuthUserByEmail(email string) (*ent.User, error)
	SetEmailVerified(email string) error
}

type AuthMiddleware interface {
	RequireSession(c *gin.Context)
	MaybeSession(c *gin.Context)
	GenerateSession(id int, ip, userAgent string) (string, error)
	SetNewCookie(id int, userAgent string, c *gin.Context)
	GetSession(c *gin.Context) (*dao.Session, error)
	PopCookie(c *gin.Context)
}

type MailSender interface {
	SendEmail(subj, body, from string, to ...string) error
}

type Handler struct {
	users UserService
	auth  AuthService
	mail  MailSender
	sess  AuthMiddleware
}

func NewHandler(users UserService, auth AuthService, mail MailSender, sess AuthMiddleware) *Handler {
	return &Handler{users: users, auth: auth, mail: mail, sess: sess}
}

func (h *Handler) InitRoutes(rg *gin.RouterGroup, mailSet bool) {

	rg.StaticFile("/docs", "docs/OpenAPI.yaml")

	auth := rg.Group("/auth")
	{
		auth.POST("/password", h.signInByPassword)
		auth.POST("/email", h.signInByEmail)

		session := rg.Group("/session")
		{
			session.GET("", h.sess.RequireSession, h.getMe)
			session.DELETE("", h.signOut)
		}
	}

	user := rg.Group("/user")
	{
		user.GET("/:username", h.getUserByUsername)
		user.PATCH("", h.sess.RequireSession, h.updateUser)
		user.PATCH("/email", h.sess.RequireSession, h.updateEmail)
		user.PATCH("/password", h.sess.RequireSession, h.updatePassword)
		user.GET("/check-name/:name", h.checkUsername)
	}

	if mailSet {
		email := rg.Group("/email")
		{
			email.POST("/send-code", h.sendCodeToEmail)
		}
	}
}

type ErrHandler interface {
	HandleErrors(c *gin.Context)
}

type QueryHandler interface {
	HandleQueries(c *gin.Context)
}

type Middlewares struct {
	erh ErrHandler
	qh  QueryHandler
}

func NewMiddleWares(erh ErrHandler, qh QueryHandler) *Middlewares {
	return &Middlewares{erh: erh, qh: qh}
}

func (m *Middlewares) InitGlobalMiddleWares(r *gin.Engine) {
	r.Use(m.qh.HandleQueries, gin.Recovery(), m.erh.HandleErrors)
}
