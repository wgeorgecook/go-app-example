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
	// run db migrations either up or down depending
	// on environment
	if err := db.Migrate(POSTGRES_URI); err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Hello World! This is the data access layer.")

	// build database connection
	if err := db.Connect(POSTGRES_URI); err != nil {
		panic(err)
	}

	// start listening for incoming grpc connections
	go server.Start(GRPC_PORT)

	log.Println("server is running, CTRL + C to shutdown")

	// graceful shutdown logic
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGTERM, syscall.SIGINT)

	// block until signal is received
	<-shutdownChan
	log.Println("shutdown received")

	// disconnect database
	ctx := context.Background()
	if err := db.Shutdown(ctx); err != nil {
		log.Println("could not close db connection: ", err.Error())
	}

	// graceful close gprc server
	server.Shutdown()
	log.Println("shutdown complete, goodbye")
}
