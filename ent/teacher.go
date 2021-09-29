// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/eminetto/go-with-frameworks/ent/teacher"
	"github.com/google/uuid"
)

// Teacher is the model entity for the Teacher schema.
type Teacher struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeacherQuery when eager-loading is set.
	Edges TeacherEdges `json:"edges"`
}

// TeacherEdges holds the relations/edges for other nodes in the graph.
type TeacherEdges struct {
	// Courses holds the value of the courses edge.
	Courses []*Course `json:"courses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CoursesOrErr returns the Courses value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) CoursesOrErr() ([]*Course, error) {
	if e.loadedTypes[0] {
		return e.Courses, nil
	}
	return nil, &NotLoadedError{edge: "courses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Teacher) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case teacher.FieldEmail, teacher.FieldFirstName, teacher.FieldLastName:
			values[i] = new(sql.NullString)
		case teacher.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Teacher", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Teacher fields.
func (t *Teacher) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case teacher.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case teacher.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				t.Email = value.String
			}
		case teacher.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				t.FirstName = value.String
			}
		case teacher.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				t.LastName = value.String
			}
		}
	}
	return nil
}

// QueryCourses queries the "courses" edge of the Teacher entity.
func (t *Teacher) QueryCourses() *CourseQuery {
	return (&TeacherClient{config: t.config}).QueryCourses(t)
}

// Update returns a builder for updating this Teacher.
// Note that you need to call Teacher.Unwrap() before calling this method if this Teacher
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Teacher) Update() *TeacherUpdateOne {
	return (&TeacherClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Teacher entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Teacher) Unwrap() *Teacher {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Teacher is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Teacher) String() string {
	var builder strings.Builder
	builder.WriteString("Teacher(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", email=")
	builder.WriteString(t.Email)
	builder.WriteString(", first_name=")
	builder.WriteString(t.FirstName)
	builder.WriteString(", last_name=")
	builder.WriteString(t.LastName)
	builder.WriteByte(')')
	return builder.String()
}

// Teachers is a parsable slice of Teacher.
type Teachers []*Teacher

func (t Teachers) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
