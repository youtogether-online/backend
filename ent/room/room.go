// Code generated by ent, DO NOT EDIT.

package room

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the room type in the database.
	Label = "room"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCustomName holds the string denoting the custom_name field in the database.
	FieldCustomName = "custom_name"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldPrivacy holds the string denoting the privacy field in the database.
	FieldPrivacy = "privacy"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldHasChat holds the string denoting the has_chat field in the database.
	FieldHasChat = "has_chat"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the room in the database.
	Table = "rooms"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_rooms"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
)

// Columns holds all SQL columns for room fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldCustomName,
	FieldOwnerID,
	FieldPrivacy,
	FieldPasswordHash,
	FieldHasChat,
	FieldDescription,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "room_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/wtkeqrf0/you-together/ent/runtime"
var (
	Hooks [2]ent.Hook
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CustomNameValidator is a validator for the "custom_name" field. It is called by the builders before save.
	CustomNameValidator func(string) error
	// OwnerIDValidator is a validator for the "owner_id" field. It is called by the builders before save.
	OwnerIDValidator func(int) error
	// DefaultPrivacy holds the default value on creation for the "privacy" field.
	DefaultPrivacy string
	// DefaultHasChat holds the default value on creation for the "has_chat" field.
	DefaultHasChat bool
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)
