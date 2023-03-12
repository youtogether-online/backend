package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
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
		c.Error(exceptions.ServerError.AddErr(err.Error()))
		return
	}

	dto.CutEmail(&user.Email)

	c.JSON(http.StatusOK, user)
}

// getUserById return the user with main info
func (h Handler) getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		c.Error(exceptions.ValidError.AddErr(err.Error()))
		return
	}

	user, err := h.users.FindUserById(id)
	if err != nil {
		if ent.IsNotFound(err) {
			c.Error(exceptions.LoginUnknown.AddErr(err.Error()))
		} else {
			c.Error(exceptions.ServerError.AddErr(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
