package main

import (
	"context"
	"fmt"
	"log"
	pb "myrpc/bigboss/student"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNetworkServer
}

func (s *server) WelcomeStudent(ctx context.Context, stud *pb.Student) (*pb.Welcome, error) {
	log.Printf("received: %v\n", stud.String())
	log.Printf("name: %v\n", stud.GetName())
	log.Printf("age: %v\n", stud.GetAge())
	log.Printf("weight: %v\n", stud.GetWeight())
	message := fmt.Sprintf("welcome %s", stud.GetName())
	return &pb.Welcome{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterNetworkServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
