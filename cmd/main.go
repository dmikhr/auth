package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/dmikhr/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedAuthV1Server
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("User data: %s | %s | %s | %s | %v", req.Name, req.Email, req.Password, req.PasswordConfirm, req.Role)
	return &desc.CreateResponse{}, nil
}

func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Get id: %d", req.Id)
	return &desc.GetResponse{}, nil
}

func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update data: %d | %s | %s | %v", req.Id, req.Name, req.Email, req.Role)
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Chat to delete id: %d", req.Id)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
