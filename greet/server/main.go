package main

import (
	"log"
	"net"

	pb "example.com/m/greet/proto"
	"google.golang.org/grpc"
)


var address = "0.0.0.0:50051";

type Address struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen ON: %v", err)
	}

	log.Printf("listening on %v", address)

	s :=grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}