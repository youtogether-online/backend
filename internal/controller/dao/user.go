package dao

type Session struct {
	ID      int    `json:"id" redis:"ID"`
	IP      string `json:"ip" redis:"IP"`
	Device  string `json:"device" redis:"Device"`
	Browser string `json:"browser" redis:"Browser"`
	Updated int64  `json:"updated" redis:"Updated"`
}

type Me struct {
	Name            string   `json:"name,omitempty" sql:"name"`
	Email           string   `json:"email,omitempty" sql:"email"`
	IsEmailVerified bool     `json:"isEmailVerified,omitempty" sql:"is_email_verified"`
	Biography       string   `json:"biography,omitempty" sql:"biography"`
	Role            string   `json:"role,omitempty" sql:"role"`
	FriendsIds      []string `json:"friendsIds,omitempty" sql:"friends_ids"`
	Language        string   `json:"language,omitempty" sql:"language"`
	Theme           string   `json:"theme,omitempty" sql:"theme"`
	FirstName       string   `json:"firstName,omitempty" sql:"first_name"`
	LastName        string   `json:"lastName,omitempty" sql:"last_name"`
}

type User struct {
	Name       string   `json:"name,omitempty" sql:"name"`
	Biography  string   `json:"biography,omitempty" sql:"biography"`
	Role       string   `json:"role,omitempty" sql:"role"`
	FriendsIds []string `json:"friendsIds,omitempty" sql:"friends_ids"`
	FirstName  string   `json:"firstName,omitempty" sql:"first_name"`
	LastName   string   `json:"lastName,omitempty" sql:"last_name"`
}
