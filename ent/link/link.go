// Code generated by entc, DO NOT EDIT.

package link

import (
	"time"
)

const (
	// Label holds the string label denoting the link type in the database.
	Label = "link"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLongURI holds the string denoting the long_uri field in the database.
	FieldLongURI = "long_uri"
	// FieldAccessedTimes holds the string denoting the accessed_times field in the database.
	FieldAccessedTimes = "accessed_times"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the link in the database.
	Table = "links"
)

// Columns holds all SQL columns for link fields.
var Columns = []string{
	FieldID,
	FieldLongURI,
	FieldAccessedTimes,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// LongURIValidator is a validator for the "long_uri" field. It is called by the builders before save.
	LongURIValidator func(string) error
	// DefaultAccessedTimes holds the default value on creation for the "accessed_times" field.
	DefaultAccessedTimes int
	// AccessedTimesValidator is a validator for the "accessed_times" field. It is called by the builders before save.
	AccessedTimesValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
