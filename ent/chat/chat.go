// Code generated by ent, DO NOT EDIT.

package chat

import (
	"time"
)

const (
	// Label holds the string label denoting the chat type in the database.
	Label = "chat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// EdgeRoom holds the string denoting the room edge name in mutations.
	EdgeRoom = "room"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the chat in the database.
	Table = "chats"
	// RoomTable is the table that holds the room relation/edge.
	RoomTable = "rooms"
	// RoomInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	RoomInverseTable = "rooms"
	// RoomColumn is the table column denoting the room relation/edge.
	RoomColumn = "room_chat"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "users"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "chat_user"
)

// Columns holds all SQL columns for chat fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldMessage,
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
	// MessageValidator is a validator for the "message" field. It is called by the builders before save.
	MessageValidator func(string) error
)