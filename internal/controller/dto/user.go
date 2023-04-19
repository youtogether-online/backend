package dto

type DTO interface {
	MyUserDTO |
		UserDTO |
		UpdateUserDTO |
		UpdateEmailDTO |
		UpdatePasswordDTO |
		UpdateNameDTO |
		EmailWithPasswordDTO |
		EmailDTO |
		EmailWithCodeDTO |
		NameDTO
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

type UpdateNameDTO struct {
	NewUsername string `json:"newName,omitempty" validate:"required,gte=5,lte=20"`
}

type EmailWithPasswordDTO struct {
	Email    string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Password string `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" example:"onkr3451"`
	Language string `json:"-" enum:"EN,RU" default:"EN"`
	Theme    string `json:"theme,omitempty" example:"DARK" enum:"SYSTEM,DARK,WHITE" default:"SYSTEM"`
}

type EmailDTO struct {
	Email string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
}

type EmailWithCodeDTO struct {
	Email    string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Code     string `json:"code,omitempty" validate:"required,len=5" length:"5" example:"I1ELB"`
	Language string `json:"-" enum:"EN,RU" default:"EN"`
	Theme    string `json:"theme,omitempty" example:"DARK" enum:"SYSTEM,DARK,WHITE" default:"SYSTEM"`
}

type NameDTO struct {
	Name string `json:"name,omitempty" validate:"required,gte=5,lte=20"`
}

type TypeDTO struct {
	Type   string `json:"type,omitempty" example:"user"`
	Object any    `json:"object,omitempty" example:"dto.UserDTO"`
}
