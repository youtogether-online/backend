package dto

type Room struct {
	Privacy     *string `json:"privacy,omitempty" validate:"omitempty,enum=public*private*friends"`
	Password    *string `json:"password,omitempty" validate:"omitempty,gte=4,lte=20,password"`
	HasChat     bool    `json:"hasChat,omitempty"`
	Title       string  `json:"title,omitempty" validate:"omitempty,gte=3,lte=32"`
	Description *string `json:"description,omitempty" validate:"omitempty,lte=140"`
}

type RoomId struct {
	Id int `uri:"id" validate:"gte=0"`
}
