// Code generated by entc, DO NOT EDIT.

package lesson

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the lesson type in the database.
	Label = "lesson"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeCourse holds the string denoting the course edge name in mutations.
	EdgeCourse = "course"
	// Table holds the table name of the lesson in the database.
	Table = "lessons"
	// CourseTable is the table that holds the course relation/edge.
	CourseTable = "lessons"
	// CourseInverseTable is the table name for the Course entity.
	// It exists in this package in order to avoid circular dependency with the "course" package.
	CourseInverseTable = "courses"
	// CourseColumn is the table column denoting the course relation/edge.
	CourseColumn = "course_lessons"
)

// Columns holds all SQL columns for lesson fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "lessons"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"course_lessons",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
