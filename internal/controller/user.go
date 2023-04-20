package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/internal/middleware/exceptions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// GetUserByUsername godoc
// @Summary Get type of the user (NOT WORKING)
// @Description Returns type of object (NOT WORKING)
// @Tags Get
// @Param name path string true "Name of something"
// @Success 200 {object} dto.UserDTO "string type with object"
// @Failure 400 {object} exceptions.MyError
// @Failure 404 {object} exceptions.MyError "User doesn't exist"
// @Failure 500 {object} exceptions.MyError
// @Router /{name} [get]
func (h Handler) getTypeByName(c *gin.Context) {
	name := c.Param("name")
	if err := valid.Struct(&dto.NameDTO{Name: name}); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}

	if user, err := h.users.FindUserByUsername(name); err == nil {
		c.JSON(http.StatusOK, dto.TypeDTO{
			Type:   "user",
			Object: user,
		})
	} // else if room, err := h.rooms
}

// GetMe godoc
// @Summary Get detail information about user by session
// @Description Return detail information about the user (cookie required)
// @Tags Sessions
// @Success 200 {object} dto.MyUserDTO
// @Failure 401 {object} exceptions.MyError "User isn't logged in"
// @Failure 404 {object} exceptions.MyError "User doesn't exist"
// @Failure 500 {object} exceptions.MyError
// @Router /session [get]
func (h Handler) getMe(c *gin.Context) {
	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	user, err := h.users.FindMe(info.ID)
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
// @Tags User Get
// @Param username path string true "Name of the user"
// @Success 200 {object} dto.UserDTO "main info, without cookie or dto.MyUserDTO"
// @Failure 400 {object} exceptions.MyError
// @Failure 404 {object} exceptions.MyError "User doesn't exist"
// @Failure 500 {object} exceptions.MyError
// @Router /user/{username} [get]
func (h Handler) getUserByUsername(c *gin.Context) {
	username := c.Param("username")
	if err := valid.Struct(&dto.NameDTO{Name: username}); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}

	var (
		user *dto.UserDTO
		me   *dto.MyUserDTO
		info *dto.Session
		err  error
	)

	if info, err = h.sessions.GetSession(c); err != nil {
		user, err = h.users.FindUserByUsername(username)

	} else {
		me, err = h.users.FindMe(info.ID)

		if err != nil && me.Name != username {
			user, err = h.users.FindUserByUsername(username)
		}
	}

	if err != nil {
		if exceptions.NoSuchUser == err {
			c.Error(exceptions.NoSuchUser)
			fmt.Println(err)
		} else {
			c.Error(exceptions.ServerError.AddErr(err))
		}
		return
	}

	if user != nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusOK, me)
	}
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
// @Router /user [patch]
func (h Handler) updateUser(c *gin.Context) {
	upd, ok := fillStruct[dto.UpdateUserDTO](c)
	if !ok {
		return
	}

	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	if err = h.users.UpdateUser(upd, info.ID); err != nil {
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
// @Router /user/email [patch]
func (h Handler) updateEmail(c *gin.Context) {
	upd, ok := fillStruct[dto.UpdateEmailDTO](c)
	if !ok {
		return
	}

	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	user, err := h.users.FindUserByID(info.ID)

	if err != nil {
		if _, ok := err.(exceptions.MyError); ok {
			c.Error(exceptions.NoSuchUser)
		} else {
			c.Error(exceptions.ServerError.AddErr(err))
		}
		return
	} else if user.PasswordHash == nil {
		c.Error(exceptions.PasswordNotFound)
		return
	}

	if err = bcrypt.CompareHashAndPassword(*user.PasswordHash, []byte(upd.Password)); err != nil {
		c.Error(exceptions.PasswordError.AddErr(err))
		return
	}

	if err = h.users.UpdateEmail(upd.NewEmail, info.ID); err != nil {
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
// @Router /user/password [patch]
func (h Handler) updatePassword(c *gin.Context) {
	upd, ok := fillStruct[dto.UpdatePasswordDTO](c)
	if !ok {
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

	if err = h.users.UpdatePassword(upd.NewPassword, info.ID); err != nil {
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
// @Router /user/name [patch]
func (h Handler) updateUsername(c *gin.Context) {
	upd, ok := fillStruct[dto.UpdateNameDTO](c)
	if !ok {
		return
	}
	info, err := h.sessions.GetSession(c)
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	if err = h.users.UpdateUsername(upd.NewUsername, info.ID); err != nil {
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
// @Param username path string true "Name of the user"
// @Success 200 "name isn't used"
// @Failure 400 {object} exceptions.MyError
// @Failure 403 "name already used"
// @Failure 500 {object} exceptions.MyError
// @Router /user/check-name/{username} [get]
func (h Handler) checkUsername(c *gin.Context) {
	username := c.Param("username")

	if err := valid.Struct(&dto.NameDTO{Name: username}); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}

	if h.users.UsernameExist(username) {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusOK)
	}
}
