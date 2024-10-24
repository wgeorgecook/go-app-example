package server

import (
	"fmt"
	"log"
	"net"

	"buf.build/gen/go/wgeorgecook/go-app-example/grpc/go/wgeorgecook/petapis/v1/petapisv1grpc"
	"google.golang.org/grpc"
)

var srv *grpc.Server

// petServiceServer embeds the UnimplentedPetServiceServer
// for forwards compatibility out of the box.
type petServiceServer struct {
	petapisv1grpc.UnimplementedPetServiceServer
}

// Start spings up a new GRPC server instance and
// listens for incoming connects with the provided configuration
// options. This is a blocking operation and the caller
// is responsible for shutting down the server.
func Start(listenerPort string, grpcOpts ...grpc.ServerOption) error {
	if srv != nil {
		log.Println("server already initiated")
		return nil
	}
	log.Println("starting GRPC server...")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", listenerPort))
	if err != nil {
		return err
	}
	srv = grpc.NewServer(grpcOpts...)
	petapisv1grpc.RegisterPetServiceServer(srv, &petServiceServer{})
	log.Printf("start listening on port %s\n", listenerPort)
	if err := srv.Serve(listener); err != nil {
		return err
	}

	log.Println("listen complete")
	return nil
}

// Shutdown calls the underlying *grpc.Server.GracefulShutdown()
func Shutdown() {
	if srv == nil {
		log.Println("server is nil, returning")
		return
	}

	srv.GracefulStop()
}
