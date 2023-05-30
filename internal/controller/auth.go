package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
)

func (h *Handler) signInByPassword(c *gin.Context, auth *dto.EmailWithPassword) error {

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

	h.sess.SetNewCookie(customer.ID, c)

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) sendCodeToEmail(c *gin.Context, to *dto.Email) error {

	code := generateSecretCode()
	if err := h.auth.SetCodes(to.Email, code); err != nil {
		return errs.ServerError.AddErr(err)
	}

	if err := h.mail.SendEmail("Verify email for you-together account", code, cfg.Email.From, to.Email); err != nil {
		return errs.EmailError.AddErr(err)
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) signInByEmail(c *gin.Context, auth *dto.EmailWithCode) error {
	if oki, err := h.auth.EqualsPopCode(auth.Email, auth.Code); err != nil {
		return errs.ServerError.AddErr(err)
	} else if !oki {
		return errs.CodeError.AddErr(err)
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
	h.sess.SetNewCookie(customer.ID, c)

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) signOut(c *gin.Context) {
	h.sess.PopCookie(c)
	c.Status(http.StatusOK)
}

// generateSecretCode for email auth
func generateSecretCode() string {
	b := make([]rune, 5)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
