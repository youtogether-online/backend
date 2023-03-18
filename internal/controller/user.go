package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you_together/pkg/middlewares/exceptions"
	"net/http"
	"strconv"
)

// getMe using JWT. Return the user with detail info
func (h Handler) getMe(c *gin.Context) {
	userId, ok := c.Get("ID")
	if !ok {
		c.Error(exceptions.ServerError)
		return
	}

	user, err := h.users.FindMe(userId.(int))
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

// getUserById return the user with main info
func (h Handler) getUserById(c *gin.Context) {
	val := c.Param("ID")
	if err := h.valid.Var(val, "required,numeric"); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}
	id, _ := strconv.Atoi(val)

	if err := h.valid.Var(id, "required,gte=0"); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}

	var user any
	var err error

	if userId, ok := c.Get("ID"); ok && userId.(int) == id {
		user, err = h.users.FindMe(userId.(int))
	} else {
		user, err = h.users.FindUserById(id)
	}

	if err != nil {
		c.Error(exceptions.LoginUnknown.AddErr(err))
		return
	}

	c.JSON(http.StatusOK, user)
}
