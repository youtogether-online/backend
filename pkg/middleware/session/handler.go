package session

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"net/http"
)

func (a Auth) Session(handler func(*gin.Context, *dao.Session) error) func(c *gin.Context) error {
	return func(c *gin.Context) error {
		session, _ := c.Cookie(a.cfg.Session.CookieName)
		info, ok, err := a.ValidateSession(session)
		if err != nil {
			return errs.UnAuthorized.AddErr(err)
		}

		if ok {
			c.SetSameSite(http.SameSiteNoneMode)
			c.SetCookie(a.cfg.Session.CookieName, session, int(a.cfg.Session.Duration.Seconds()),
				a.cfg.Session.CookiePath, a.cfg.Session.Domain, true, true)
		}

		return handler(c, info)
	}
}

func (a Auth) SessionFunc(c *gin.Context) (*dao.Session, error) {
	session, _ := c.Cookie(a.cfg.Session.CookieName)
	info, ok, err := a.ValidateSession(session)
	if err != nil {
		return nil, errs.UnAuthorized.AddErr(err)
	}

	if ok {
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie(a.cfg.Session.CookieName, session, int(a.cfg.Session.Duration.Seconds()),
			a.cfg.Session.CookiePath, a.cfg.Session.Domain, true, true)
	}

	return info, nil
}

func HandleBody[T any](handler func(*gin.Context, T, *dao.Session) error, auth func(*gin.Context) (*dao.Session, error), v *validator.Validate) func(*gin.Context) error {
	return func(c *gin.Context) error {

		info, err := auth(c)
		if err != nil {
			return err
		}

		var t T
		if err = c.ShouldBindJSON(&t); err != nil {
			return err
		} else if err = v.Struct(&t); err != nil {
			return err
		}

		return handler(c, t, info)

	}
}
