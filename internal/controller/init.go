package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/pkg/middleware/bind"
	"github.com/wtkeqrf0/you-together/pkg/middleware/session"
	"net/http"
)

type ErrHandler interface {
	HandleError(handler func(*gin.Context) error) gin.HandlerFunc
}

type SessionHandler interface {
	Session(handler func(*gin.Context, *dao.Session) error) func(c *gin.Context) error
	SessionFunc(c *gin.Context) (*dao.Session, error)
}

type QueryHandler interface {
	HandleQueries() gin.HandlerFunc
}

type Setter struct {
	r     *gin.Engine
	valid *validator.Validate
	erh   ErrHandler
	qh    QueryHandler
	sess  SessionHandler
}

func NewSetter(r *gin.Engine, valid *validator.Validate, erh ErrHandler, qh QueryHandler, sess SessionHandler) *Setter {
	return &Setter{r: r, valid: valid, erh: erh, qh: qh, sess: sess}
}

func (h *Handler) InitRoutes(s *Setter) {
	initMiddlewares(s.r, s.qh)

	rg := s.r.Group(h.cfg.Listen.QueryPath)

	auth := rg.Group("/auth")
	{
		auth.POST("/password", s.erh.HandleError(bind.HandleBodyWithHeader(h.signInByPassword, s.valid)))
		auth.POST("/email", s.erh.HandleError(bind.HandleBodyWithHeader(h.signInByEmail, s.valid)))

		sess := auth.Group("/session")
		{
			sess.GET("", s.erh.HandleError(s.sess.Session(h.getMe)))
			sess.DELETE("", s.erh.HandleError(h.signOut))
		}
	}

	user := rg.Group("/user")
	{
		user.GET("/:name", s.erh.HandleError(bind.HandleParam(h.getUserByUsername, s.valid)))
		user.PATCH("", s.erh.HandleError(session.HandleForm(h.updateUser, s.sess.SessionFunc, s.valid)))
		user.PATCH("/email", s.erh.HandleError(session.HandleJSONBody(h.updateEmail, s.sess.SessionFunc, s.valid)))
		user.PATCH("/image", s.erh.HandleError(session.HandleJSONBody(h.updateImage, s.sess.SessionFunc, s.valid)))
		user.PATCH("/password", s.erh.HandleError(session.HandleJSONBody(h.updatePassword, s.sess.SessionFunc, s.valid)))
		user.PATCH("/name", s.erh.HandleError(session.HandleJSONBody(h.updateUsername, s.sess.SessionFunc, s.valid)))
		user.GET("/check-name/:name", s.erh.HandleError(bind.HandleParam(h.checkUsername, s.valid)))
	}

	room := rg.Group("/room")
	{
		room.PUT("", s.erh.HandleError(session.HandleJSONBody(h.createRoom, s.sess.SessionFunc, s.valid)))
		room.GET("/:id", s.erh.HandleError(bind.HandleParam(h.joinRoom, s.valid)))
	}

	rg.Static("/file", "./"+h.cfg.Files.Path)

	if h.mail != nil {
		email := rg.Group("/email")
		{
			email.POST("/send-code", s.erh.HandleError(bind.HandleBody(h.sendCodeToEmail, s.valid)))
		}
	}
}

func initMiddlewares(r *gin.Engine, qh QueryHandler) {
	config := cors.Config{
		AllowOrigins:     []string{"https://youtogether.frkam.dev", "https://youtogether-online.github.io", "http://localhost:3000", "http://localhost:80", "http://localhost"},
		AllowMethods:     []string{http.MethodGet, http.MethodOptions, http.MethodPatch, http.MethodDelete, http.MethodPost, http.MethodHead},
		AllowHeaders:     []string{"Content-Code", "Content-Length", "Cache-Control", "User-Agent", "Accept-Language", "Accept", "DomainName", "Accept-Encoding", "Connection", "Set-Cookie", "Cookie", "Date", "Postman-Token", "Host"},
		AllowCredentials: true,
		AllowWebSockets:  true,
	}

	r.Use(qh.HandleQueries(), cors.New(config), gin.Recovery())
}
