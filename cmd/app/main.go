package main

import (
	"go_ddd/internal/app/controller"
	"go_ddd/internal/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	task := &controller.TaskController{}
	pb.RegisterTaskServiceServer(s, task)

	log.Printf("gRPC server listening on " + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf(err.Error())
	}
}
