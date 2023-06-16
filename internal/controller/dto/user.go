package dto

type UpdateUser struct {
	Biography *string `json:"biography,omitempty" sql:"biography" validate:"omitempty,lte=512"`
	Language  *string `json:"language,omitempty" sql:"language" validate:"omitempty,enum=EN*RU"`
	Theme     *string `json:"theme,omitempty" sql:"theme" validate:"omitempty,enum=SYSTEM*WHITE*DARK"`
	FirstName *string `json:"firstName,omitempty" sql:"first_name" validate:"omitempty,gte=3,lte=32"`
	LastName  *string `json:"lastName,omitempty" sql:"last_name" validate:"omitempty,gte=3,lte=32"`
}

type UpdateEmail struct {
	Password string `json:"password,omitempty" validate:"required,gte=4,lte=20,password"`
	NewEmail string `json:"newEmail,omitempty" validate:"required,email"`
}

type UpdatePassword struct {
	NewPassword string `json:"newPassword,omitempty" validate:"required,gte=4,lte=20,password"`
	OldPassword string `json:"oldPassword,omitempty" validate:"required,gte=4,lte=20,password"`
}

type UpdateName struct {
	NewName string `json:"newName,omitempty" validate:"required,gte=5,lte=20,name"`
}
