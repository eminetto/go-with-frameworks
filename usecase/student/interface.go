package student

import (
	"github.com/eminetto/go-with-frameworks/ent/schema"
	"github.com/google/uuid"
)
//Reader interface
type Reader interface {
	Get(id uuid.UUID) (*schema.Student, error)
	Search(name string) ([]*schema.Student, error)
	List() ([]*schema.Student, error)
}

//Writer user writer
type Writer interface {
	Create(e *schema.Student) error
	Update(e *schema.Student) error
	Delete(id uuid.UUID) error
}

//UseCase interface
type UseCase interface {
	Reader
	Writer
}