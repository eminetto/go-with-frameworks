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
