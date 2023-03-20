package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you_together/pkg/conf"
	"github.com/wtkeqrf0/you_together/pkg/middlewares/exceptions"
)

var cfg = conf.GetConfig()

// RequireSession authorizes the user
func (a Auth) RequireSession(c *gin.Context) {
	id, _ := c.Cookie(cfg.Session.CookieName)
	info, err := a.ValidateSession(id)
	if err != nil {
		c.Error(exceptions.UnAuthorized.AddErr(err))
		return
	}

	if a.auth.GetTTL(id) <= cfg.Session.Duration/2 {
		id, info, err = a.GenerateSession(info["username"], info["ip"], info["device"])

		if err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}

		c.SetCookie(cfg.Session.CookieName, id, cfg.Session.DurationInSeconds,
			"/api", cfg.Listen.Host, false, true)
	}

	c.Set("user_info", info)
	c.Next()
}

// MaybeSession authorizes the user, if the token exist. User can not be authorized
func (a Auth) MaybeSession(c *gin.Context) {
	id, _ := c.Cookie(cfg.Session.CookieName)
	info, err := a.ValidateSession(id)
	if err != nil {
		c.Next()
		return
	}

	if a.auth.GetTTL(id) <= cfg.Session.Duration/2 {
		id, info, err = a.GenerateSession(info["username"], info["ip"], info["device"])

		if err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}

		c.SetCookie(cfg.Session.CookieName, id, cfg.Session.DurationInSeconds,
			"/api", cfg.Listen.Host, false, true)
	}

	c.Set("user_info", info)
	c.Next()
}
