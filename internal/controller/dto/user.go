package dto

import (
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
	Email    string `json:"Email,omitempty" binding:"required"`
	Password string `json:"Password,omitempty" binding:"required"`
}

type EmailDTO struct {
	Email string `json:"Email,omitempty" binding:"required"`
}

type EmailWithCodeDTO struct {
	Email string `json:"Email,omitempty" binding:"required"`
	Code  string `json:"Code,omitempty" binding:"required"`
}

func CutEmail(email *string) {
	*email = (*email)[:1] + "**" + (*email)[strings.Index(*email, "@")-1:]
}
