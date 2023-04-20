package dto

// @Description User's session
type Session struct {
	ID      int    `json:"id" redis:"ID"`
	IP      string `json:"ip" redis:"IP"`
	Device  string `json:"device" redis:"Device"`
	Browser string `json:"browser" redis:"Browser"`
	Updated int64  `json:"updated" redis:"Updated"`
}

// @Description User detail information
type MyUserDTO struct {
	Name            string   `json:"name,omitempty" sql:"name" example:"bobbas"`
	Email           string   `json:"email,omitempty" sql:"email" example:"myemail@gmail.com"`
	IsEmailVerified bool     `json:"isEmailVerified" sql:"is_email_verified"`
	Biography       string   `json:"biography,omitempty" sql:"biography" example:"I'd like to relax"`
	Role            string   `json:"role,omitempty" sql:"role"`
	FriendsIds      []string `json:"friendsIds,omitempty" sql:"friends_ids" example:"tldtb,kigfv"`
	Language        string   `json:"language,omitempty" sql:"language"`
	Theme           string   `json:"theme,omitempty" sql:"theme"`
	FirstName       string   `json:"firstName,omitempty" sql:"first_name" example:"Tele"`
	LastName        string   `json:"lastName,omitempty" sql:"last_name" example:"phone"`
}

// @Description User's main information
type UserDTO struct {
	Name       string   `json:"name,omitempty" sql:"name" example:"bobbas"`
	Biography  string   `json:"biography,omitempty" sql:"biography" example:"I'd like to relax"`
	Role       string   `json:"role,omitempty" sql:"role"`
	FriendsIds []string `json:"friendsIds,omitempty" sql:"friends_ids" example:"tldtb,kigfv"`
	FirstName  string   `json:"firstName,omitempty" sql:"first_name" example:"Tele"`
	LastName   string   `json:"lastName,omitempty" sql:"last_name" example:"phone"`
}

type UpdateUserDTO struct {
	Biography *string `json:"biography,omitempty" validate:"required_without_all,gt=140" sql:"biography" example:"I'd like to relax"`
	Language  *string `json:"language,omitempty" validate:"required_without_all" sql:"language"`
	Theme     *string `json:"theme,omitempty" validate:"required_without_all" sql:"theme"`
	FirstName *string `json:"firstName,omitempty" validate:"required_without_all,lte=30" sql:"first_name" example:"Tele"`
	LastName  *string `json:"lastName,omitempty" validate:"required_without_all,lte=30" sql:"last_name" example:"phone"`
}
