package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you_together/internal/middleware/exceptions"
	"net/http"
)

// getMe using the session ID. Returns the user with detail info

// GetMe godoc
// @Summary Get detail information about user by session
// @Description Return detail information about the user
// @Param session_id header string true "user's session id"
// @Produce application/json
// @Tags User
// @Success 200 {object} dto.MyUserDTO
// @Failure 401 {object} exceptions.MyError "User isn't logged in"
// @Failure 500 {object} exceptions.MyError
// @Router /user [get]
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

// GetUserByUsername godoc
// @Summary Get information about the user
// @Description Return basic information about the user. If the user tries to find out information about himself, return detailed information about the user
// @Param username path string true "find user by username"
// @Produce application/json
// @Tags User
// @Success 200 {object} dto.UserDTO "dto.UserDTO or dto.MyUserDTO"
// @Failure 400 {object} exceptions.MyError
// @Failure 404 {object} exceptions.MyError "User doesn't exist"
// @Router /user/{username} [get]
func (h Handler) getUserByUsername(c *gin.Context) {
	username := c.Param("username")
	if err := h.valid.Var(username, "required,gte=5,lte=20"); err != nil {
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
