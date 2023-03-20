package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"github.com/wtkeqrf0/you_together/internal/middlewares/exceptions"
	"github.com/wtkeqrf0/you_together/pkg/conf"
	"time"
)

var cfg = conf.GetConfig()

// UserService interacts with the users table
type UserService interface {
	FindUserByUsername(username string) (dto.UserDTO, error)
	FindMe(username string) (dto.MyUserDTO, error)
	FindAllUsers(limit int) ([]dto.UserDTO, error)
	UpdateUser(user ent.User) error
	DeleteUser(id int) error

	SetVariable(key string, value any, exp time.Duration) error
	ContainsKeys(keys ...string) (int64, error)
}

type AuthService interface {
	SetCodes(key string, value ...any) error
	EqualsPopCode(email string, code string) (bool, error)
	GetSession(sessionId string) (map[string]string, error)
	SetSession(sessionId string, info map[string]string) error
	DelSession(sessionId string)

	CreateUserWithPassword(email, password string) ([]byte, string, error)
	CreateUserByEmail(email string) (string, error)
	AuthUserByEmail(email string) ([]byte, string, error)
	AuthUserWithInfo(email string) (bool, string, error)
	SetEmailVerified(email string) error
}

type AuthMiddleware interface {
	RequireSession(c *gin.Context)
	MaybeSession(c *gin.Context)
	GenerateSession(email, ip, device string) (string, map[string]string, error)
	ValidateSession(sessionId string) (map[string]string, error)
}

type Handler struct {
	users    UserService
	sessions AuthMiddleware
	auth     AuthService
	valid    *validator.Validate
}

func NewHandler(users UserService, sessions AuthMiddleware, auth AuthService, valid *validator.Validate) *Handler {
	return &Handler{users: users, sessions: sessions, auth: auth, valid: valid}
}

func (h Handler) InitRoutes(r *gin.Engine) {
	//TODO disable gin errors print
	r.Use(gin.Recovery(), gin.Logger(), exceptions.ErrorHandler)

	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/save-mail", h.saveMail)
		auth.POST("/check-mail", h.checkMail)
		auth.POST("/sign-out", h.signOut)
	}

	{
		api.GET("/:username", h.sessions.MaybeSession, h.getUserByUsername)
		api.GET("/", h.sessions.RequireSession, h.getMe)
	}

}
