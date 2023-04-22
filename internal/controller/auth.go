package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/internal/middleware/exceptions"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mail.v2"
	"math/rand"
	"net/http"
)

const acceptLanguage string = "Accept-Language"

// SignInWithPassword godoc
// @Summary Sign in by password
// @Description Compare the user's password with an existing user's password. If it matches, create session of the user. If the user does not exist, create a new user
// @Param EmailWithPasswordDTO body dto.EmailWithPasswordDTO true "User's email, password"
// @Param Accept-Language header string false "User's language" Enums(EN,RU)
// @Tags Authorization
// @Success 200 "user's session"
// @Failure 400 {object} exceptions.MyError
// @Failure 404 {object} exceptions.MyError "Password is not registered for this account"
// @Failure 500 {object} exceptions.MyError
// @Router /auth/password/sign-in [post]
func (h Handler) signInByPassword(c *gin.Context) {
	auth, ok := fillStruct[dto.EmailWithPasswordDTO](c)
	if !ok {
		return
	}

	auth.Language = validLanguage(c.GetHeader(acceptLanguage))

	customer, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		customer, err = h.auth.CreateUserWithPassword(auth)

		if err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}
	} else if customer.PasswordHash == nil {
		c.Error(exceptions.PasswordNotFound)
		return
	}

	if err = bcrypt.CompareHashAndPassword(*customer.PasswordHash, []byte(auth.Password)); err != nil {
		c.Error(exceptions.PasswordError.AddErr(err))
		return
	}
	h.sessions.SetNewCookie(customer.ID, c)

	c.Status(http.StatusOK)
}

// SignInSendCode godoc
// @Summary Send code to the user's email
// @Description Send a secret 5-digit code to the specified email
// @Param email body dto.EmailDTO true "User's email"
// @Tags Email
// @Success 200
// @Failure 400 {object} exceptions.MyError
// @Failure 500 {object} exceptions.MyError
// @Router /email/send-code [post]
func (h Handler) sendEmailCode(c *gin.Context) {
	to, ok := fillStruct[dto.EmailDTO](c)
	if !ok {
		return
	}

	code := generateSecretCode()
	if err := h.auth.SetCodes(to.Email, code); err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	go func() {
		if err := sendEmail(to.Email, code); err != nil {
			c.Error(exceptions.EmailError.AddErr(err))
		}
	}()

	c.Status(http.StatusOK)
}

// SignInCheckCode godoc
// @Summary Sign in by email
// @Description Compare the secret code with the previously sent codes. If at least one matches, create session of the user. If the user does not exist, create a new user
// @Param user_info body dto.EmailWithCodeDTO true "User's email, secret code and device"
// @Param Accept-Language header string false "User's language" Enums(EN,RU)
// @Tags Authorization
// @Success 200 "user's session"
// @Failure 400 {object} exceptions.MyError
// @Failure 500 {object} exceptions.MyError
// @Router /auth/email/sign-in [post]
func (h Handler) signInByEmail(c *gin.Context) {
	auth, ok := fillStruct[dto.EmailWithCodeDTO](c)
	if !ok {
		return
	}

	auth.Language = validLanguage(c.GetHeader(acceptLanguage))

	if ok, err := h.auth.EqualsPopCode(auth.Email, auth.Code); !ok || err != nil {
		c.Error(exceptions.CodeError.AddErr(err))
		return
	}

	customer, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		customer, err = h.auth.CreateUserByEmail(auth)

		if err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}

	} else if !customer.IsEmailVerified {
		err = h.auth.SetEmailVerified(auth.Email)

		if err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}
	}
	h.sessions.SetNewCookie(customer.ID, c)

	c.Status(http.StatusOK)
}

// SignOut godoc
// @Summary Delete user's session
// @Description Make the user's session invalid (can accept cookie)
// @Tags Sessions
// @Success 200
// @Router /session [delete]
func (h Handler) signOut(c *gin.Context) {
	h.sessions.PopCookie(c)
	c.Status(http.StatusOK)
}

var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

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

func validLanguage(language string) string {
	switch language {
	case "RU":
		return "RU"
	default:
		return "EN"
	}
}
