package dto

type UpdatePasswordAuth struct {
	NewPassword string `json:"newPassword,omitempty" validate:"required,gte=4,lte=20,password"`
	Email       string `json:"email,omitempty" validate:"required,email"`
	Code        string `json:"code,omitempty" validate:"required,len=5"`
}

type Email struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}

type EmailWithPassword struct {
	Email    string  `json:"email,omitempty" header:"-" validate:"required,email"`
	Password string  `json:"password,omitempty" header:"-" validate:"required,gte=4,lte=20,password"`
	Language *string `json:"-" header:"Accept-Language" validate:"omitempty,enum=EN*RU"`
}

type EmailWithCode struct {
	Email    string  `json:"email,omitempty" validate:"required,email"`
	Code     string  `json:"code,omitempty" validate:"required,len=5"`
	Language *string `json:"-" header:"Accept-Language" validate:"omitempty,enum=EN*RU"`
}
