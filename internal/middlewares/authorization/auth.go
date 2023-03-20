package authorization

import (
	"fmt"
	"github.com/google/uuid"
)

type AuthService interface {
	SetSession(sessionId string, info map[string]string) error
	GetSession(sessionId string) (map[string]string, error)
	ExpandExpireSession(sessionId string) error
	UserExists(username string) bool
	FindSessionsByUsername(userName string) []map[string]string
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
		return nil, fmt.Errorf("session id is not found")
	}

	info, err := a.auth.GetSession(sessionId)
	if err != nil {
		return nil, err
	}

	if !a.auth.UserExists(info["username"]) {
		return nil, fmt.Errorf("user does not exists")
	}

	if err = a.auth.ExpandExpireSession(sessionId); err != nil {
		return nil, fmt.Errorf("session does not exist: %v", err)
	}

	return info, nil
}

// GenerateSession generates a new session, if not exists
func (a Auth) GenerateSession(username, ip, device string) (string, map[string]string, error) {
	sessionId := uuid.New().String()
	info := map[string]string{
		"username": username,
		"ip":       ip,
		"device":   device,
	}

	return sessionId, info, a.auth.SetSession(sessionId, info)
}
