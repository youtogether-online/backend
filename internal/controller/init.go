package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"github.com/wtkeqrf0/you-together/pkg/middleware/bind"
	"github.com/wtkeqrf0/you-together/pkg/middleware/session"
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
	cfg   *conf.Config
}

func NewSetter(r *gin.Engine, valid *validator.Validate, erh ErrHandler, qh QueryHandler, sess SessionHandler, cfg *conf.Config) *Setter {
	return &Setter{r: r, valid: valid, erh: erh, qh: qh, sess: sess, cfg: cfg}
}

func (h *Handler) InitRoutes(s *Setter) {
	initMiddlewares(s.r, s.qh)

	rg := s.r.Group(s.cfg.Listen.QueryPath)

	auth := rg.Group("/auth")
	{
		auth.POST("/password", s.erh.HandleError(bind.HandleBodyWithHeader(h.signInByPassword, s.valid)))
		auth.POST("/email", s.erh.HandleError(bind.HandleBodyWithHeader(h.signInByEmail, s.valid)))

		sess := auth.Group("/session")
		{
			sess.GET("", s.erh.HandleError(s.sess.Session(h.getMe)))
			sess.DELETE("", s.erh.HandleError(func(c *gin.Context) error { return h.signOut(c, s.cfg.Session.CookieName) }))
		}
	}

	user := rg.Group("/user")
	{
		user.GET("/:name", s.erh.HandleError(bind.HandleParam(h.getUserByUsername, "name", "required,gte=5,lte=20,name", s.valid)))
		user.PATCH("", s.erh.HandleError(session.HandleBody(h.updateUser, s.sess.SessionFunc, s.valid)))
		user.PATCH("/email", s.erh.HandleError(session.HandleBody(h.updateEmail, s.sess.SessionFunc, s.valid)))
		user.PATCH("/password", s.erh.HandleError(session.HandleBody(h.updatePassword, s.sess.SessionFunc, s.valid)))
		user.PATCH("/name", s.erh.HandleError(session.HandleBody(h.updateUsername, s.sess.SessionFunc, s.valid)))
		user.GET("/checkName/:name", s.erh.HandleError(bind.HandleParam(h.checkUsername, "name", "required,gte=5,lte=20,name", s.valid)))
	}

	room := rg.Group("/room")
	{
		room.PUT("", s.erh.HandleError(session.HandleBody(h.createRoom, s.sess.SessionFunc, s.valid)))
		room.GET("/:name", s.erh.HandleError(h.joinRoom))
	}

	if h.mail != nil {
		email := rg.Group("/email")
		{
			email.POST("/sendCode", s.erh.HandleError(bind.HandleBody(h.sendCodeToEmail, s.valid)))
		}
	}
}

func initMiddlewares(r *gin.Engine, qh QueryHandler) {
	r.Use(qh.HandleQueries(), gin.Recovery())
}
