package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"github.com/wtkeqrf0/you_together/internal/middleware/exceptions"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mail.v2"
	"math/rand"
	"net/http"
)

var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// signInWithPassword authentication by email and password. Returns a new session id in cookie, and deletes old, if exists

// SignInWithPassword godoc
// @Summary Sign in by password
// @Description Compare the user's password with an existing user's password. If it matches, create session of the user. If the user does not exist, create a new user
// @Param user_info_with_password body dto.SignInDTO true "User's email, password and device"
// @Accept application/json
// @Produce application/json
// @Tags Authorization
// @Success 200 "return session id in cookie"
// @Failure 400 {object} exceptions.MyError
// @Failure 404 {object} exceptions.MyError "The password is not registered for this account"
// @Failure 500 {object} exceptions.MyError
// @Router /auth/sign-in-with-password [post]
func (h Handler) signInWithPassword(c *gin.Context) {
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
	} else if len(passwordHash) < 1 {
		c.Error(exceptions.PasswordNotFound)
		return
	}

	if err = bcrypt.CompareHashAndPassword(passwordHash, []byte(auth.Password)); err != nil {
		c.Error(exceptions.PasswordError.AddErr(err))
		return
	}
	c.SetSameSite(http.SameSiteNoneMode)

	id, _ := c.Cookie(cfg.Session.CookieName)
	if err = h.valid.Var(id, "uuid4"); err == nil {
		h.auth.DelKeys(id)
	}
	id, _, err = h.sessions.GenerateSession(username, c.ClientIP(), auth.Device)

	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	c.SetCookie(cfg.Session.CookieName, id, cfg.Session.DurationInSeconds,
		"/api", cfg.Listen.Host, true, true)

	c.Status(http.StatusOK)
}

// signInSendCode to users email and saves it

// SignInSendCode godoc
// @Summary Send code to the user's email
// @Description Send a secret 5-digit code to the specified email
// @Param email body dto.EmailDTO true "User's email"
// @Accept application/json
// @Produce application/json
// @Tags Authorization
// @Success 200
// @Failure 400 {object} exceptions.MyError
// @Failure 500 {object} exceptions.MyError
// @Router /auth/sign-in-send-code [post]
func (h Handler) signInSendCode(c *gin.Context) {
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

	go sendEmail(to.Email, code)

	c.Status(http.StatusOK)
}

// signInCheckCode with saved code by email. Returns a new session id in cookie, and deletes old, if exists

// SignInCheckCode godoc
// @Summary Sign in by email
// @Description Compare the secret code with the previously sent codes. If at least one matches, create session of the user. If the user does not exist, create a new user
// @Param user_info body dto.EmailWithCodeDTO true "User's email, secret code and device"
// @Accept application/json
// @Produce application/json
// @Tags Authorization
// @Success 200 "return session id in cookie"
// @Failure 400 {object} exceptions.MyError
// @Failure 500 {object} exceptions.MyError
// @Router /auth/sign-in-check-code [post]
func (h Handler) signInCheckCode(c *gin.Context) {
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
	c.SetSameSite(http.SameSiteNoneMode)

	id, _ := c.Cookie(cfg.Session.CookieName)
	if err = h.valid.Var(id, "uuid4"); err == nil {
		h.auth.DelKeys(id)
	}
	id, _, err = h.sessions.GenerateSession(username, c.ClientIP(), auth.Device)
	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	c.SetCookie(cfg.Session.CookieName, id, cfg.Session.DurationInSeconds,
		"/api", cfg.Listen.Host, true, true)

	c.Status(http.StatusOK)
}

// signOut deletes the session id from the database, which makes it invalid

// SignOut godoc
// @Summary Delete user's session
// @Description Make the user's session invalid
// @Tags Authorization
// @Success 200
// @Router /auth/sign-out [post]
func (h Handler) signOut(c *gin.Context) {
	c.SetSameSite(http.SameSiteNoneMode)

	id, _ := c.Cookie(cfg.Session.CookieName)
	if err := h.valid.Var(id, "uuid4"); err == nil {
		h.auth.DelKeys(id)
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
func sendEmail(ToEmail string, code string) {
	d := mail.NewDialer(cfg.Email.Host, cfg.Email.Port, cfg.Email.User, cfg.Email.Password)
	m := mail.NewMessage()
	m.SetHeader("From", cfg.Email.From)
	m.SetHeader("To", ToEmail)
	m.SetHeader("Subject", "Подтвердите ваш email")
	m.SetBody("text/plain", code)

	d.DialAndSend(m)
}
