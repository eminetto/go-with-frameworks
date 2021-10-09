package student_test

import (
	"context"
	"github.com/eminetto/go-with-frameworks/ent"
	"github.com/eminetto/go-with-frameworks/ent/enttest"
	"github.com/eminetto/go-with-frameworks/usecase/student"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newStudent() *ent.Student {
	return &ent.Student{
		ID:        uuid.New(),
		Email:     "sjobs@apple.com",
		FirstName: "Steve",
		LastName:  "Jobs",
	}
}

func Test_Create(t *testing.T) {
	cl := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer cl.Close()
	service := student.NewService(context.Background(), cl)
	ns := newStudent()
	err := service.Create(ns)
	assert.Nil(t, err)
	err = service.Create(ns)
	assert.Errorf(t, err,"UNIQUE constraint failed")

}

func Test_Get(t *testing.T) {
	cl := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer cl.Close()
	service := student.NewService(context.Background(), cl)
	ns := newStudent()
	err := service.Create(ns)
	assert.Nil(t, err)
	st, err := service.Get(ns.ID)
	assert.Nil(t, err)
	assert.Equal(t, ns.Email, st.Email)
}

func Test_Search(t *testing.T) {
	cl := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer cl.Close()
	service := student.NewService(context.Background(), cl)
	ns := newStudent()
	err := service.Create(ns)
	assert.Nil(t, err)
	tests := []struct {
		Search string
		ExpectedCount int
	}{
		{
			Search: "steve",
			ExpectedCount: 1,
		},
		{
			Search: "jobs",
			ExpectedCount: 1,
		},
		{
			Search: "Steve",
			ExpectedCount: 1,
		},
		{
			Search: "Jobs",
			ExpectedCount: 1,
		},
		{
			Search: "bill",
			ExpectedCount: 0,
		},

	}
	for _, tt := range tests {
		st, err := service.Search(tt.Search)
		assert.Equal(t, tt.ExpectedCount, len(st))
		assert.Nil(t, err)
		if tt.ExpectedCount > 0 {
			assert.Equal(t, ns.Email, st[0].Email)
		}
	}
}

func Test_List(t *testing.T) {
	cl := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer cl.Close()
	service := student.NewService(context.Background(), cl)
	ns := newStudent()
	err := service.Create(ns)
	assert.Nil(t, err)
	st, err := service.List()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(st))
	assert.Equal(t, ns.Email, st[0].Email)
}

