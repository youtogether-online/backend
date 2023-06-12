package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func (h *Handler) signInByPassword(c *gin.Context, auth dto.EmailWithPassword) error {

	customer, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		customer, err = h.auth.CreateUserWithPassword(auth.Email, []byte(auth.Password), auth.Language)

		if err != nil {
			return err
		}
	} else if customer.PasswordHash == nil {
		return errs.PasswordNotFound
	}

	if err = bcrypt.CompareHashAndPassword(*customer.PasswordHash, []byte(auth.Password)); err != nil {
		return errs.PasswordError.AddErr(err)
	}

	err = h.sess.SetNewCookie(customer.ID, c)
	if err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) sendCodeToEmail(c *gin.Context, to dto.Email) error {

	code := h.sess.GenerateSecretCode()
	if err := h.auth.SetCodes(to.Email, code); err != nil {
		return err
	}

	if err := h.mail.SendEmail("Verify email for you-together account", code, "", to.Email); err != nil {
		return errs.EmailError.AddErr(err)
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) signInByEmail(c *gin.Context, auth dto.EmailWithCode) error {
	if oki, err := h.auth.EqualsPopCode(auth.Email, auth.Code); err != nil {
		return errs.ServerError.AddErr(err)
	} else if !oki {
		return errs.MailCodeError.AddErr(err)
	}

	customer, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		customer, err = h.auth.CreateUserByEmail(auth.Email, auth.Language)

		if err != nil {
			return err
		}
	} else if !customer.IsEmailVerified {
		err = h.auth.SetEmailVerified(auth.Email)

		if err != nil {
			return err
		}
	}
	err = h.sess.SetNewCookie(customer.ID, c)
	if err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) signOut(c *gin.Context, info *dao.Session) error {
	h.auth.DelKeys(strconv.Itoa(info.ID))
	c.Status(http.StatusOK)
	return nil
}
