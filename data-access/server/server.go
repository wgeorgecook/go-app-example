package server

import (
	"log"
	"net"

	"buf.build/gen/go/wgeorgecook/go-app-example/grpc/go/wgeorgecook/petapis/v1/petapisv1grpc"
	"google.golang.org/grpc"
)

// petServiceServer embeds the UnimplentedPetServiceServer
// for forwards compatibility out of the box.
type petServiceServer struct {
	petapisv1grpc.UnimplementedPetServiceServer
}

// Start spings up a new GRPC server instance and
// listens for incoming connects with the provided configuration
// options. This is a blocking operation.
func Start(listenerPort int, grpcOpts ...grpc.ServerOption) error {
	log.Println("starting GRPC server...")
	listener, err := net.Listen("tcp", ":%d")
	if err != nil {
		return err
	}
	srv := grpc.NewServer(grpcOpts...)
	petapisv1grpc.RegisterPetServiceServer(srv, &petServiceServer{})
	log.Printf("start listening on port %d\n", listenerPort)
	if err := srv.Serve(listener); err != nil {
		return err
	}
	return nil
}
