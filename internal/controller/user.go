package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you_together/pkg/middlewares/exceptions"
	"net/http"
)

// getMe using the session ID. Returns the user with detail info
func (h Handler) getMe(c *gin.Context) {
	get, ok := c.Get("user_info")
	if !ok {
		c.Error(exceptions.ServerError)
		return
	}
	info := get.(map[string]string)

	user, err := h.users.FindMe(info["username"])
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

// getUserByUsername returns the user with the main info or detail info (if session id username equals username)
func (h Handler) getUserByUsername(c *gin.Context) {
	username := c.Param("username")
	if err := h.valid.Var(username, "required,gte=5"); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}

	var user any
	var err error

	if get, ok := c.Get("user_info"); ok && get.(map[string]string)["username"] == username {
		user, err = h.users.FindMe(username)
	} else {
		user, err = h.users.FindUserByUsername(username)
	}

	if err != nil {
		c.Error(exceptions.SignInUnknown.AddErr(err))
		return
	}

	c.JSON(http.StatusOK, user)
}
