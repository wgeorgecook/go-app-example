package server

import (
	"context"
	"log"

	petapisv1 "buf.build/gen/go/wgeorgecook/go-app-example/protocolbuffers/go/wgeorgecook/petapis/v1"
)

func (petServiceServer) GetPets(ctx context.Context, req *petapisv1.GetPetsRequest) (*petapisv1.GetPetsResponse, error) {
	log.Println("start get pets")
	defer log.Println("get pets complete")
	return nil, nil
}
