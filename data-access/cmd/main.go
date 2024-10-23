package main

import (
	"log"

	"github.com/wgeorgecook/go-app-example/data-access/internal/pkg/db"
)

func init() {
	if err := db.Migrate(); err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Hello World! This is the data access layer.")
}
