package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"github.com/wtkeqrf0/you_together/internal/middlewares/exceptions"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mail.v2"
	"math/rand"
	"net/http"
)

var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// signIn authentication by email and password. Returns a new session id in cookie, and deletes old, if exists
func (h Handler) signIn(c *gin.Context) {

	auth := &dto.SignInDTO{}
	if err := c.ShouldBindJSON(auth); err != nil {
		c.Error(exceptions.DataError.AddErr(err))
		return
	}

	if err := h.valid.Struct(auth); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}

	passwordHash, username, err := h.auth.AuthUserByEmail(auth.Email)
	if err != nil {
		passwordHash, username, err = h.auth.CreateUserWithPassword(auth.Email, auth.Password)

		if err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}
	}

	if err = bcrypt.CompareHashAndPassword(passwordHash, []byte(auth.Password)); err != nil {
		c.Error(exceptions.PasswordError.AddErr(err))
		return
	}

	id, _ := c.Cookie(cfg.Session.CookieName)
	if err = h.valid.Var(id, "uuid4"); err == nil {
		h.auth.DelSession(id)
	}
	id, _, err = h.sessions.GenerateSession(username, c.ClientIP(), auth.Device)
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	c.SetCookie(cfg.Session.CookieName, id, cfg.Session.DurationInSeconds,
		"/api", cfg.Listen.Host, false, true)

	c.Status(http.StatusOK)
}

// saveMail to users email and saves it
func (h Handler) saveMail(c *gin.Context) {
	to := &dto.EmailDTO{}
	if err := c.ShouldBindJSON(to); err != nil {
		c.Error(exceptions.DataError.AddErr(err))
		return
	}

	if err := h.valid.Struct(to); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}

	code := generateSecretCode()
	if err := h.auth.SetCodes(to.Email, code); err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	if err := sendEmail(to.Email, code); err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	c.Status(http.StatusOK)
}

// checkMail with saved code by email. Returns a new session id in cookie, and deletes old, if exists
func (h Handler) checkMail(c *gin.Context) {
	auth := &dto.EmailWithCodeDTO{}
	if err := c.ShouldBindJSON(auth); err != nil {
		c.Error(exceptions.DataError.AddErr(err))
		return
	}

	if err := h.valid.Struct(auth); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}

	if ok, err := h.auth.EqualsPopCode(auth.Email, auth.Code); !ok || err != nil {
		c.Error(exceptions.CodeError.AddErr(err))
		return
	}

	ok, username, err := h.auth.AuthUserWithInfo(auth.Email)

	if err != nil {
		username, err = h.auth.CreateUserByEmail(auth.Email)

		if err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}

	} else if !ok {
		err = h.auth.SetEmailVerified(auth.Email)

		if err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}
	}

	id, _ := c.Cookie(cfg.Session.CookieName)
	if err = h.valid.Var(id, "uuid4"); err == nil {
		h.auth.DelSession(id)
	}
	id, _, err = h.sessions.GenerateSession(username, c.ClientIP(), auth.Device)
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	c.SetCookie(cfg.Session.CookieName, id, cfg.Session.DurationInSeconds,
		"/api", cfg.Listen.Host, false, true)

	c.Status(http.StatusOK)
}

// signOut deletes the session id from the database, which makes it invalid
func (h Handler) signOut(c *gin.Context) {
	id, _ := c.Cookie(cfg.Session.CookieName)
	if err := h.valid.Var(id, "uuid4"); err == nil {
		h.auth.DelSession(id)
	}

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

// sendEmail to user with a secret code
func sendEmail(ToEmail string, code string) error {
	d := mail.NewDialer(cfg.Email.Host, cfg.Email.Port, cfg.Email.User, cfg.Email.Password)
	m := mail.NewMessage()
	m.SetHeader("From", cfg.Email.From)
	m.SetHeader("To", ToEmail)
	m.SetHeader("Subject", "Подтвердите ваш email")
	m.SetBody("text/plain", code)

	return d.DialAndSend(m)
}
