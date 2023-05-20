package dto

type UpdateEmail struct {
	Password string `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20"`
	NewEmail string `json:"newEmail,omitempty" validate:"required,email"`
}

type UpdatePassword struct {
	NewPassword string `json:"newPassword,omitempty" validate:"required,printascii,gte=4,lte=20"`
	Email       string `json:"email,omitempty" validate:"required,email"`
	Code        string `json:"code,omitempty" validate:"required,len=5"`
}

type UpdateName struct {
	NewName string `json:"newName,omitempty" validate:"required,gte=5,lte=20,name"`
}

type Name struct {
	Name string `json:"name,omitempty" validate:"required,gte=5,lte=20,name"`
}

type EmailWithPassword struct {
	Email     string  `json:"email,omitempty" validate:"required,email"`
	Password  string  `json:"password,omitempty" validate:"required,printascii,gte=4,lte=20"`
	UserAgent string  `json:"-" header:"User-Agent"`
	Language  *string `json:"language,omitempty" header:"Accept-Language" enum:"EN,RU" default:"EN"`
	Theme     *string `json:"theme,omitempty" enum:"SYSTEM,DARK,WHITE" default:"SYSTEM"`
}

type Email struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}

type EmailWithCode struct {
	Email     string  `json:"email,omitempty" validate:"required,email"`
	Code      string  `json:"code,omitempty" validate:"required,len=5"`
	UserAgent string  `json:"-" header:"User-Agent"`
	Language  *string `json:"language,omitempty" header:"Accept-Language" enum:"EN,RU" default:"EN"`
	Theme     *string `json:"theme,omitempty" enum:"SYSTEM,DARK,WHITE" default:"SYSTEM"`
}

type UpdateUser struct {
	Biography *string `json:"biography,omitempty" sql:"biography"`
	Language  *string `json:"language,omitempty" enum:"EN,RU" default:"EN" sql:"language"`
	Theme     *string `json:"theme,omitempty" enum:"SYSTEM,DARK,WHITE" default:"SYSTEM" sql:"theme"`
	FirstName *string `json:"firstName,omitempty" sql:"first_name"`
	LastName  *string `json:"lastName,omitempty" sql:"last_name"`
}
