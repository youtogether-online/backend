package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"golang.org/x/crypto/bcrypt"
	"image"
	"net/http"
)

func (h *Handler) getMe(c *gin.Context, info *dao.Session) error {

	user, err := h.user.FindMe(info.ID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (h *Handler) getUserByUsername(c *gin.Context, name dto.NameParam) error {

	user, err := h.user.FindUserByUsername(name.Name)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (h *Handler) updateUser(c *gin.Context, upd dto.UpdateUser, info *dao.Session) error {

	if err := h.user.UpdateUser(upd, info.ID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) updateImage(c *gin.Context, upd dto.UpdateImage, info *dao.Session) error {
	if upd.Image == nil {
		return errs.UnsupportedImageType
	}

	img, err := upd.Image.Open()
	if err != nil {
		return err
	}

	var imgConfig image.Config
	imgConfig, _, err = image.DecodeConfig(img)
	if err != nil {
		return errs.UnsupportedImageType.AddErr(err)
	}

	if imgConfig.Width > 3000 || imgConfig.Height > 3000 {
		return errs.ImageTooLarge
	}

	imageName, err := h.sess.GenerateFileName(c, upd.Image)
	if err != nil {
		return err
	}

	if err = h.user.UpdateImage(imageName, info.ID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) updateEmail(c *gin.Context, upd dto.UpdateEmail, info *dao.Session) error {

	user, err := h.user.FindUserByID(info.ID)

	if err != nil {
		return err
	} else if user.PasswordHash == nil {
		return errs.PasswordNotFound
	}

	if err = bcrypt.CompareHashAndPassword(*user.PasswordHash, []byte(upd.Password)); err != nil {
		return errs.InvalidPassword.AddErr(err)
	}

	if err = h.user.UpdateEmail(upd.NewEmail, info.ID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) updatePassword(c *gin.Context, upd dto.UpdatePassword, info *dao.Session) error {

	user, err := h.user.FindUserByID(info.ID)
	if err != nil {
		return err
	}
	if user.PasswordHash != nil {
		if err = h.auth.CompareHashAndPassword(*user.PasswordHash, []byte(upd.OldPassword)); err != nil {
			return errs.InvalidPassword.AddErr(err)
		}
	}

	if err = h.user.UpdatePassword([]byte(upd.NewPassword), info.ID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) updateUsername(c *gin.Context, upd dto.UpdateName, info *dao.Session) error {

	if err := h.user.UpdateUsername(upd.NewName, info.ID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) checkUsername(c *gin.Context, name dto.NameParam) error {
	if ok, err := h.user.UsernameExist(name.Name); err != nil {
		return err
	} else if ok {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusOK)
	}
	return nil
}
