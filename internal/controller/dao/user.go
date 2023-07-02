package dao

import (
	"github.com/wtkeqrf0/you-together/ent"
	"strings"
	"time"
)

type Me struct {
	Name            string    `json:"name,omitempty"`
	Email           string    `json:"email,omitempty"`
	IsEmailVerified bool      `json:"isEmailVerified"`
	Biography       *string   `json:"biography,omitempty"`
	Role            string    `json:"role,omitempty"`
	FriendsIds      []string  `json:"friendsIds,omitempty"`
	Language        string    `json:"language,omitempty"`
	Theme           string    `json:"theme,omitempty"`
	FirstName       *string   `json:"firstName,omitempty"`
	LastName        *string   `json:"lastName,omitempty"`
	Image           string    `json:"image,omitempty"`
	CreateTime      time.Time `json:"createTime,omitempty"`
}

func TransformToMe(user *ent.User) *Me {
	return &Me{
		Email:           user.Email[:1] + "**" + user.Email[strings.Index(user.Email, "@")-1:],
		IsEmailVerified: user.IsEmailVerified,
		Language:        user.Language,
		Theme:           user.Theme,
		Name:            user.Name,
		Biography:       user.Biography,
		Role:            user.Role,
		FriendsIds:      user.FriendsIds,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Image:           user.Image,
		CreateTime:      user.CreateTime,
	}
}

type User struct {
	Name       string    `json:"name,omitempty"`
	Biography  *string   `json:"biography,omitempty"`
	Role       string    `json:"role,omitempty"`
	FriendsIds []string  `json:"friendsIds,omitempty"`
	FirstName  *string   `json:"firstName,omitempty"`
	LastName   *string   `json:"lastName,omitempty"`
	Image      string    `json:"image,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
}

func TransformToUser(user *ent.User) *User {
	return &User{
		Name:       user.Name,
		Biography:  user.Biography,
		Role:       user.Role,
		FriendsIds: user.FriendsIds,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Image:      user.Image,
		CreateTime: user.CreateTime,
	}
}
