// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/wtkeqrf0/you-together/ent/room"
	"github.com/wtkeqrf0/you-together/ent/schema"
	"github.com/wtkeqrf0/you-together/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	roomMixin := schema.Room{}.Mixin()
	roomMixinFields0 := roomMixin[0].Fields()
	_ = roomMixinFields0
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescCreateTime is the schema descriptor for create_time field.
	roomDescCreateTime := roomMixinFields0[0].Descriptor()
	// room.DefaultCreateTime holds the default value on creation for the create_time field.
	room.DefaultCreateTime = roomDescCreateTime.Default.(func() time.Time)
	// roomDescUpdateTime is the schema descriptor for update_time field.
	roomDescUpdateTime := roomMixinFields0[1].Descriptor()
	// room.DefaultUpdateTime holds the default value on creation for the update_time field.
	room.DefaultUpdateTime = roomDescUpdateTime.Default.(func() time.Time)
	// room.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	room.UpdateDefaultUpdateTime = roomDescUpdateTime.UpdateDefault.(func() time.Time)
	// roomDescName is the schema descriptor for name field.
	roomDescName := roomFields[0].Descriptor()
	// room.DefaultName holds the default value on creation for the name field.
	room.DefaultName = roomDescName.Default.(func() string)
	// room.NameValidator is a validator for the "name" field. It is called by the builders before save.
	room.NameValidator = func() func(string) error {
		validators := roomDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// roomDescCustomName is the schema descriptor for custom_name field.
	roomDescCustomName := roomFields[1].Descriptor()
	// room.CustomNameValidator is a validator for the "custom_name" field. It is called by the builders before save.
	room.CustomNameValidator = func() func(string) error {
		validators := roomDescCustomName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(custom_name string) error {
			for _, fn := range fns {
				if err := fn(custom_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// roomDescOwnerID is the schema descriptor for owner_id field.
	roomDescOwnerID := roomFields[2].Descriptor()
	// room.OwnerIDValidator is a validator for the "owner_id" field. It is called by the builders before save.
	room.OwnerIDValidator = roomDescOwnerID.Validators[0].(func(int) error)
	// roomDescHasChat is the schema descriptor for has_chat field.
	roomDescHasChat := roomFields[5].Descriptor()
	// room.DefaultHasChat holds the default value on creation for the has_chat field.
	room.DefaultHasChat = roomDescHasChat.Default.(bool)
	// roomDescDescription is the schema descriptor for description field.
	roomDescDescription := roomFields[6].Descriptor()
	// room.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	room.DescriptionValidator = roomDescDescription.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(func() string)
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescIsEmailVerified is the schema descriptor for is_email_verified field.
	userDescIsEmailVerified := userFields[2].Descriptor()
	// user.DefaultIsEmailVerified holds the default value on creation for the is_email_verified field.
	user.DefaultIsEmailVerified = userDescIsEmailVerified.Default.(bool)
	// userDescBiography is the schema descriptor for biography field.
	userDescBiography := userFields[4].Descriptor()
	// user.BiographyValidator is a validator for the "biography" field. It is called by the builders before save.
	user.BiographyValidator = userDescBiography.Validators[0].(func(string) error)
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[5].Descriptor()
	// user.DefaultRole holds the default value on creation for the role field.
	user.DefaultRole = userDescRole.Default.(string)
	// userDescLanguage is the schema descriptor for language field.
	userDescLanguage := userFields[7].Descriptor()
	// user.DefaultLanguage holds the default value on creation for the language field.
	user.DefaultLanguage = userDescLanguage.Default.(string)
	// userDescTheme is the schema descriptor for theme field.
	userDescTheme := userFields[8].Descriptor()
	// user.DefaultTheme holds the default value on creation for the theme field.
	user.DefaultTheme = userDescTheme.Default.(string)
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[9].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = func() func(string) error {
		validators := userDescFirstName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(first_name string) error {
			for _, fn := range fns {
				if err := fn(first_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[10].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = func() func(string) error {
		validators := userDescLastName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(last_name string) error {
			for _, fn := range fns {
				if err := fn(last_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
