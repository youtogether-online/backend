package dto

import (
	"github.com/wtkeqrf0/you_together/ent/user"
	"strings"
)

// MyUserDTO with detail info
type MyUserDTO struct {
	Username        string        `json:"username,omitempty" sql:"username"`
	Email           string        `json:"email,omitempty" sql:"email"`
	IsEmailVerified bool          `json:"isEmailVerified" sql:"is_email_verified"`
	Biography       string        `json:"biography,omitempty" sql:"biography"`
	Role            user.Role     `json:"role,omitempty" sql:"role"`
	Avatar          string        `json:"avatar,omitempty" sql:"avatar"`
	FriendsIds      []int         `json:"friendsIds,omitempty" sql:"friends_ids"`
	Language        user.Language `json:"language,omitempty" sql:"language"`
	Theme           user.Theme    `json:"theme,omitempty" sql:"theme"`
	FirstName       string        `json:"firstName,omitempty" sql:"first_name"`
	LastName        string        `json:"lastName,omitempty" sql:"last_name"`
}

func (m *MyUserDTO) CutEmail() {
	m.Email = (m.Email)[:1] + "**" + (m.Email)[strings.Index(m.Email, "@")-1:]
}

// UserDTO with main info
type UserDTO struct {
	Username   string    `json:"username,omitempty" sql:"username"`
	Biography  string    `json:"biography,omitempty" sql:"biography"`
	Role       user.Role `json:"role,omitempty" sql:"role"`
	Avatar     string    `json:"avatar,omitempty" sql:"avatar"`
	FriendsIds []int     `json:"friendsIds,omitempty" sql:"friends_ids"`
	FirstName  string    `json:"firstName,omitempty" sql:"first_name"`
	LastName   string    `json:"lastName,omitempty" sql:"last_name"`
}

type SignInDTO struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,gte=4,lte=20"`
	Device   string `json:"device,omitempty" validate:"required,lt=50"`
}

type EmailDTO struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}

type EmailWithCodeDTO struct {
	Email  string `json:"email,omitempty" validate:"required,email"`
	Code   string `json:"code,omitempty" validate:"required,len=5"`
	Device string `json:"device,omitempty" validate:"required,lt=50"`
}
