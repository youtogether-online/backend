package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mssola/useragent"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type AuthService interface {
	SetSession(sessionId string, info dao.Session) error
	GetSession(sessionId string) (*dao.Session, error)
	ExpandExpireSession(sessionId string) (bool, error)
	IDExist(id int) (bool, error)
	AddSession(id int, sessions ...string) error
}

type Auth struct {
	auth AuthService
	cfg  *conf.Config
}

func NewAuth(auth AuthService, cfg *conf.Config) *Auth {
	return &Auth{auth: auth, cfg: cfg}
}

// ValidateSession and identify the user in redis. Returns true, if session was expanded.
func (a *Auth) ValidateSession(sessionId string) (info *dao.Session, ok bool, err error) {
	if sessionId == "" {
		return nil, false, fmt.Errorf("session id is not found")
	}

	if info, err = a.auth.GetSession(sessionId); err != nil {
		return nil, false, err
	}

	if ok, err = a.auth.IDExist(info.ID); err != nil {
		return nil, false, err
	} else if !ok {
		return nil, false, fmt.Errorf("user does not exists")
	}

	ok, err = a.auth.ExpandExpireSession(sessionId)
	if err != nil {
		return nil, false, fmt.Errorf("session does not exist: %v", err)
	}

	return
}

// generateSession and save it to redis
func (a *Auth) generateSession(id int, ip, userAgent string) (string, error) {
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

func (a *Auth) getSession(c *gin.Context) *dao.Session {
	get, ok := c.Get("user_info")
	if !ok {
		return nil
	}

	res, ok := get.(*dao.Session)
	if !ok {
		return nil
	}
	return res
}

func (a *Auth) SetNewCookie(id int, c *gin.Context) error {
	session, err := a.generateSession(id, c.ClientIP(), c.Request.UserAgent())
	if err != nil {
		return err
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(a.cfg.Session.CookieName, session, int(a.cfg.Session.Duration.Seconds()),
		a.cfg.Session.CookiePath, a.cfg.Session.Domain, true, true)
	return nil
}

// GenerateSecretCode for email auth
func (a *Auth) GenerateSecretCode() string {
	b := make([]rune, 5)
	for i := range b {
		b[i] = rand.Int31n(26) + 65
	}
	return string(b)
}

// GenerateFileName for file creation
func (a *Auth) GenerateFileName(c *gin.Context, file *multipart.FileHeader) (string, error) {
	b := make([]rune, rand.Int31n(30)+20)
	for i := range b {
		b[i] = rand.Int31n(26) + 97
	}
	generated := string(b) + filepath.Ext(file.Filename)
	if err := c.SaveUploadedFile(file, a.cfg.Files.Path+generated); err != nil {
		if os.IsExist(err) {
			return a.GenerateFileName(c, file)
		}
		return "", err
	}
	return generated, nil
}
