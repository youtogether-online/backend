package authorization

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mssola/useragent"
	"github.com/wtkeqrf0/you_together/internal/middleware/exceptions"
	"net/http"
	"regexp"
	"time"
)

type AuthService interface {
	SetSession(sessionId string, info map[string]string) error
	GetSession(sessionId string) (map[string]string, error)
	ExpandExpireSession(sessionId string) (bool, error)
	IDExist(id string) bool
	DelKeys(keys ...string)
	AddSession(id string, sessions ...string) error
}

type Auth struct {
	auth AuthService
}

func NewAuth(auth AuthService) *Auth {
	return &Auth{auth: auth}
}

const userAgent string = "User-Agent"

// ValidateSession validates the session and identifies the user in DB. Returns an error in case of unsuccessful validation
func (a Auth) ValidateSession(sessionId string) (map[string]string, bool, error) {
	if sessionId == "" {
		return nil, false, fmt.Errorf("session id is not found")
	}

	info, err := a.auth.GetSession(sessionId)
	if err != nil {
		return nil, false, err
	}

	if !a.auth.IDExist(info["id"]) {
		return nil, false, fmt.Errorf("user does not exists")
	}

	ok, err := a.auth.ExpandExpireSession(sessionId)
	if err != nil {
		return nil, false, fmt.Errorf("session does not exist: %v", err)
	}

	return info, ok, nil
}

// GenerateSession generates a new session
func (a Auth) GenerateSession(id, ip, userAgent string) (string, error) {
	ua := useragent.New(userAgent)
	name, ver := ua.Browser()
	sessionId := uuid.New().String()

	err := a.auth.AddSession(id, sessionId)
	if err != nil {
		return "", err
	}

	return sessionId, a.auth.SetSession(sessionId, map[string]string{
		"id":      id,
		"ip":      ip,
		"device":  ua.OS(),
		"browser": name + ver,
		"created": time.Now().String(),
	})
}

func (a Auth) GetSession(c *gin.Context) (map[string]string, error) {
	get, ok := c.Get("user_info")
	if !ok {
		return nil, fmt.Errorf("the user is not logged in")
	}

	res, ok := get.(map[string]string)
	if !ok {
		return nil, fmt.Errorf("the user is not logged in")
	}
	return res, nil
}

func (a Auth) SetNewCookie(id string, c *gin.Context) {
	a.PopCookie(c)

	session, err := a.GenerateSession(
		id, c.ClientIP(), c.GetHeader(userAgent),
	)

	if err != nil {
		c.Error(exceptions.ServerError.AddErr(err))
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(cfg.Session.CookieName, session, cfg.Session.DurationInSeconds,
		cfg.Session.CookiePath, cfg.Listen.Host, true, true)
}

func (a Auth) PopCookie(c *gin.Context) {
	session, _ := c.Cookie(cfg.Session.CookieName)
	if ok, _ := regexp.MatchString(cfg.Regexp.UUID4, session); ok {
		a.auth.DelKeys(session)
	}
}
