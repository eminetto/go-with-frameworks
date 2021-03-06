// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/eminetto/go-with-frameworks/ent/course"
	"github.com/eminetto/go-with-frameworks/ent/lesson"
	"github.com/eminetto/go-with-frameworks/ent/schema"
	"github.com/eminetto/go-with-frameworks/ent/student"
	"github.com/eminetto/go-with-frameworks/ent/teacher"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	courseFields := schema.Course{}.Fields()
	_ = courseFields
	// courseDescID is the schema descriptor for id field.
	courseDescID := courseFields[0].Descriptor()
	// course.DefaultID holds the default value on creation for the id field.
	course.DefaultID = courseDescID.Default.(func() uuid.UUID)
	lessonFields := schema.Lesson{}.Fields()
	_ = lessonFields
	// lessonDescID is the schema descriptor for id field.
	lessonDescID := lessonFields[0].Descriptor()
	// lesson.DefaultID holds the default value on creation for the id field.
	lesson.DefaultID = lessonDescID.Default.(func() uuid.UUID)
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescFirstName is the schema descriptor for first_name field.
	studentDescFirstName := studentFields[2].Descriptor()
	// student.DefaultFirstName holds the default value on creation for the first_name field.
	student.DefaultFirstName = studentDescFirstName.Default.(string)
	// studentDescLastName is the schema descriptor for last_name field.
	studentDescLastName := studentFields[3].Descriptor()
	// student.DefaultLastName holds the default value on creation for the last_name field.
	student.DefaultLastName = studentDescLastName.Default.(string)
	// studentDescID is the schema descriptor for id field.
	studentDescID := studentFields[0].Descriptor()
	// student.DefaultID holds the default value on creation for the id field.
	student.DefaultID = studentDescID.Default.(func() uuid.UUID)
	teacherFields := schema.Teacher{}.Fields()
	_ = teacherFields
	// teacherDescFirstName is the schema descriptor for first_name field.
	teacherDescFirstName := teacherFields[2].Descriptor()
	// teacher.DefaultFirstName holds the default value on creation for the first_name field.
	teacher.DefaultFirstName = teacherDescFirstName.Default.(string)
	// teacherDescLastName is the schema descriptor for last_name field.
	teacherDescLastName := teacherFields[3].Descriptor()
	// teacher.DefaultLastName holds the default value on creation for the last_name field.
	teacher.DefaultLastName = teacherDescLastName.Default.(string)
	// teacherDescID is the schema descriptor for id field.
	teacherDescID := teacherFields[0].Descriptor()
	// teacher.DefaultID holds the default value on creation for the id field.
	teacher.DefaultID = teacherDescID.Default.(func() uuid.UUID)
}
