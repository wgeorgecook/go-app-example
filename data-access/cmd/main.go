package main

import (
	"log"

	petapisv1 "buf.build/gen/go/wgeorgecook/go-app-example/protocolbuffers/go/wgeorgecook/petapis/v1"
	"github.com/wgeorgecook/go-app-example/data-access/internal/pkg/db"
)

const POSTGRES_URI = "postgres://postgres:postgres@postgres:5432/pets?sslmode=disable"

func init() {
	if err := db.Migrate(POSTGRES_URI); err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Hello World! This is the data access layer.")

	conn, err := db.Connect(POSTGRES_URI)
	if err != nil {
		panic(err)
	}
	pets, err := db.Query(conn, []petapisv1.Query{})
	if err != nil {
		panic(err)
	}

	log.Printf("pets: %+v\n", pets)
}
