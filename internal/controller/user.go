package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/bind"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// GetMe godoc
// @Summary Get detail data about the user by session
// @Description Returns detail information about me (session required)
// @Tags Sessions
// @Success 200 {object} dao.Me
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 404 {object} errs.MyError "User doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /session [get]
func (h Handler) getMe(c *gin.Context) {
	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	}

	user, err := h.users.FindMe(info.ID)
	if err != nil {
		if _, ok := err.(errs.MyError); ok {
			c.Error(errs.NoSuchUser)
		} else {
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByUsername godoc
// @Summary Get main data about the user
// @Description Returns main information about the user
// @Tags User Get
// @Param username path string true "Name of the user"
// @Success 200 {object} dao.User "main info"
// @Failure 400 {object} errs.MyError "Param is not valid"
// @Failure 404 {object} errs.MyError "User doesn't exist"
// @Failure 500 {object} errs.MyError
// @Router /user/{username} [get]
func (h Handler) getUserByUsername(c *gin.Context) {
	username := c.Param("username")
	if err := binding.Validator.ValidateStruct(&dto.Name{Name: username}); err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}

	user, err := h.users.FindUserByUsername(username)

	if err != nil {
		if err == errs.NoSuchUser {
			c.Error(errs.NoSuchUser)
		} else {
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update user's data
// @Description Change user's main information (session required)
// @Param UpdateUser body dto.UpdateUser true "New user data"
// @Tags User Update
// @Success 200
// @Failure 400 {object} errs.MyError "Data is not valid"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /user [patch]
func (h Handler) updateUser(c *gin.Context) {
	upd, ok := bind.FillStruct[dto.UpdateUser](c)
	if !ok {
		return
	}

	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	}

	if err = h.users.UpdateUser(upd, info.ID); err != nil {
		if ent.IsNotFound(err) {
			c.Error(errs.NoSuchUser.AddErr(err))
		} else {
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	}

	c.Status(http.StatusOK)
}

// UpdateEmail godoc
// @Summary Update user's email
// @Description Change user's email by password (session required)
// @Param UpdateEmail body dto.UpdateEmail true "user's password and new email"
// @Tags User Update
// @Success 200
// @Failure 400 {object} errs.MyError "Data is not valid"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /user/email [patch]
func (h Handler) updateEmail(c *gin.Context) {
	upd, ok := bind.FillStruct[dto.UpdateEmail](c)
	if !ok {
		return
	}

	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	}

	user, err := h.users.FindUserByID(info.ID)

	if err != nil {
		if _, ok := err.(errs.MyError); ok {
			c.Error(errs.NoSuchUser)
		} else {
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	} else if user.PasswordHash == nil {
		c.Error(errs.PasswordNotFound)
		return
	}

	if err = bcrypt.CompareHashAndPassword(*user.PasswordHash, []byte(upd.Password)); err != nil {
		c.Error(errs.PasswordError.AddErr(err))
		return
	}

	if err = h.users.UpdateEmail(upd.NewEmail, info.ID); err != nil {
		if ent.IsValidationError(err) {
			c.Error(errs.AlreadyExist.AddErr(err))
		} else {
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	}

	c.Status(http.StatusOK)
}

// UpdatePassword godoc
// @Summary Update user's password
// @Description Change user's password by email (session required)
// @Param UpdatePassword body dto.UpdatePassword true "user's email, code and new password"
// @Tags User Update
// @Success 200
// @Failure 400 {object} errs.MyError "Data is not valid"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /user/password [patch]
func (h Handler) updatePassword(c *gin.Context) {
	upd, ok := bind.FillStruct[dto.UpdatePassword](c)
	if !ok {
		return
	}
	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	}

	if ok, err := h.auth.EqualsPopCode(upd.Email, upd.Code); !ok {
		c.Error(errs.CodeError.AddErr(err))
		return
	} else if err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	}

	if err = h.users.UpdatePassword(upd.NewPassword, info.ID); err != nil {
		if ent.IsNotFound(err) {
			c.Error(errs.NoSuchUser.AddErr(err))
		} else {
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	}

	c.Status(http.StatusOK)
}

// UpdateUsername godoc
// @Summary Update user's name
// @Description Change user's name (session required)
// @Tags User Update
// @Success 200
// @Failure 400 {object} errs.MyError "Data is not valid"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /user/name [patch]
func (h Handler) updateUsername(c *gin.Context) {
	upd, ok := bind.FillStruct[dto.UpdateName](c)
	if !ok {
		return
	}
	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	}

	if err = h.users.UpdateUsername(upd.NewName, info.ID); err != nil {
		switch {
		case ent.IsNotFound(err):
			c.Error(errs.NoSuchUser.AddErr(err))
		case ent.IsValidationError(err):
			c.Error(errs.ValidError.AddErr(err))
		default:
			c.Error(errs.ServerError.AddErr(err))
		}
		return
	}

	c.Status(http.StatusOK)
}

// CheckUsername godoc
// @Summary Check username on exist
// @Description Status 200 if username not used or 403 if username is already used
// @Tags User Get
// @Param username path string true "Name of the user"
// @Success 200 "name isn't used"
// @Failure 400 {object} errs.MyError "Username is not valid"
// @Failure 403 "name already used"
// @Failure 500 {object} errs.MyError
// @Router /user/check-name/{username} [get]
func (h Handler) checkUsername(c *gin.Context) {
	username := c.Param("username")
	if err := binding.Validator.ValidateStruct(&dto.Name{Name: username}); err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}

	if h.users.UsernameExist(username) {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusOK)
	}
}
