package dto

import "mime/multipart"

type UpdateUser struct {
	Biography *string `json:"biography,omitempty" validate:"omitempty,lte=512"`
	Language  *string `json:"language,omitempty" validate:"omitempty,enum=EN*RU"`
	Theme     *string `json:"theme,omitempty" validate:"omitempty,enum=system*light*dark"`
	FirstName *string `json:"firstName,omitempty" validate:"omitempty,gte=3,lte=32"`
	LastName  *string `json:"lastName,omitempty" validate:"omitempty,gte=3,lte=32"`
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

type UpdateImage struct {
	Image *multipart.FileHeader `form:"file,omitempty"`
}

type NameParam struct {
	Name string `uri:"name" validate:"required,gte=5,lte=20,name"`
}
