// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Size: 32, Default: schema.Expr("'room' || setval(pg_get_serial_sequence('rooms','id'),nextval(pg_get_serial_sequence('rooms','id'))-1)")},
		{Name: "privacy", Type: field.TypeString, Default: "friends"},
		{Name: "password_hash", Type: field.TypeBytes, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 140},
		{Name: "user_room", Type: field.TypeInt, Unique: true},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:       "rooms",
		Columns:    RoomsColumns,
		PrimaryKey: []*schema.Column{RoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rooms_users_room",
				Columns:    []*schema.Column{RoomsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true, Default: schema.Expr("'user' || setval(pg_get_serial_sequence('users','id'),nextval(pg_get_serial_sequence('users','id'))-1)")},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "is_email_verified", Type: field.TypeBool, Default: false},
		{Name: "password_hash", Type: field.TypeBytes, Nullable: true},
		{Name: "biography", Type: field.TypeString, Nullable: true, Size: 512},
		{Name: "role", Type: field.TypeString, Default: "user"},
		{Name: "friends_ids", Type: field.TypeJSON, Nullable: true},
		{Name: "language", Type: field.TypeString, Default: "en"},
		{Name: "theme", Type: field.TypeString, Default: "system"},
		{Name: "first_name", Type: field.TypeString, Nullable: true, Size: 32},
		{Name: "last_name", Type: field.TypeString, Nullable: true, Size: 32},
		{Name: "sessions", Type: field.TypeJSON, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RoomsTable,
		UsersTable,
	}
)

func init() {
	RoomsTable.ForeignKeys[0].RefTable = UsersTable
}
