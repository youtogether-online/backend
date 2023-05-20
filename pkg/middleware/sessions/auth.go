package sessions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mssola/useragent"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/pkg/bind"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"net/http"
	"time"
)

type AuthService interface {
	SetSession(sessionId string, info dao.Session) error
	GetSession(sessionId string) (*dao.Session, error)
	ExpandExpireSession(sessionId string) (bool, error)
	IDExist(id int) bool
	DelKeys(keys ...string)
	AddSession(id int, sessions ...string) error
}

type Auth struct {
	auth AuthService
}

func NewAuth(auth AuthService) *Auth {
	return &Auth{auth: auth}
}

// ValidateSession and identify the user in redis. Returns true, if session was expanded.
func (a Auth) ValidateSession(sessionId string) (*dao.Session, bool, error) {
	if sessionId == "" {
		return nil, false, fmt.Errorf("session id is not found")
	}

	info, err := a.auth.GetSession(sessionId)
	if err != nil {
		return nil, false, err
	}

	if !a.auth.IDExist(info.ID) {
		return nil, false, fmt.Errorf("user does not exists")
	}

	ok, err := a.auth.ExpandExpireSession(sessionId)
	if err != nil {
		return nil, false, fmt.Errorf("session does not exist: %v", err)
	}

	return info, ok, nil
}

// GenerateSession and save it to redis
func (a Auth) GenerateSession(id int, ip, userAgent string) (string, error) {
	ua := useragent.New(userAgent)
	name, ver := ua.Browser()
	sessionId := uuid.NewString()

	err := a.auth.AddSession(id, sessionId)
	if err != nil {
		return "", err
	}

	return sessionId, a.auth.SetSession(sessionId, dao.Session{
		ID:      id,
		IP:      ip,
		Device:  ua.OS(),
		Browser: name + ver,
		Updated: time.Now().Unix(),
	})
}

func (a Auth) GetSession(c *gin.Context) (*dao.Session, error) {
	get, ok := c.Get("user_info")
	if !ok {
		return nil, fmt.Errorf("session not found in context")
	}

	res, ok := get.(*dao.Session)
	if !ok {
		return nil, fmt.Errorf("cannot parse session")
	}
	return res, nil
}

func (a Auth) SetNewCookie(id int, userAgent string, c *gin.Context) {
	a.PopCookie(c)

	session, err := a.GenerateSession(id, c.ClientIP(), userAgent)
	if err != nil {
		c.Error(errs.ServerError.AddErr(err))
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(cfg.Session.CookieName, session, int(cfg.Session.Duration.Seconds()),
		cfg.Session.CookiePath, cfg.Listen.Host, true, true)
}

// PopCookie from cookie storage only if equals to uuid4
func (a Auth) PopCookie(c *gin.Context) {
	session, _ := c.Cookie(cfg.Session.CookieName)
	if ok := bind.UUID4.MatchString(session); ok {
		a.auth.DelKeys(session)
	}
}
