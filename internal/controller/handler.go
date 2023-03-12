package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"github.com/wtkeqrf0/you_together/pkg/conf"
	"github.com/wtkeqrf0/you_together/pkg/middlewares/exceptions"
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
	SetVariable(key string, value ...any) error
	GetVariables(key string) ([]string, error)
	ContainsVariable(key string, code string) bool
}

// Authorization requests
type Authorization interface {
	CreateUser(email, password string) (*ent.User, error)
	AuthUserByEmail(email string) (*ent.User, error)
	RequireAuth(c *gin.Context)
	GenerateJWT(id int, c *gin.Context) error
}

type Handler struct {
	users UserService
	redis RedisService
	auth  Authorization
}

func NewHandler(users UserService, redis RedisService, auth Authorization) *Handler {
	return &Handler{users: users, redis: redis, auth: auth}
}

func (h Handler) InitRoutes(r *gin.Engine) {
	r.Use(gin.Recovery(), gin.Logger(), exceptions.ErrorHandler)

	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/send-code", h.sendSecretCode)
		auth.POST("/check-code", h.compareSecretCode)
	}

	users := api.Group("/user")
	{
		users.GET("/:ID", h.getUserById)
		users.GET("/", h.auth.RequireAuth, h.getMe)
	}

}
