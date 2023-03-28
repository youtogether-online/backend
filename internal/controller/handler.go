package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"github.com/wtkeqrf0/you_together/internal/middleware/exceptions"
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
	DelKeys(keys ...string)

	CreateUserWithPassword(email, password string) ([]byte, string, error)
	UserExistsByEmail(email string) bool
	CreateUserByEmail(email string) (string, error)
	AuthUserByEmail(email string) ([]byte, string, error)
	AuthUserWithInfo(email string) (bool, string, error)
	SetEmailVerified(email string) error
	FindSessionsByUsername(userName string) []map[string]string
}

type AuthMiddleware interface {
	RequireSession(c *gin.Context)
	MaybeSession(c *gin.Context)
	GenerateSession(email, ip, device string) (string, map[string]string, error)
	ValidateSession(sessionId string) (map[string]string, bool, error)
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
	r.Use(gin.Logger(), gin.Recovery(), exceptions.ErrorHandler)
	api := r.Group("/api")

	docs := api.Group("/docs")
	{
		docs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	auth := api.Group("/auth")
	{
		auth.POST("/sign-in-with-password", h.signInWithPassword)
		auth.POST("/sign-in-send-code", h.signInSendCode)
		auth.POST("/sign-in-check-code", h.signInCheckCode)
		auth.POST("/sign-out", h.signOut)
	}

	user := api.Group("/user")
	{
		user.GET("/:username", h.sessions.MaybeSession, h.getUserByUsername)
		user.GET("", h.sessions.RequireSession, h.getMe)
	}

}
