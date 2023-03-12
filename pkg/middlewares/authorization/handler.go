package authorization

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wtkeqrf0/you_together/pkg/conf"
	"github.com/wtkeqrf0/you_together/pkg/middlewares/exceptions"
	"time"
)

var cfg = conf.GetConfig()

// RequireAuth authorizes the user
func (a Auth) RequireAuth(c *gin.Context) {
	at, _ := c.Cookie(cfg.Token.AccessName)
	userId, err := a.validateJWT(at)
	if err != nil {
		rt, err := c.Cookie(cfg.Token.RefreshName)
		if err != nil {
			c.Error(exceptions.UnAuthorized.AddErr(err.Error()))
			return
		}

		userId, err = a.validateJWT(rt)
		if err != nil {
			c.Error(exceptions.UnAuthorized.AddErr(err.Error()))
			return
		}

		err = a.GenerateJWT(userId, c)
		if err != nil {
			c.Error(exceptions.ServerError.AddErr(err.Error()))
			return
		}
		c.Set("ID", userId)
	} else {
		c.Set("ID", userId)
	}
	c.Next()
}

// validateJWT validates the token and identifies the user with isExist function. Returns an error in case of unsuccessful validation. Otherwise, returns the user ID.
func (a Auth) validateJWT(token string) (int, error) {
	if token == "" {
		return 0, fmt.Errorf("token not found")
	}

	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to decrypt token")
		}

		return []byte(cfg.Token.Secret), nil
	})

	if claims, ok := at.Claims.(jwt.MapClaims); ok && at.Valid && err == nil {
		if userId, oki := claims["user_id"].(float64); oki && float64(time.Now().Unix()) < claims["exp"].(float64) {
			if a.UserExists(int(userId)) {
				return int(userId), nil
			}
			return 0, fmt.Errorf("user not found by token")
		}
		return 0, fmt.Errorf("token fields not found: %v", err)
	}
	return 0, fmt.Errorf("token is not valid: %v", err)
}

// GenerateJWT generates a new pair of JWT and inserts it into Cookie
func (a Auth) GenerateJWT(id int, c *gin.Context) error {

	at, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	}).SignedString([]byte(cfg.Token.Secret))
	if err != nil {
		return err
	}

	rt, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}).SignedString([]byte(cfg.Token.Secret))
	if err != nil {
		return err
	}

	c.SetCookie(cfg.Token.AccessName, at, cfg.Token.AccessDurationInSeconds,
		"/api", cfg.Listen.Host, false, true)

	c.SetCookie(cfg.Token.RefreshName, rt, cfg.Token.RefreshDurationInSeconds,
		"/api", cfg.Listen.Host, false, true)

	return nil
}
