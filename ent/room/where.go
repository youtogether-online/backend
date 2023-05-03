// Code generated by ent, DO NOT EDIT.

package room

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/wtkeqrf0/you-together/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldName, v))
}

// CustomName applies equality check predicate on the "custom_name" field. It's identical to CustomNameEQ.
func CustomName(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldCustomName, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v int) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldOwnerID, v))
}

// PasswordHash applies equality check predicate on the "password_hash" field. It's identical to PasswordHashEQ.
func PasswordHash(v []byte) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldPasswordHash, v))
}

// HasChat applies equality check predicate on the "has_chat" field. It's identical to HasChatEQ.
func HasChat(v bool) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldHasChat, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldDescription, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Room {
	return predicate.Room(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Room {
	return predicate.Room(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Room {
	return predicate.Room(sql.FieldContainsFold(FieldName, v))
}

// CustomNameEQ applies the EQ predicate on the "custom_name" field.
func CustomNameEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldCustomName, v))
}

// CustomNameNEQ applies the NEQ predicate on the "custom_name" field.
func CustomNameNEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldCustomName, v))
}

// CustomNameIn applies the In predicate on the "custom_name" field.
func CustomNameIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldCustomName, vs...))
}

// CustomNameNotIn applies the NotIn predicate on the "custom_name" field.
func CustomNameNotIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldCustomName, vs...))
}

// CustomNameGT applies the GT predicate on the "custom_name" field.
func CustomNameGT(v string) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldCustomName, v))
}

// CustomNameGTE applies the GTE predicate on the "custom_name" field.
func CustomNameGTE(v string) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldCustomName, v))
}

// CustomNameLT applies the LT predicate on the "custom_name" field.
func CustomNameLT(v string) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldCustomName, v))
}

// CustomNameLTE applies the LTE predicate on the "custom_name" field.
func CustomNameLTE(v string) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldCustomName, v))
}

// CustomNameContains applies the Contains predicate on the "custom_name" field.
func CustomNameContains(v string) predicate.Room {
	return predicate.Room(sql.FieldContains(FieldCustomName, v))
}

// CustomNameHasPrefix applies the HasPrefix predicate on the "custom_name" field.
func CustomNameHasPrefix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasPrefix(FieldCustomName, v))
}

// CustomNameHasSuffix applies the HasSuffix predicate on the "custom_name" field.
func CustomNameHasSuffix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasSuffix(FieldCustomName, v))
}

// CustomNameIsNil applies the IsNil predicate on the "custom_name" field.
func CustomNameIsNil() predicate.Room {
	return predicate.Room(sql.FieldIsNull(FieldCustomName))
}

// CustomNameNotNil applies the NotNil predicate on the "custom_name" field.
func CustomNameNotNil() predicate.Room {
	return predicate.Room(sql.FieldNotNull(FieldCustomName))
}

// CustomNameEqualFold applies the EqualFold predicate on the "custom_name" field.
func CustomNameEqualFold(v string) predicate.Room {
	return predicate.Room(sql.FieldEqualFold(FieldCustomName, v))
}

// CustomNameContainsFold applies the ContainsFold predicate on the "custom_name" field.
func CustomNameContainsFold(v string) predicate.Room {
	return predicate.Room(sql.FieldContainsFold(FieldCustomName, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v int) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v int) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...int) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...int) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldOwnerID, vs...))
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v int) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldOwnerID, v))
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v int) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldOwnerID, v))
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v int) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldOwnerID, v))
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v int) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldOwnerID, v))
}

// PrivacyEQ applies the EQ predicate on the "privacy" field.
func PrivacyEQ(v Privacy) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldPrivacy, v))
}

// PrivacyNEQ applies the NEQ predicate on the "privacy" field.
func PrivacyNEQ(v Privacy) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldPrivacy, v))
}

// PrivacyIn applies the In predicate on the "privacy" field.
func PrivacyIn(vs ...Privacy) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldPrivacy, vs...))
}

// PrivacyNotIn applies the NotIn predicate on the "privacy" field.
func PrivacyNotIn(vs ...Privacy) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldPrivacy, vs...))
}

// PasswordHashEQ applies the EQ predicate on the "password_hash" field.
func PasswordHashEQ(v []byte) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldPasswordHash, v))
}

// PasswordHashNEQ applies the NEQ predicate on the "password_hash" field.
func PasswordHashNEQ(v []byte) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldPasswordHash, v))
}

// PasswordHashIn applies the In predicate on the "password_hash" field.
func PasswordHashIn(vs ...[]byte) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldPasswordHash, vs...))
}

// PasswordHashNotIn applies the NotIn predicate on the "password_hash" field.
func PasswordHashNotIn(vs ...[]byte) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldPasswordHash, vs...))
}

// PasswordHashGT applies the GT predicate on the "password_hash" field.
func PasswordHashGT(v []byte) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldPasswordHash, v))
}

// PasswordHashGTE applies the GTE predicate on the "password_hash" field.
func PasswordHashGTE(v []byte) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldPasswordHash, v))
}

// PasswordHashLT applies the LT predicate on the "password_hash" field.
func PasswordHashLT(v []byte) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldPasswordHash, v))
}

// PasswordHashLTE applies the LTE predicate on the "password_hash" field.
func PasswordHashLTE(v []byte) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldPasswordHash, v))
}

// PasswordHashIsNil applies the IsNil predicate on the "password_hash" field.
func PasswordHashIsNil() predicate.Room {
	return predicate.Room(sql.FieldIsNull(FieldPasswordHash))
}

// PasswordHashNotNil applies the NotNil predicate on the "password_hash" field.
func PasswordHashNotNil() predicate.Room {
	return predicate.Room(sql.FieldNotNull(FieldPasswordHash))
}

// HasChatEQ applies the EQ predicate on the "has_chat" field.
func HasChatEQ(v bool) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldHasChat, v))
}

// HasChatNEQ applies the NEQ predicate on the "has_chat" field.
func HasChatNEQ(v bool) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldHasChat, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Room {
	return predicate.Room(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Room {
	return predicate.Room(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Room {
	return predicate.Room(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Room {
	return predicate.Room(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Room {
	return predicate.Room(sql.FieldContainsFold(FieldDescription, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Room) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Room) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Room) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		p(s.Not())
	})
}
