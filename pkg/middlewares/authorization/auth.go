package authorization

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type AuthService interface {
	SetSession(sessionId string, info map[string]string) error
	GetSession(sessionId string) (map[string]string, error)
	UserExists(username string) bool
	GetTTL(sessionId string) time.Duration
}

type Auth struct {
	auth AuthService
}

func NewAuth(auth AuthService) *Auth {
	return &Auth{auth: auth}
}

// ValidateSession validates the session and identifies the user in DB. Returns an error in case of unsuccessful validation
func (a Auth) ValidateSession(sessionId string) (map[string]string, error) {
	if sessionId == "" {
		return make(map[string]string), fmt.Errorf("session id is not found")
	}

	info, err := a.auth.GetSession(sessionId)
	if err != nil {
		return make(map[string]string), err
	}

	if !a.auth.UserExists(info["username"]) {
		return make(map[string]string), fmt.Errorf("user does not exists")
	}
	return info, nil
}

// GenerateSession generates a new session, if not exists
func (a Auth) GenerateSession(username, ip, device string) (string, map[string]string, error) {
	// TODO session updating while user sign-in again (not creating new session)
	sessionId := uuid.New().String()
	info := map[string]string{
		"username": username,
		"ip":       ip,
		"device":   device,
	}

	return sessionId, info, a.auth.SetSession(sessionId, info)
}
