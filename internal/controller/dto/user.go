package dto

type UpdateEmail struct {
	Password string `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" example:"onkr3451"`
	NewEmail string `json:"newEmail,omitempty" validate:"required,email" example:"myemail@gmail.com"`
}

type UpdatePassword struct {
	NewPassword string `json:"newPassword,omitempty" validate:"required,printascii,gte=4,lte=20" example:"onkr3451"`
	Email       string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Code        string `json:"code,omitempty" validate:"required,len=5" length:"5" example:"I1ELB"`
}

type UpdateName struct {
	NewName string `json:"newName,omitempty" validate:"required,gte=5,lte=20,name"`
}

type Name struct {
	Name string `json:"name,omitempty" validate:"required,gte=5,lte=20,name"`
}

type EmailWithPassword struct {
	Email    string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Password string `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20" example:"onkr3451"`
	Language string `json:"-" enum:"EN,RU" default:"EN"`
	Theme    string `json:"theme,omitempty" example:"DARK" enum:"SYSTEM,DARK,WHITE" default:"SYSTEM"`
}

type Email struct {
	Email string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
}

type EmailWithCode struct {
	Email    string `json:"email,omitempty" validate:"required,email" example:"myemail@gmail.com"`
	Code     string `json:"code,omitempty" validate:"required,len=5" length:"5" example:"I1ELB"`
	Language string `json:"-" enum:"EN,RU" default:"EN"`
	Theme    string `json:"theme,omitempty" example:"DARK" enum:"SYSTEM,DARK,WHITE" default:"SYSTEM"`
}

type UpdateUser struct {
	Biography *string `json:"biography,omitempty" validate:"required_without_all,gt=140" sql:"biography" example:"I'd like to relax"`
	Language  *string `json:"language,omitempty" validate:"required_without_all" sql:"language"`
	Theme     *string `json:"theme,omitempty" validate:"required_without_all" sql:"theme"`
	FirstName *string `json:"firstName,omitempty" validate:"required_without_all,lte=30" sql:"first_name" example:"Tele"`
	LastName  *string `json:"lastName,omitempty" validate:"required_without_all,lte=30" sql:"last_name" example:"phone"`
}
