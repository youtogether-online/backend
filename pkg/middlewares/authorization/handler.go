package authorization

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wtkeqrf0/you_together/internal/controller/dto"
	"github.com/wtkeqrf0/you_together/pkg/conf"
	"github.com/wtkeqrf0/you_together/pkg/middlewares/exceptions"
	"time"
)

var cfg = conf.GetConfig()

// RequireAuth authorizes the user
func (a Auth) RequireAuth(c *gin.Context) {
	at, _ := c.Cookie(cfg.Token.AccessName)
	claims, err := a.ValidateJWT(at)
	if err != nil {
		rt, err := c.Cookie(cfg.Token.RefreshName)
		if err != nil {
			c.Error(exceptions.UnAuthorized.AddErr(err))
			return
		}

		claims, err = a.ValidateJWT(rt)
		if err != nil {
			c.Error(exceptions.UnAuthorized.AddErr(err))
			return
		}

		t, err := a.GenerateJWT(claims["sub"].(float64))
		if err != nil {
			c.Error(exceptions.ServerError.AddErr(err))
			return
		}

		c.SetCookie(cfg.Token.AccessName, t.AT, cfg.Token.AccessDurationInSeconds,
			"/api", cfg.Listen.Host, false, true)

		c.SetCookie(cfg.Token.RefreshName, t.RT, cfg.Token.RefreshDurationInSeconds,
			"/api", cfg.Listen.Host, false, true)

		c.Set("claims", t.AClaims)
	} else {
		c.Set("claims", claims)
	}
	c.Next()
}

// MaybeAuth authorizes the user, if the token exist. User can not be authorized
func (a Auth) MaybeAuth(c *gin.Context) {
	at, _ := c.Cookie(cfg.Token.AccessName)
	claims, err := a.ValidateJWT(at)
	if err != nil {
		rt, err := c.Cookie(cfg.Token.RefreshName)
		if err != nil {
			c.Next()
			return
		}

		claims, err = a.ValidateJWT(rt)
		if err != nil {
			c.Next()
			return
		}

		t, err := a.GenerateJWT(claims["sub"].(float64))
		if err != nil {
			c.Next()
			return
		}

		c.SetCookie(cfg.Token.AccessName, t.AT, cfg.Token.AccessDurationInSeconds,
			"/api", cfg.Listen.Host, false, true)

		c.SetCookie(cfg.Token.RefreshName, t.RT, cfg.Token.RefreshDurationInSeconds,
			"/api", cfg.Listen.Host, false, true)

		c.Set("claims", t.AClaims)
	} else {
		c.Set("claims", claims)
	}
	c.Next()
}

// ValidateJWT validates the token and identifies the user with isExist function. Returns an error in case of unsuccessful validation. Otherwise, returns the user ID.
func (a Auth) ValidateJWT(token string) (jwt.MapClaims, error) {
	if token == "" {
		return jwt.MapClaims{}, fmt.Errorf("token not found")
	}

	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to decrypt token")
		}

		return []byte(cfg.Token.Secret), nil
	})

	if claims, ok := at.Claims.(jwt.MapClaims); ok && at.Valid && err == nil {
		if userId, oki := claims["sub"].(float64); oki && float64(time.Now().Unix()) < claims["exp"].(float64) {
			if a.UserExists(int(userId)) {
				return claims, nil
			}
			return jwt.MapClaims{}, fmt.Errorf("user not found by token")
		}
		return jwt.MapClaims{}, fmt.Errorf("token fields not found: %v", err)
	}
	return jwt.MapClaims{}, fmt.Errorf("token is not valid: %v", err)
}

// GenerateJWT generates a new pair of JWT and inserts it into Cookie
func (a Auth) GenerateJWT(id float64) (dto.TokensDTO, error) {
	t := dto.TokensDTO{AClaims: jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	}, RClaims: jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}}

	at, err := jwt.NewWithClaims(jwt.SigningMethodHS256, t.AClaims).SignedString([]byte(cfg.Token.Secret))
	if err != nil {
		return dto.TokensDTO{}, err
	}
	t.AT = at

	rt, err := jwt.NewWithClaims(jwt.SigningMethodHS256, t.RClaims).SignedString([]byte(cfg.Token.Secret))
	if err != nil {
		return dto.TokensDTO{}, err
	}
	t.RT = rt

	return t, nil
}
