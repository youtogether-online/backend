package dao

// @Description User's session information
type Session struct {
	ID      int    `json:"id" redis:"ID"`
	IP      string `json:"ip" redis:"IP"`
	Device  string `json:"device" redis:"Device"`
	Browser string `json:"browser" redis:"Browser"`
	Updated int64  `json:"updated" redis:"Updated"`
}

// @Description Detail information about the user
type Me struct {
	Name            string   `json:"name,omitempty" sql:"name" example:"bobbas"`
	Email           string   `json:"email,omitempty" sql:"email" example:"myemail@gmail.com"`
	IsEmailVerified bool     `json:"isEmailVerified,omitempty" sql:"is_email_verified" example:"true"`
	Biography       string   `json:"biography,omitempty" sql:"biography" example:"I'd like to relax"`
	Role            string   `json:"role,omitempty" sql:"role" example:"USER"`
	FriendsIds      []string `json:"friendsIds,omitempty" sql:"friends_ids" example:"bobba, imaxied"`
	Language        string   `json:"language,omitempty" sql:"language" example:"RU"`
	Theme           string   `json:"theme,omitempty" sql:"theme" example:"DARK"`
	FirstName       string   `json:"firstName,omitempty" sql:"first_name" example:"Tele"`
	LastName        string   `json:"lastName,omitempty" sql:"last_name" example:"phone"`
}

// @Description Main information about the user
type User struct {
	Name       string   `json:"name,omitempty" sql:"name" example:"bobbas"`
	Biography  string   `json:"biography,omitempty" sql:"biography" example:"I'd like to relax"`
	Role       string   `json:"role,omitempty" sql:"role" example:"USER"`
	FriendsIds []string `json:"friendsIds,omitempty" sql:"friends_ids" example:"tldtb,kigfv"`
	FirstName  string   `json:"firstName,omitempty" sql:"first_name" example:"Tele"`
	LastName   string   `json:"lastName,omitempty" sql:"last_name" example:"phone"`
}
