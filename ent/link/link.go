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
	// FieldShortPath holds the string denoting the short_path field in the database.
	FieldShortPath = "short_path"
	// FieldLongURI holds the string denoting the long_uri field in the database.
	FieldLongURI = "long_uri"
	// FieldAccessedTimes holds the string denoting the accessed_times field in the database.
	FieldAccessedTimes = "accessed_times"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the link in the database.
	Table = "links"
)

// Columns holds all SQL columns for link fields.
var Columns = []string{
	FieldID,
	FieldShortPath,
	FieldLongURI,
	FieldAccessedTimes,
	FieldCreatedAt,
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
	// ShortPathValidator is a validator for the "short_path" field. It is called by the builders before save.
	ShortPathValidator func(string) error
	// LongURIValidator is a validator for the "long_uri" field. It is called by the builders before save.
	LongURIValidator func(string) error
	// DefaultAccessedTimes holds the default value on creation for the "accessed_times" field.
	DefaultAccessedTimes int
	// AccessedTimesValidator is a validator for the "accessed_times" field. It is called by the builders before save.
	AccessedTimesValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
