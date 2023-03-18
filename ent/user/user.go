// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldUserName holds the string denoting the user_name field in the database.
	FieldUserName = "user_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldBiography holds the string denoting the biography field in the database.
	FieldBiography = "biography"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldFriendsIds holds the string denoting the friends_ids field in the database.
	FieldFriendsIds = "friends_ids"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldTheme holds the string denoting the theme field in the database.
	FieldTheme = "theme"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldUserName,
	FieldEmail,
	FieldPasswordHash,
	FieldBiography,
	FieldRole,
	FieldAvatar,
	FieldFriendsIds,
	FieldLanguage,
	FieldTheme,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// UserNameValidator is a validation for the "user_name" field. It is called by the builders before save.
	UserNameValidator func(string) error
	// EmailValidator is a validation for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// BiographyValidator is a validation for the "biography" field. It is called by the builders before save.
	BiographyValidator func(string) error
	// NameValidator is a validation for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// Role defines the type for the "role" enum field.
type Role string

// RoleUSER is the default value of the Role enum.
const DefaultRole = RoleUSER

// Role values.
const (
	RoleUSER  Role = "USER"
	RoleADMIN Role = "ADMIN"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validation for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleUSER, RoleADMIN:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

// Language defines the type for the "language" enum field.
type Language string

// LanguageEN is the default value of the Language enum.
const DefaultLanguage = LanguageEN

// Language values.
const (
	LanguageEN Language = "EN"
	LanguageRU Language = "RU"
)

func (l Language) String() string {
	return string(l)
}

// LanguageValidator is a validation for the "language" field enum values. It is called by the builders before save.
func LanguageValidator(l Language) error {
	switch l {
	case LanguageEN, LanguageRU:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for language field: %q", l)
	}
}

// Theme defines the type for the "theme" enum field.
type Theme string

// ThemeSYSTEM is the default value of the Theme enum.
const DefaultTheme = ThemeSYSTEM

// Theme values.
const (
	ThemeWHITE  Theme = "WHITE"
	ThemeDARK   Theme = "DARK"
	ThemeSYSTEM Theme = "SYSTEM"
)

func (t Theme) String() string {
	return string(t)
}

// ThemeValidator is a validation for the "theme" field enum values. It is called by the builders before save.
func ThemeValidator(t Theme) error {
	switch t {
	case ThemeWHITE, ThemeDARK, ThemeSYSTEM:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for theme field: %q", t)
	}
}
