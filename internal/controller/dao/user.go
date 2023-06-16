package dao

import (
	"github.com/wtkeqrf0/you-together/ent"
	"strings"
	"time"
)

type Me struct {
	Name            string    `json:"name,omitempty" sql:"name"`
	Email           string    `json:"email,omitempty" sql:"email"`
	IsEmailVerified bool      `json:"isEmailVerified,omitempty" sql:"is_email_verified"`
	Biography       *string   `json:"biography,omitempty" sql:"biography"`
	Role            string    `json:"role,omitempty" sql:"role"`
	FriendsIds      []string  `json:"friendsIds,omitempty" sql:"friends_ids"`
	Language        string    `json:"language,omitempty" sql:"language"`
	Theme           string    `json:"theme,omitempty" sql:"theme"`
	FirstName       *string   `json:"firstName,omitempty" sql:"first_name"`
	LastName        *string   `json:"lastName,omitempty" sql:"last_name"`
	Sessions        []string  `json:"sessions,omitempty" sql:"sessions"`
	CreateTime      time.Time `json:"createTime,omitempty" sql:"create_time"`
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
		Sessions:        user.Sessions,
		CreateTime:      user.CreateTime,
	}
}

type User struct {
	Name       string    `json:"name,omitempty" sql:"name"`
	Biography  *string   `json:"biography,omitempty" sql:"biography"`
	Role       string    `json:"role,omitempty" sql:"role"`
	FriendsIds []string  `json:"friendsIds,omitempty" sql:"friends_ids"`
	FirstName  *string   `json:"firstName,omitempty" sql:"first_name"`
	LastName   *string   `json:"lastName,omitempty" sql:"last_name"`
	CreateTime time.Time `json:"createTime,omitempty" sql:"create_time"`
}

func TransformToUser(user *ent.User) *User {
	return &User{
		Name:       user.Name,
		Biography:  user.Biography,
		Role:       user.Role,
		FriendsIds: user.FriendsIds,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		CreateTime: user.CreateTime,
	}
}
