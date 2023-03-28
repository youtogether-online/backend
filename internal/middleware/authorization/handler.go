package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you_together/internal/middleware/exceptions"
	"github.com/wtkeqrf0/you_together/pkg/conf"
	"net/http"
)

var cfg = conf.GetConfig()

// RequireSession authorizes the user
func (a Auth) RequireSession(c *gin.Context) {
	id, _ := c.Cookie(cfg.Session.CookieName)
	info, ok, err := a.ValidateSession(id)
	if err != nil {
		c.Error(exceptions.UnAuthorized.AddErr(err))
		return
	}

	if ok {
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie(cfg.Session.CookieName, id, cfg.Session.DurationInSeconds,
			"/api", cfg.Listen.Host, true, true)
	}

	c.Set("user_info", info)
	c.Next()
}

// MaybeSession authorizes the user, if the token exist. User can not be authorized
func (a Auth) MaybeSession(c *gin.Context) {
	id, _ := c.Cookie(cfg.Session.CookieName)
	info, ok, err := a.ValidateSession(id)
	if err != nil {
		c.Next()
		return
	}

	if ok {
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie(cfg.Session.CookieName, id, cfg.Session.DurationInSeconds,
			"/api", cfg.Listen.Host, true, true)
	}

	c.Set("user_info", info)
	c.Next()
}
