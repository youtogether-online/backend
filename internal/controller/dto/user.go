package dto

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/wtkeqrf0/you_together/ent/user"
	"strings"
)

// MyUserDTO with detail info
type MyUserDTO struct {
	UserName   string        `json:"UserName,omitempty" sql:"user_name"`
	Email      string        `json:"Email,omitempty" sql:"email"`
	Biography  string        `json:"Biography,omitempty" sql:"biography"`
	Role       user.Role     `json:"Role,omitempty" sql:"role"`
	Avatar     string        `json:"Avatar,omitempty" sql:"avatar"`
	FriendsIds []int         `json:"FriendsIds,omitempty" sql:"friends_ids"`
	Language   user.Language `json:"Language,omitempty" sql:"language"`
	Theme      user.Theme    `json:"Theme,omitempty" sql:"theme"`
	Name       string        `json:"Name,omitempty" sql:"name"`
}

func (m *MyUserDTO) CutEmail() {
	m.Email = (m.Email)[:1] + "**" + (m.Email)[strings.Index(m.Email, "@")-1:]
}

// UserDTO with main info
type UserDTO struct {
	UserName   string    `json:"UserName,omitempty" sql:"user_name"`
	Biography  string    `json:"Biography,omitempty" sql:"biography"`
	Role       user.Role `json:"Role,omitempty" sql:"role"`
	Avatar     string    `json:"Avatar,omitempty" sql:"avatar"`
	FriendsIds []int     `json:"FriendsIds,omitempty" sql:"friends_ids"`
	Name       string    `json:"Name,omitempty" sql:"name"`
}

type SignInDTO struct {
	Email    string `json:"Email,omitempty" validate:"required,email"`
	Password string `json:"Password,omitempty" validate:"required,gte=6,lte=20"`
}

type EmailDTO struct {
	Email string `json:"Email,omitempty" validate:"required,email"`
}

type EmailWithCodeDTO struct {
	Email string `json:"Email,omitempty" validate:"required,email"`
	Code  string `json:"Code,omitempty" validate:"required,len=5"`
}

type TokensDTO struct {
	AT      string
	RT      string
	AClaims jwt.MapClaims
	RClaims jwt.MapClaims
}
