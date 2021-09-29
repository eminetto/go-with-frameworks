// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eminetto/go-with-frameworks/ent/course"
	"github.com/eminetto/go-with-frameworks/ent/lesson"
	"github.com/google/uuid"
)

// LessonCreate is the builder for creating a Lesson entity.
type LessonCreate struct {
	config
	mutation *LessonMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (lc *LessonCreate) SetName(s string) *LessonCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetID sets the "id" field.
func (lc *LessonCreate) SetID(u uuid.UUID) *LessonCreate {
	lc.mutation.SetID(u)
	return lc
}

// SetCourseID sets the "course" edge to the Course entity by ID.
func (lc *LessonCreate) SetCourseID(id uuid.UUID) *LessonCreate {
	lc.mutation.SetCourseID(id)
	return lc
}

// SetNillableCourseID sets the "course" edge to the Course entity by ID if the given value is not nil.
func (lc *LessonCreate) SetNillableCourseID(id *uuid.UUID) *LessonCreate {
	if id != nil {
		lc = lc.SetCourseID(*id)
	}
	return lc
}

// SetCourse sets the "course" edge to the Course entity.
func (lc *LessonCreate) SetCourse(c *Course) *LessonCreate {
	return lc.SetCourseID(c.ID)
}

// Mutation returns the LessonMutation object of the builder.
func (lc *LessonCreate) Mutation() *LessonMutation {
	return lc.mutation
}

// Save creates the Lesson in the database.
func (lc *LessonCreate) Save(ctx context.Context) (*Lesson, error) {
	var (
		err  error
		node *Lesson
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LessonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LessonCreate) SaveX(ctx context.Context) *Lesson {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LessonCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LessonCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LessonCreate) defaults() {
	if _, ok := lc.mutation.ID(); !ok {
		v := lesson.DefaultID()
		lc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LessonCreate) check() error {
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	return nil
}

func (lc *LessonCreate) sqlSave(ctx context.Context) (*Lesson, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
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

func (lc *LessonCreate) createSpec() (*Lesson, *sqlgraph.CreateSpec) {
	var (
		_node = &Lesson{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lesson.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: lesson.FieldID,
			},
		}
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lesson.FieldName,
		})
		_node.Name = value
	}
	if nodes := lc.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lesson.CourseTable,
			Columns: []string{lesson.CourseColumn},
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
		_node.course_lessons = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LessonCreateBulk is the builder for creating many Lesson entities in bulk.
type LessonCreateBulk struct {
	config
	builders []*LessonCreate
}

// Save creates the Lesson entities in the database.
func (lcb *LessonCreateBulk) Save(ctx context.Context) ([]*Lesson, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Lesson, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LessonMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LessonCreateBulk) SaveX(ctx context.Context) []*Lesson {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LessonCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LessonCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}