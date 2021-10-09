package student

import (
	"github.com/eminetto/go-with-frameworks/ent"
	"github.com/google/uuid"
)
//Reader interface
type Reader interface {
	Get(id uuid.UUID) (*ent.Student, error)
	Search(name string) ([]*ent.Student, error)
	List() ([]*ent.Student, error)
}

//Writer user writer
type Writer interface {
	Create(e *ent.Student) error
	Update(e *ent.Student) error
	Delete(id uuid.UUID) error
	AddCourse(st *ent.Student, c *ent.Course) error
}

//UseCase interface
type UseCase interface {
	Reader
	Writer
	GetCourses(id uuid.UUID) ([]*ent.Course, error)
}