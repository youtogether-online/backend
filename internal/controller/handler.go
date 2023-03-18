package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"github.com/wtkeqrf0/you_together/pkg/conf"
	"github.com/wtkeqrf0/you_together/pkg/middlewares/exceptions"
	"time"
)

var cfg = conf.GetConfig()

// UserService interacts with the users table
type UserService interface {
	FindUserById(id int) (dto.UserDTO, error)
	FindMe(id int) (dto.MyUserDTO, error)
	FindAllUsers(limit int) ([]dto.UserDTO, error)
	UpdateUser(user ent.User) error
	DeleteUser(id int) error
}

// RedisService interacts with redis
type RedisService interface {
	SetVariable(key string, value any, exp time.Duration) error
	ContainsKeys(keys ...string) (int64, error)
	SetCodes(key string, value ...any) error
	ContainsPopCode(email string, code string) bool
}

// Authorization requests
type Authorization interface {
	CreateUser(email, password string) (*ent.User, error)
	AuthUserByEmail(email string) (*ent.User, error)
	RequireAuth(c *gin.Context)
	MaybeAuth(c *gin.Context)
	GenerateJWT(id float64) (dto.TokensDTO, error)
	ValidateJWT(token string) (jwt.MapClaims, error)
}

type Handler struct {
	users UserService
	redis RedisService
	auth  Authorization
	valid *validator.Validate
}

func NewHandler(users UserService, redis RedisService, auth Authorization, valid *validator.Validate) *Handler {
	return &Handler{users: users, redis: redis, auth: auth, valid: valid}
}

func (h Handler) InitRoutes(r *gin.Engine) {
	eng := en.New()
	ut.New(eng, eng)
	r.Use(gin.Recovery(), gin.Logger(), exceptions.ErrorHandler)

	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/send-code", h.sendSecretCode)
		auth.POST("/check-code", h.compareSecretCode)
		auth.POST("/logout", h.logout)
	}

	{
		api.GET("/:ID", h.auth.MaybeAuth, h.getUserById)
		api.GET("/", h.auth.RequireAuth, h.getMe)
	}

}
