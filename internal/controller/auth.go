package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"net/http"
)

func (h *Handler) signInByPassword(c *gin.Context, auth dto.EmailWithPassword) error {
	auth.Language = h.auth.FormatLanguage(auth.Language)

	customer, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		customer, err = h.auth.CreateUserWithPassword(auth)

		if err != nil {
			return err
		}
	} else if customer.PasswordHash == nil {
		return errs.PasswordNotFound
	}

	if err = h.auth.CompareHashAndPassword(*customer.PasswordHash, []byte(auth.Password)); err != nil {
		return errs.InvalidPassword.AddErr(err)
	}

	err = h.sess.SetNewCookie(customer.ID, c)
	if err != nil {
		return err
	}

	c.Status(http.StatusCreated)
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
	auth.Language = h.auth.FormatLanguage(auth.Language)

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

	c.Status(http.StatusCreated)
	return nil
}

func (h *Handler) signOut(c *gin.Context, cookieName string) error {
	session, _ := c.Cookie(cookieName)
	info, _, err := h.sess.ValidateSession(session)
	if err != nil {
		return errs.UnAuthorized.AddErr(err)
	}

	h.auth.DelKeys(session)

	user, err := h.user.FindUserByID(info.ID)
	if err != nil {
		return err
	}

	for i, v := range user.Sessions {
		if v == session {
			if err = user.Update().SetSessions(append(user.Sessions[:i], user.Sessions[i+1:]...)).
				Exec(context.Background()); err != nil {
				return err
			}
		}
	}

	c.Status(http.StatusOK)
	return nil
}
