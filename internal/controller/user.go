package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Handler) getMe(c *gin.Context, info *dao.Session) error {

	user, err := h.users.FindMe(info.ID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (h *Handler) getUserByUsername(c *gin.Context, username string) error {

	if err := h.v.Var(&username, "required,name"); err != nil {
		return err
	}

	user, err := h.users.FindUserByUsername(username)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (h *Handler) updateUser(c *gin.Context, upd *dto.UpdateUser, info *dao.Session) error {

	if err := h.users.UpdateUser(upd, info.ID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) updateEmail(c *gin.Context, upd *dto.UpdateEmail, info *dao.Session) error {

	user, err := h.users.FindUserByID(info.ID)

	if err != nil {
		return err
	} else if user.PasswordHash == nil {
		return errs.PasswordNotFound
	}

	if err = bcrypt.CompareHashAndPassword(*user.PasswordHash, []byte(upd.Password)); err != nil {
		return errs.PasswordNotFound.AddErr(err)
	}

	if err = h.users.UpdateEmail(upd.NewEmail, info.ID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) updatePassword(c *gin.Context, upd *dto.UpdatePassword, info *dao.Session) error {

	if ok, err := h.auth.EqualsPopCode(upd.Email, upd.Code); err != nil {
		return errs.ServerError.AddErr(err)
	} else if !ok {
		return errs.CodeError.AddErr(err)
	}

	if err := h.users.UpdatePassword([]byte(upd.NewPassword), info.ID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) updateUsername(c *gin.Context, upd *dto.UpdateName, info *dao.Session) error {

	if err := h.users.UpdateUsername(upd.NewName, info.ID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) checkUsername(c *gin.Context, name string) error {
	if err := h.v.Var(&name, "required,name"); err != nil {
		return err
	}

	if ok, err := h.users.UsernameExist(name); err != nil {
		return err
	} else if ok {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusOK)
	}
	return nil
}
