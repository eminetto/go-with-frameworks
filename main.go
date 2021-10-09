package main

import (
	"context"
	"fmt"
	"github.com/eminetto/go-with-frameworks/usecase/student"
	"github.com/google/uuid"
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

	service := student.NewService(context.Background(), client)
	list, err := service.List()
	if err != nil {
		log.Fatalf("failed listing students: %v", err)
	}
	fmt.Println(list)
	uid := uuid.MustParse("2d5e5b6f-b943-40cd-8fc4-1eb8c1644bbb")
	st, err := service.Get(uid)
	if err != nil {
		log.Fatalf("failed getting student: %v", err)
	}
	fmt.Println(st)
	list, err = service.Search("minettos")
	if err != nil {
		log.Fatalf("failed searching students: %v", err)
	}
	fmt.Println(list)

	//newSt := &ent.Student{
	//	ID:        uuid.New(),
	//	Email:     "aminetto@gmail.com",
	//	FirstName: "Alice",
	//	LastName:  "Minetto",
	//}
	//err = service.Create(newSt)
	//if err != nil {
	//	log.Fatalf("failed creating student: %v", err)
	//}

}
