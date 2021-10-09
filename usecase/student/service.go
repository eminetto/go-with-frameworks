package student

import (
	"context"
	"fmt"
	"github.com/eminetto/go-with-frameworks/ent"
	"github.com/eminetto/go-with-frameworks/ent/student"
	"github.com/google/uuid"
)

type Service struct {
	ctx context.Context
	client *ent.Client
}

//NewService create a new service
func NewService(ctx context.Context, client *ent.Client) *Service {
	return &Service{
		ctx,
		client,
	}
}

func (s *Service) Get(id uuid.UUID) (*ent.Student, error) {
	st, err := s.client.Student.
		Query().
		Where(student.ID(id)).
		Only(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying student: %w", err)
	}
	return st, nil
}

func (s *Service) Search(name string) ([]*ent.Student, error) {
	st, err := s.client.Student.
		Query().
		Where(student.Or(student.FirstNameContains(name), student.LastNameContains(name))).
		All(s.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying student: %w", err)
	}
	return st, nil
}

func (s *Service) List() ([]*ent.Student, error) {
	st, err := s.client.Student.
		Query().
		Where().
		All(s.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed listing student: %w", err)
	}
	return st, nil

}

func (s *Service) Create(e *ent.Student) error {
	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}
	_, err = tx.Student.
		Create().
		SetID(e.ID).
		SetEmail(e.Email).
		SetFirstName(e.FirstName).
		SetLastName(e.LastName).
		Save(s.ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("failed creating the student: %w", err))
	}

	return tx.Commit()
}

func (s *Service) Update(e *ent.Student) error {
	return nil
}

func (s *Service) Delete(id uuid.UUID) error {
	return nil
}

func (s *Service) GetCourses(id uuid.UUID) ([]*ent.Course, error) {
	return nil, nil

}

func (s *Service) AddCourse(st *ent.Student, c *ent.Course) error {
	return nil
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

