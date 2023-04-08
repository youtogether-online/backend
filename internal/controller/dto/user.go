package dto

import (
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/ent/user"
)

type DTO interface {
	MyUserDTO |
		UserDTO |
		UpdateUserDTO |
		UpdateEmailDTO |
		UpdatePasswordDTO |
		UpdateUsernameDTO |
		EmailWithPasswordDTO |
		EmailDTO |
		EmailWithCodeDTO |
		UsernameDTO
}

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

// @Description User's main information
type UserDTO struct {
	Biography  string    `json:"biography,omitempty" sql:"biography" example:"I'd like to relax"`
	Role       user.Role `json:"role,omitempty" sql:"role"`
	FriendsIds []string  `json:"friendsIds,omitempty" sql:"friends_ids" example:"tldtb,kigfv"`
	FirstName  string    `json:"firstName,omitempty" sql:"first_name" example:"Tele"`
	LastName   string    `json:"lastName,omitempty" sql:"last_name" example:"phone"`
}

type UpdateUserDTO struct {
	Biography string        `json:"biography,omitempty" validate:"required_without_all,gt=140" sql:"biography" example:"I'd like to relax"`
	Language  user.Language `json:"language,omitempty" validate:"required_without_all" sql:"language"`
	Theme     user.Theme    `json:"theme,omitempty" validate:"required_without_all" sql:"theme"`
	FirstName string        `json:"firstName,omitempty" validate:"required_without_all,lte=30" sql:"first_name" example:"Tele"`
	LastName  string        `json:"lastName,omitempty" validate:"required_without_all,lte=30" sql:"last_name" example:"phone"`
}

type UpdateEmailDTO struct {
	Password string `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" example:"onkr3451"`
	NewEmail string `json:"newEmail,omitempty" validate:"required,email" example:"myemail@gmail.com"`
}

type UpdatePasswordDTO struct {
	NewPassword string `json:"newPassword,omitempty" validate:"required,printascii,gte=4,lte=20" example:"onkr3451"`
	Email       string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Code        string `json:"code,omitempty" validate:"required,len=5" length:"5" example:"I1ELB"`
}

type UpdateUsernameDTO struct {
	NewUsername string `json:"newUsername,omitempty" validate:"required,gte=5,lte=20"`
}

type EmailWithPasswordDTO struct {
	Email    string        `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Password string        `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" example:"onkr3451"`
	Language user.Language `json:"-"`
	Theme    user.Theme    `json:"theme,omitempty" example:"DARK"`
}

type EmailDTO struct {
	Email string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
}

type EmailWithCodeDTO struct {
	Email    string        `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Code     string        `json:"code,omitempty" validate:"required,len=5" length:"5" example:"I1ELB"`
	Language user.Language `json:"-"`
	Theme    user.Theme    `json:"theme,omitempty" example:"DARK"`
}

type UsernameDTO struct {
	Username string `json:"username,omitempty" validate:"required,gte=5,lte=20"`
}

func Convert(user *ent.User) MyUserDTO {
	return MyUserDTO{
		Username:        user.Name,
		Email:           user.Email,
		IsEmailVerified: user.IsEmailVerified,
		Biography:       user.Biography,
		Role:            user.Role,
		FriendsIds:      user.FriendsIds,
		Language:        user.Language,
		Theme:           user.Theme,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
	}
}
