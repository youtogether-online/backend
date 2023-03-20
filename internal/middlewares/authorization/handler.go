package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you_together/internal/middlewares/exceptions"
	"github.com/wtkeqrf0/you_together/pkg/conf"
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

	c.Set("user_info", info)
	c.Next()
}
