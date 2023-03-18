package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"github.com/wtkeqrf0/you_together/pkg/middlewares/exceptions"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mail.v2"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// sign-in authentication by email and password. Returns a pair of tokens in cookies
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

	customer, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		customer, err = h.auth.CreateUser(auth.Email, auth.Password)

		if err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}
	}

	if err = bcrypt.CompareHashAndPassword(*customer.PasswordHash, []byte(auth.Password)); err != nil {
		c.Error(exceptions.PasswordError.AddErr(err))
		return
	}

	t, err := h.auth.GenerateJWT(float64(customer.ID))
	if err != nil {
		c.Error(exceptions.ServerError)
		return
	}

	c.SetCookie(cfg.Token.AccessName, t.AT, cfg.Token.AccessDurationInSeconds,
		"/api", cfg.Listen.Host, false, true)

	c.SetCookie(cfg.Token.RefreshName, t.RT, cfg.Token.RefreshDurationInSeconds,
		"/api", cfg.Listen.Host, false, true)

}

// sendSecretCode to users email and save it
func (h Handler) sendSecretCode(c *gin.Context) {
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
	if err := h.redis.SetCodes(to.Email, code); err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	if err := sendEmail(to.Email, code); err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	c.Status(http.StatusOK)
}

// compareSecretCode with saved code by email. Returns a pair of tokens in cookies
func (h Handler) compareSecretCode(c *gin.Context) {
	auth := &dto.EmailWithCodeDTO{}
	if err := c.ShouldBindJSON(auth); err != nil {
		c.Error(exceptions.DataError.AddErr(err))
		return
	}

	if err := h.valid.Struct(auth); err != nil {
		c.Error(exceptions.ValidError.AddErr(err))
		return
	}

	if !h.redis.ContainsPopCode(auth.Email, auth.Code) {
		c.Error(exceptions.CodeError)
		return
	}

	customer, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		customer, err = h.auth.CreateUser(auth.Email, "")

		if err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}
	}

	t, err := h.auth.GenerateJWT(float64(customer.ID))
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}
	c.SetCookie(cfg.Token.AccessName, t.AT, cfg.Token.AccessDurationInSeconds,
		"/api", cfg.Listen.Host, false, true)

	c.SetCookie(cfg.Token.RefreshName, t.RT, cfg.Token.RefreshDurationInSeconds,
		"/api", cfg.Listen.Host, false, true)
}

func (h Handler) logout(c *gin.Context) {
	at, _ := c.Cookie(cfg.Token.AccessName)
	atClaims, err := h.auth.ValidateJWT(at)
	if err == nil {
		if err = h.redis.SetVariable(at, atClaims["sub"].(float64), time.Now().
			Sub(time.UnixMilli(int64(atClaims["exp"].(float64))))); err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}
	}

	rt, _ := c.Cookie(cfg.Token.RefreshName)
	rtClaims, err := h.auth.ValidateJWT(rt)
	if err == nil {
		if err = h.redis.SetVariable(at, rtClaims["sub"].(float64), time.Now().
			Sub(time.UnixMilli(int64(rtClaims["exp"].(float64))))); err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}
	}

	c.Status(http.StatusOK)
}

// generateSecretPassword for email auth
func generateSecretCode() string {
	all := []rune(chars)
	var b strings.Builder
	for i := 0; i < 5; i++ {
		b.WriteRune(all[rand.Intn(len(all))])
	}
	return b.String()
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
