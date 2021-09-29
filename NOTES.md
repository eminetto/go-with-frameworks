# Ent

## Install

    go get entgo.io/ent/cmd/ent

## Creating entities

    go run entgo.io/ent/cmd/ent init Student Course Lesson Teacher

## Configure entities

- [x] Add fields

- [x] Add relationships 

Execute: 

    go mod tidy
    go generate ./ent 
    
to generate the schema files

Let's create an main.go just to test the tabel generation.


Install PostgreSQL lib:

    go get github.com/lib/pq

main.go content:

```go

package main

import (
	"context"
	"log"

	"github.com/eminetto/go-with-frameworks/ent"
	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=prest password=prest dbname=prest sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

```

Execute:

    go run main.go

And check the tables created in the database


# UseCases

Create directory:

    mkdir -p usecase/student

Create the interface in `usecase/student/interface.go`:

```go
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
```


