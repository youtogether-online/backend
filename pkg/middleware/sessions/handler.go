package sessions

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"net/http"
)

var cfg = conf.GetConfig()

// RequireSession authorizes the user and saves session info to context
func (a Auth) RequireSession(c *gin.Context) {
	session, _ := c.Cookie(cfg.Session.CookieName)
	info, ok, err := a.ValidateSession(session)
	if err != nil {
		c.Error(errs.UnAuthorized.AddErr(err))
		return
	}

	if ok {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie(cfg.Session.CookieName, session, int(cfg.Session.Duration.Seconds()),
			cfg.Session.CookiePath, cfg.Listen.DomainName, true, true)
	}

	c.Set("user_info", info)
	c.Next()
}

// MaybeSession authorizes the user and saves session info to context. User could not be authorized
func (a Auth) MaybeSession(c *gin.Context) {
	session, _ := c.Cookie(cfg.Session.CookieName)
	info, ok, err := a.ValidateSession(session)
	if err != nil {
		c.Next()
		return
	}

	if ok {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie(cfg.Session.CookieName, session, int(cfg.Session.Duration.Seconds()),
			cfg.Session.CookiePath, cfg.Listen.DomainName, true, true)
	}

	c.Set("user_info", info)
	c.Next()
}
