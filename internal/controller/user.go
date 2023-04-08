package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"github.com/wtkeqrf0/you_together/internal/middleware/exceptions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// GetMe godoc
// @Summary Get detail information about user by session
// @Description Return detail information about the user (cookie required)
// @Tags User Get
// @Success 200 {object} dto.MyUserDTO
// @Failure 401 {object} exceptions.MyError "User isn't logged in"
// @Failure 404 {object} exceptions.MyError "User doesn't exist"
// @Failure 500 {object} exceptions.MyError
// @Router /user [get]
func (h Handler) getMe(c *gin.Context) {
	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	user, err := h.users.FindMe(info["id"])
	if err != nil {
		if _, ok := err.(exceptions.MyError); ok {
			c.Error(exceptions.NoSuchUser)
		} else {
			c.Error(exceptions.ServerError.AddErr(err))
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByUsername godoc
// @Summary Get main information about the user
// @Description Return main information about the user. If the user tries to find out information about himself, return detailed information about the user (can accept cookie)
// @Param username header string true "the name of the desired user to find"
// @Tags User Get
// @Success 200 {object} dto.UserDTO "dto.UserDTO or dto.MyUserDTO"
// @Failure 400 {object} exceptions.MyError
// @Failure 404 {object} exceptions.MyError "User doesn't exist"
// @Failure 500 {object} exceptions.MyError
// @Router /user/{username} [get]
func (h Handler) getUserByUsername(c *gin.Context) {
	username := c.Param("username")
	if err := Valid.Struct(&dto.UsernameDTO{Username: username}); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}

	var (
		user any
		err  error
	)

	info, err := h.sessions.GetSession(c)
	if err != nil {
		user, err = h.users.FindUserByUsername(username)
	} else {
		customer, err := h.users.FindUserByID(info["id"])
		if err != nil {
			user, err = h.users.FindUserByUsername(username)
		} else {
			user = dto.Convert(customer)
		}
	}

	if err != nil {
		if _, ok := err.(exceptions.MyError); ok {
			c.Error(exceptions.NoSuchUser)
		} else {
			c.Error(exceptions.ServerError.AddErr(err))
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update user's information
// @Description Change user's main information (cookie required)
// @Param UpdateUserDTO body dto.UpdateUserDTO true "New user Information"
// @Tags User Update
// @Success 200
// @Failure 400 {object} exceptions.MyError
// @Failure 401 {object} exceptions.MyError "User isn't logged in"
// @Failure 500 {object} exceptions.MyError
// @Router /user/upd [patch]
func (h Handler) updateUser(c *gin.Context) {
	upd := fillStruct[dto.UpdateUserDTO](c)
	if c.Errors.Last() != nil {
		return
	}

	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	if err = h.users.UpdateUser(upd, info["id"]); err != nil {
		if ent.IsNotFound(err) {
			c.Error(exceptions.NoSuchUser.AddErr(err))
		} else {
			c.Error(exceptions.ServerError.AddErr(err))
		}
		return
	}

	c.Status(http.StatusOK)
}

// UpdateEmail godoc
// @Summary Update user's email
// @Description Change user's email by password (cookie required)
// @Param UpdateEmailDTO body dto.UpdateEmailDTO true "user's password and new email"
// @Tags User Update
// @Success 200
// @Failure 400 {object} exceptions.MyError
// @Failure 401 {object} exceptions.MyError "User isn't logged in"
// @Failure 500 {object} exceptions.MyError
// @Router /user/upd/mail [patch]
func (h Handler) updateEmail(c *gin.Context) {
	upd := fillStruct[dto.UpdateEmailDTO](c)
	if c.Errors.Last() != nil {
		return
	}

	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	user, err := h.users.FindUserByID(info["id"])

	if err != nil {
		if _, ok := err.(exceptions.MyError); ok {
			c.Error(exceptions.NoSuchUser)
		} else {
			c.Error(exceptions.ServerError.AddErr(err))
		}
		return
	} else if user.PasswordHash == nil || len(user.PasswordHash) < 1 {
		c.Error(exceptions.PasswordNotFound)
		return
	}

	if err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(upd.Password)); err != nil {
		c.Error(exceptions.PasswordError.AddErr(err))
		return
	}

	if err = h.users.UpdateEmail(upd.NewEmail, info["id"]); err != nil {
		if ent.IsValidationError(err) {
			c.Error(exceptions.AlreadyExist.AddErr(err))
		} else {
			c.Error(exceptions.ServerError.AddErr(err))
		}
		return
	}

	c.Status(http.StatusOK)
}

// UpdatePassword godoc
// @Summary Update user's password
// @Description Change user's password by email (cookie required)
// @Param UpdatePasswordDTO body dto.UpdatePasswordDTO true "user's email, code and new password"
// @Tags User Update
// @Success 200
// @Failure 400 {object} exceptions.MyError
// @Failure 401 {object} exceptions.MyError "User isn't logged in"
// @Failure 500 {object} exceptions.MyError
// @Router /user/upd/pass [patch]
func (h Handler) updatePassword(c *gin.Context) {
	upd := fillStruct[dto.UpdatePasswordDTO](c)
	if c.Errors.Last() != nil {
		return
	}

	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	if ok, err := h.auth.EqualsPopCode(upd.Email, upd.Code); !ok {
		c.Error(exceptions.CodeError.AddErr(err))
		return
	} else if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	if err = h.users.UpdatePassword(upd.NewPassword, info["id"]); err != nil {
		if ent.IsNotFound(err) {
			c.Error(exceptions.NoSuchUser.AddErr(err))
		} else {
			c.Error(exceptions.ServerError.AddErr(err))
		}
		return
	}

	c.Status(http.StatusOK)
}

// UpdateUsername godoc
// @Summary Update user's username
// @Description Change user's username (cookie required)
// @Tags User Update
// @Success 200
// @Failure 400 {object} exceptions.MyError
// @Failure 401 {object} exceptions.MyError "User isn't logged in"
// @Failure 500 {object} exceptions.MyError
// @Router /user/upd/name [patch]
func (h Handler) updateUsername(c *gin.Context) {
	upd := fillStruct[dto.UpdateUsernameDTO](c)
	if c.Errors.Last() != nil {
		return
	}

	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	if err = h.users.UpdateUsername(upd.NewUsername, info["id"]); err != nil {
		switch {
		case ent.IsNotFound(err):
			c.Error(exceptions.NoSuchUser.AddErr(err))
		case ent.IsValidationError(err):
			c.Error(exceptions.ValidError.AddErr(err))
		default:
			c.Error(exceptions.ServerError.AddErr(err))
		}
		return
	}

	c.Status(http.StatusOK)
}

// CheckUsername godoc
// @Summary Check username on exist
// @Description Status 200 if username not used or 403 if username already used
// @Tags User Get
// @Success 200 "name isn't used"
// @Failure 400 {object} exceptions.MyError
// @Failure 403 "name already used"
// @Failure 500 {object} exceptions.MyError
// @Router /user/upd/:username [get]
func (h Handler) checkUsername(c *gin.Context) {
	username := c.Param("username")

	if err := Valid.Struct(&dto.UsernameDTO{Username: username}); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}

	if h.users.UsernameExist(username) {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusOK)
	}
}
