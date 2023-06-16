package dto

type Room struct {
	Privacy     *string `json:"privacy,omitempty" validate:"omitempty,enum=PUBLIC*PRIVATE*FRIENDS"`
	Password    *string `json:"password,omitempty" validate:"omitempty,gte=4,lte=20,password"`
	CustomName  *string `json:"custom_name,omitempty" validate:"omitempty,gte=3,lte=32"`
	Description *string `json:"description,omitempty" validate:"omitempty,lte=140"`
	HasChat     *bool   `json:"has_chat,omitempty"`
}
