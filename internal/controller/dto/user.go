package dto

import (
	"github.com/wtkeqrf0/you_together/ent/user"
	"strings"
)

// @Description User detail information
type MyUserDTO struct {
	Username        string        `json:"username,omitempty" sql:"username" example:"bobbas"`
	Email           string        `json:"email,omitempty" sql:"email" example:"myemail@gmail.com"`
	IsEmailVerified bool          `json:"isEmailVerified" sql:"is_email_verified"`
	Biography       string        `json:"biography,omitempty" sql:"biography" example:"I'd like to relax"`
	Role            user.Role     `json:"role,omitempty" sql:"role"`
	FriendsIds      []string      `json:"friendsIds,omitempty" sql:"friends_ids" example:"tldtb,kigfv"`
	Language        user.Language `json:"language,omitempty" sql:"language"`
	Theme           user.Theme    `json:"theme,omitempty" sql:"theme"`
	FirstName       string        `json:"firstName,omitempty" sql:"first_name" example:"Tele"`
	LastName        string        `json:"lastName,omitempty" sql:"last_name" example:"phone"`
}

func (m *MyUserDTO) CutEmail() {
	m.Email = (m.Email)[:1] + "**" + (m.Email)[strings.Index(m.Email, "@")-1:]
}

// @Description User main information
type UserDTO struct {
	Username   string    `json:"username,omitempty" sql:"username" example:"bobbas"`
	Biography  string    `json:"biography,omitempty" sql:"biography" example:"I'd like to relax"`
	Role       user.Role `json:"role,omitempty" sql:"role"`
	FriendsIds []string  `json:"friendsIds,omitempty" sql:"friends_ids" example:"tldtb,kigfv"`
	FirstName  string    `json:"firstName,omitempty" sql:"first_name" example:"Tele"`
	LastName   string    `json:"lastName,omitempty" sql:"last_name" example:"phone"`
}

type SignInDTO struct {
	Email    string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Password string `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" example:"onkr3451"`
	Device   string `json:"device,omitempty" validate:"required,lt=50" example:"macOS 10.15.7 Chrome 111.0.0"`
}

type EmailDTO struct {
	Email string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
}

type EmailWithCodeDTO struct {
	Email  string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Code   string `json:"code,omitempty" validate:"required,len=5" length:"5" example:"I1ELB"`
	Device string `json:"device,omitempty" validate:"required,lt=50" example:"macOS 10.15.7 Chrome 111.0.0"`
}
