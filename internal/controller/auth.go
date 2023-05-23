package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/bind"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
)

func (h *Handler) signInByPassword(c *gin.Context) {
	auth := bind.FillStructJSON[dto.EmailWithPassword](c)
	if auth == (dto.EmailWithPassword{}) {
		return
	}

	customer, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		customer, err = h.auth.CreateUserWithPassword(auth)

		if err != nil {
			c.Error(errs.ServerError.AddErr(err))
			return
		}
	} else if customer.PasswordHash == nil {
		c.Error(errs.PasswordNotFound)
		return
	}

	if err = bcrypt.CompareHashAndPassword(*customer.PasswordHash, []byte(auth.Password)); err != nil {
		c.Error(errs.PasswordError.AddErr(err))
		return
	}

	h.sess.SetNewCookie(customer.ID, c)

	c.Status(http.StatusOK)
}

func (h *Handler) sendCodeToEmail(c *gin.Context) {
	to := bind.FillStructJSON[dto.Email](c)
	if to == (dto.Email{}) {
		return
	}

	code := generateSecretCode()
	if err := h.auth.SetCodes(to.Email, code); err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	}

	if err := h.mail.SendEmail("Verify email for you-together account", code, cfg.Email.From, to.Email); err != nil {
		c.Error(errs.EmailError.AddErr(err))
	}

	c.Status(http.StatusOK)
}

func (h *Handler) signInByEmail(c *gin.Context) {
	auth := bind.FillStructJSON[dto.EmailWithCode](c)
	if auth == (dto.EmailWithCode{}) {
		return
	}

	if oki, err := h.auth.EqualsPopCode(auth.Email, auth.Code); err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	} else if !oki {
		c.Error(errs.CodeError.AddErr(err))
		return
	}

	customer, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		customer, err = h.auth.CreateUserByEmail(auth)

		if err != nil {
			c.Error(errs.ServerError.AddErr(err))
			return
		}
	} else if !customer.IsEmailVerified {
		err = h.auth.SetEmailVerified(auth.Email)

		if err != nil {
			c.Error(errs.ServerError.AddErr(err))
			return
		}
	}
	h.sess.SetNewCookie(customer.ID, c)

	c.Status(http.StatusOK)
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
