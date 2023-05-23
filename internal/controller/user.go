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
	upd := bind.FillStructJSON[dto.UpdateUser](c)
	if upd == (dto.UpdateUser{}) {
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
	upd := bind.FillStructJSON[dto.UpdateEmail](c)
	if upd == (dto.UpdateEmail{}) {
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
	upd := bind.FillStructJSON[dto.UpdatePassword](c)
	if upd == (dto.UpdatePassword{}) {
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
	upd := bind.FillStructJSON[dto.UpdateName](c)
	if upd == (dto.UpdateName{}) {
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

func (h *Handler) checkUsername(c *gin.Context) {
	name := c.Param("name")
	if err := binding.Validator.ValidateStruct(&dto.Name{Name: name}); err != nil {
		c.Error(errs.ValidError.AddErr(err))
		return
	}

	if ok, err := h.users.UsernameExist(name); err != nil {
		c.Error(errs.ServerError.AddErr(err))
	} else if ok {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusOK)
	}
}
