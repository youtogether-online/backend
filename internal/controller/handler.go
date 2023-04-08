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
	"go/types"
	"time"
)

var cfg = conf.GetConfig()

// UserService interacts with the users table
type UserService interface {
	FindUserByUsername(username string) (dto.UserDTO, error)
	FindUserByID(id string) (*ent.User, error)
	FindMe(id string) (dto.MyUserDTO, error)
	UpdateUser(customer dto.UpdateUserDTO, id string) error
	UpdatePassword(password, id string) error
	UpdateEmail(email, id string) error
	UpdateUsername(username, id string) error
	UsernameExist(username string) bool

	SetVariable(key string, value any, exp time.Duration) error
	ContainsKeys(keys ...string) (int64, error)
}

type AuthService interface {
	SetCodes(key string, value ...any) error
	EqualsPopCode(email string, code string) (bool, error)
	GetSession(sessionId string) (map[string]string, error)
	SetSession(sessionId string, info map[string]string) error
	DelKeys(keys ...string)

	CreateUserWithPassword(auth dto.EmailWithPasswordDTO) (*ent.User, error)
	UserExistsByEmail(email string) bool
	CreateUserByEmail(auth dto.EmailWithCodeDTO) (*ent.User, error)
	AuthUserByEmail(email string) (*ent.User, error)
	SetEmailVerified(email string) error
}

type AuthMiddleware interface {
	RequireSession(c *gin.Context)
	MaybeSession(c *gin.Context)
	GenerateSession(username, ip, userAgent string) (string, error)
	SetNewCookie(username string, c *gin.Context)
	GetSession(c *gin.Context) (map[string]string, error)
	PopCookie(c *gin.Context)
}

type Handler struct {
	users    UserService
	sessions AuthMiddleware
	auth     AuthService
}

func NewHandler(users UserService, sessions AuthMiddleware, auth AuthService) *Handler {
	return &Handler{users: users, sessions: sessions, auth: auth}
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
		pass := auth.Group("/password")
		{
			pass.POST("/sign-in", h.signInByPassword)
		}

		email := auth.Group("/email")
		{
			email.POST("/send-code", h.sendEmailCode)
			email.POST("/sign-in", h.signInByEmail)
		}

		auth.POST("/sign-out", h.signOut)
	}

	user := api.Group("/user")
	{
		user.GET("/:username", h.sessions.MaybeSession, h.getUserByUsername)
		user.GET("", h.sessions.RequireSession, h.getMe)

		update := user.Group("/upd", h.sessions.RequireSession)
		{
			update.PATCH("", h.updateUser)
			update.PATCH("/mail", h.updateEmail)
			update.PATCH("/pass", h.updatePassword)
			update.PATCH("/name", h.updateUsername)
			user.GET("/upd/:username", h.checkUsername)
		}
	}
}

var Valid = validator.New()

func fillStruct[T dto.DTO | types.Nil](c *gin.Context) (t T) {
	c.ShouldBindJSON(&t)

	if err := Valid.Struct(&t); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}
	return
}
