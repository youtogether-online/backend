package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"github.com/wtkeqrf0/you-together/pkg/middleware/bind"
	"github.com/wtkeqrf0/you-together/pkg/middleware/session"
	"net/http"
)

type ErrHandler interface {
	HandleError(handler func(*gin.Context) error) gin.HandlerFunc
}

type SessionHandler interface {
	Session(handler func(ctx *gin.Context, session *dao.Session) error) func(c *gin.Context) error
	SessionFunc(c *gin.Context) (*dao.Session, error)
}

type QueryHandler interface {
	HandleQueries() gin.HandlerFunc
}

type Setter struct {
	r        *gin.Engine
	valid    *validator.Validate
	erh      ErrHandler
	qh       QueryHandler
	sess     SessionHandler
	mainPath string
	mailSet  bool
}

func NewSetter(r *gin.Engine, valid *validator.Validate, erh ErrHandler, qh QueryHandler, sess SessionHandler, mainPath string, mailSet bool) *Setter {
	return &Setter{r: r, valid: valid, erh: erh, qh: qh, sess: sess, mainPath: mainPath, mailSet: mailSet}
}

func (h *Handler) InitRoutes(s *Setter) {
	initMiddlewares(s.r, s.qh)

	rg := s.r.Group(s.mainPath)

	auth := rg.Group("/auth")
	{
		auth.POST("/password", s.erh.HandleError(bind.HandleBodyWithHeader(h.signInByPassword, s.valid)))
		auth.POST("/email", s.erh.HandleError(bind.HandleBodyWithHeader(h.signInByEmail, s.valid)))

		sess := rg.Group("/session")
		{
			sess.GET("", s.erh.HandleError(s.sess.Session(h.getMe)))
			sess.DELETE("", s.erh.HandleError(s.sess.Session(h.signOut)))
		}
	}

	user := rg.Group("/user")
	{
		user.GET("/:username", s.erh.HandleError(bind.HandleParam(h.getUserByUsername, "username", "required,gte=5,lte=20,name", s.valid)))
		user.PATCH("", s.erh.HandleError(session.HandleBody(h.updateUser, s.sess.SessionFunc, s.valid)))
		user.PATCH("/email", s.erh.HandleError(session.HandleBody(h.updateEmail, s.sess.SessionFunc, s.valid)))
		user.PATCH("/password", s.erh.HandleError(session.HandleBody(h.updatePassword, s.sess.SessionFunc, s.valid)))
		user.PATCH("/name", s.erh.HandleError(session.HandleBody(h.updateUsername, s.sess.SessionFunc, s.valid)))
		user.GET("/check-name/:name", s.erh.HandleError(bind.HandleParam(h.checkUsername, "name", "required,gte=5,lte=20,name", s.valid)))
	}

	if s.mailSet {
		email := rg.Group("/email")
		{
			email.POST("/send-code", s.erh.HandleError(bind.HandleBody(h.sendCodeToEmail, s.valid)))
		}
	}
}

func initMiddlewares(r *gin.Engine, qh QueryHandler) {
	config := cors.Config{
		AllowOrigins:     []string{"https://youtogether.frkam.dev", "https://youtogether-online.github.io", "http://localhost:3000", "http://localhost:80", "http://localhost"},
		AllowMethods:     []string{http.MethodGet, http.MethodOptions, http.MethodPatch, http.MethodDelete, http.MethodPost},
		AllowHeaders:     []string{"Content-Code", "Content-Length", "Cache-Control", "User-Agent", "Accept-Language", "Accept", "DomainName", "Accept-Encoding", "Connection", "Set-Cookie", "Cookie", "Date", "Postman-Token", "Host"},
		AllowCredentials: true,
		AllowWebSockets:  true,
	}

	r.Use(qh.HandleQueries(), cors.New(config), gin.Recovery())

	if err := r.SetTrustedProxies([]string{"95.140.155.222"}); err != nil {
		log.WithErr(err).Fatal("can't set trusted proxies")
	}
}