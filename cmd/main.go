package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/dmikhr/auth/pkg/user_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedAuthV1Server
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

// Create - создание нового пользователя
func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("User data: %s | %s | %s | %s | %v", req.GetName(), req.GetEmail(), req.GetPassword(), req.GetPasswordConfirm(), req.GetRole())
	return &desc.CreateResponse{}, nil
}

// Get - получить пользователя по id
func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Get id: %d", req.GetId())
	return &desc.GetResponse{}, nil
}

// Update - обновление данных пользователя
func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update data: %d | %s | %s | %v", req.GetId(), req.GetName(), req.GetEmail(), req.GetRole())
	return &emptypb.Empty{}, nil
}

// Delete - удаление пользователя
func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Chat to delete id: %d", req.GetId())
	return &emptypb.Empty{}, nil
}
