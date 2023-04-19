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
	FindUserByUsername(username string) (*dto.UserDTO, error)
	FindUserByID(id int) (*ent.User, error)
	FindMe(id int) (*dto.MyUserDTO, error)
	UpdateUser(customer dto.UpdateUserDTO, id int) error
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
	GetSession(sessionId string) (*dto.Session, error)
	SetSession(sessionId string, info dto.Session) error
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
	GenerateSession(id int, ip, userAgent string) (string, error)
	SetNewCookie(id int, c *gin.Context)
	GetSession(c *gin.Context) (*dto.Session, error)
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

	//api.GET("/:name", h.getTypeByName)

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
			email.POST("/sign-in", h.signInByEmail)
		}
	}

	session := api.Group("/session")
	{
		session.GET("", h.sessions.RequireSession, h.getMe)
		session.DELETE("", h.signOut)
	}

	user := api.Group("/user")
	{
		user.GET("/:username", h.sessions.MaybeSession, h.getUserByUsername)
		user.PATCH("", h.sessions.RequireSession, h.updateUser)
		user.PATCH("/email", h.sessions.RequireSession, h.updateEmail)
		user.PATCH("/password", h.sessions.RequireSession, h.updatePassword)
		user.PATCH("/name", h.sessions.RequireSession, h.updateUsername)
		user.GET("/check-name/:username", h.checkUsername)
	}

	email := api.Group("/email")
	{
		email.POST("/send-code", h.sendEmailCode)
	}
}

var valid = validator.New()

func fillStruct[T dto.DTO](c *gin.Context) (t T, ok bool) {
	c.ShouldBindJSON(&t)

	if err := valid.Struct(&t); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}
	ok = true
	return
}
