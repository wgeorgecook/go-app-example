package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wgeorgecook/go-app-example/data-access/internal/pkg/db"
	"github.com/wgeorgecook/go-app-example/data-access/server"
)

const POSTGRES_URI = "postgres://postgres:postgres@postgres:5432/pets?sslmode=disable"
const GRPC_PORT = "8088"

func init() {
	if err := db.Migrate(POSTGRES_URI); err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Hello World! This is the data access layer.")

	if err := db.Connect(POSTGRES_URI); err != nil {
		panic(err)
	}

	ctx := context.Background()
	go server.Start(GRPC_PORT)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	log.Println("server is running, CTRL + C to shutdown")
	<-exit
	log.Println("shutdown received")
	if err := db.Shutdown(ctx); err != nil {
		log.Println("could not close db connection: ", err.Error())
	}

	server.Shutdown()
	log.Println("shutdown complete")
}
