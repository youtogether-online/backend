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

func (h *Handler) getMe(c *gin.Context) {
	info, err := h.sess.GetSession(c)
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

func (h *Handler) getUserByUsername(c *gin.Context) {
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

func (h *Handler) updateUser(c *gin.Context) {
	upd, ok := bind.FillStruct[dto.UpdateUser](c)
	if !ok {
		return
	}

	info, err := h.sess.GetSession(c)
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

func (h *Handler) updateEmail(c *gin.Context) {
	upd, ok := bind.FillStruct[dto.UpdateEmail](c)
	if !ok {
		return
	}

	info, err := h.sess.GetSession(c)
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

func (h *Handler) updatePassword(c *gin.Context) {
	upd, ok := bind.FillStruct[dto.UpdatePassword](c)
	if !ok {
		return
	}
	info, err := h.sess.GetSession(c)
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

func (h *Handler) updateUsername(c *gin.Context) {
	upd, ok := bind.FillStruct[dto.UpdateName](c)
	if !ok {
		return
	}
	info, err := h.sess.GetSession(c)
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
// @Router /user/check-name/{name} [get]
func (h *Handler) checkUsername(c *gin.Context) {
	name := c.Param("name")
	if err := binding.Validator.ValidateStruct(&dto.Name{Name: name}); err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}

	if h.users.UsernameExist(name) {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusOK)
	}
}
