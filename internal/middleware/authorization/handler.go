package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/internal/middleware/exceptions"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"net/http"
)

var cfg = conf.GetConfig()

// RequireSession authorizes the user
func (a Auth) RequireSession(c *gin.Context) {
	session, _ := c.Cookie(cfg.Session.CookieName)
	info, ok, err := a.ValidateSession(session)
	if err != nil {
		c.Error(exceptions.UnAuthorized.AddErr(err))
		return
	}

	if ok {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie(cfg.Session.CookieName, session, int(cfg.Session.Duration.Seconds()),
			cfg.Session.CookiePath, cfg.Listen.Host, true, true)
	}

	c.Set("user_info", info)
	c.Next()
}

// MaybeSession authorizes the user, if the token exist. User can not be authorized
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
			cfg.Session.CookiePath, cfg.Listen.Host, true, true)
	}

	c.Set("user_info", info)
	c.Next()
}
