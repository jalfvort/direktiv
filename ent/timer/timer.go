// Code generated by entc, DO NOT EDIT.

package timer

const (
	// Label holds the string label denoting the timer type in the database.
	Label = "timer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFn holds the string denoting the fn field in the database.
	FieldFn = "fn"
	// FieldCron holds the string denoting the cron field in the database.
	FieldCron = "cron"
	// FieldOne holds the string denoting the one field in the database.
	FieldOne = "one"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// Table holds the table name of the timer in the database.
	Table = "timers"
)

// Columns holds all SQL columns for timer fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldFn,
	FieldCron,
	FieldOne,
	FieldData,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// FnValidator is a validator for the "fn" field. It is called by the builders before save.
	FnValidator func(string) error
)
