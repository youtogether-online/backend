package dto

type UpdateUser struct {
	Biography *string `json:"biography,omitempty" sql:"biography" validate:"omitempty,lte=512"`
	Language  *string `json:"language,omitempty" sql:"language" validate:"omitempty,enum=EN*RU"`
	Theme     *string `json:"theme,omitempty" sql:"theme" validate:"omitempty,enum=SYSTEM*WHITE*DARK"`
	FirstName *string `json:"firstName,omitempty" sql:"first_name" validate:"omitempty,gte=3,lte=32"`
	LastName  *string `json:"lastName,omitempty" sql:"last_name" validate:"omitempty,gte=3,lte=32"`
}

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

type Email struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}

type EmailWithPassword struct {
	Email    string  `json:"email,omitempty" header:"-" validate:"required,email"`
	Password string  `json:"password,omitempty" header:"-" validate:"required,printascii,gte=4,lte=20"`
	Language *string `json:"-" header:"Accept-Language" validate:"omitempty,enum=EN*RU"`
}

type EmailWithCode struct {
	Email    string  `json:"email,omitempty" validate:"required,email"`
	Code     string  `json:"code,omitempty" validate:"required,len=5"`
	Language *string `json:"-" header:"Accept-Language" validate:"omitempty,enum=EN*RU"`
}
