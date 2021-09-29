// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eminetto/go-with-frameworks/ent/course"
	"github.com/eminetto/go-with-frameworks/ent/student"
	"github.com/google/uuid"
)

// StudentCreate is the builder for creating a Student entity.
type StudentCreate struct {
	config
	mutation *StudentMutation
	hooks    []Hook
}

// SetEmail sets the "email" field.
func (sc *StudentCreate) SetEmail(s string) *StudentCreate {
	sc.mutation.SetEmail(s)
	return sc
}

// SetFirstName sets the "first_name" field.
func (sc *StudentCreate) SetFirstName(s string) *StudentCreate {
	sc.mutation.SetFirstName(s)
	return sc
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (sc *StudentCreate) SetNillableFirstName(s *string) *StudentCreate {
	if s != nil {
		sc.SetFirstName(*s)
	}
	return sc
}

// SetLastName sets the "last_name" field.
func (sc *StudentCreate) SetLastName(s string) *StudentCreate {
	sc.mutation.SetLastName(s)
	return sc
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (sc *StudentCreate) SetNillableLastName(s *string) *StudentCreate {
	if s != nil {
		sc.SetLastName(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *StudentCreate) SetID(u uuid.UUID) *StudentCreate {
	sc.mutation.SetID(u)
	return sc
}

// AddCourseIDs adds the "courses" edge to the Course entity by IDs.
func (sc *StudentCreate) AddCourseIDs(ids ...uuid.UUID) *StudentCreate {
	sc.mutation.AddCourseIDs(ids...)
	return sc
}

// AddCourses adds the "courses" edges to the Course entity.
func (sc *StudentCreate) AddCourses(c ...*Course) *StudentCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddCourseIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (sc *StudentCreate) Mutation() *StudentMutation {
	return sc.mutation
}

// Save creates the Student in the database.
func (sc *StudentCreate) Save(ctx context.Context) (*Student, error) {
	var (
		err  error
		node *Student
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StudentCreate) SaveX(ctx context.Context) *Student {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StudentCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StudentCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StudentCreate) defaults() {
	if _, ok := sc.mutation.FirstName(); !ok {
		v := student.DefaultFirstName
		sc.mutation.SetFirstName(v)
	}
	if _, ok := sc.mutation.LastName(); !ok {
		v := student.DefaultLastName
		sc.mutation.SetLastName(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := student.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StudentCreate) check() error {
	if _, ok := sc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "email"`)}
	}
	if _, ok := sc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "first_name"`)}
	}
	if _, ok := sc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "last_name"`)}
	}
	return nil
}

func (sc *StudentCreate) sqlSave(ctx context.Context) (*Student, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (sc *StudentCreate) createSpec() (*Student, *sqlgraph.CreateSpec) {
	var (
		_node = &Student{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: student.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: student.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := sc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := sc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldLastName,
		})
		_node.LastName = value
	}
	if nodes := sc.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   student.CoursesTable,
			Columns: student.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StudentCreateBulk is the builder for creating many Student entities in bulk.
type StudentCreateBulk struct {
	config
	builders []*StudentCreate
}

// Save creates the Student entities in the database.
func (scb *StudentCreateBulk) Save(ctx context.Context) ([]*Student, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Student, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StudentCreateBulk) SaveX(ctx context.Context) []*Student {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StudentCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StudentCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}