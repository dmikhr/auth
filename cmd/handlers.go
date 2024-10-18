package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dmikhr/auth/internal/data"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/dmikhr/auth/pkg/user_v1"
)

// CreateUser - создание нового пользователя
func (s *server) CreateUser(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("User data: %s | %s | %s | %s | %v", req.GetName(), req.GetEmail(), req.GetPassword(), req.GetPasswordConfirm(), req.GetRole())
	var id int64
	var err error
	id, err = s.models.User.Create(&data.UserNew{
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
		Role:      data.RoleToStr(req.GetRole()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()})
	if err != nil {
		fmt.Println("Error creating user:", err)
		return nil, err
	}
	return &desc.CreateResponse{Id: id}, nil
}

// GetUser - получить пользователя по id
func (s *server) GetUser(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Get id: %d", req.GetId())
	var user data.UserRetrieve
	var err error
	user, err = s.models.User.Get(req.GetId())
	if err != nil {
		fmt.Println("Error getting user:", err)
		return nil, err
	}
	return &desc.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      data.StrToRole(user.Role),
		CreatedAt: data.TimeToTimestamppb(user.CreatedAt),
		UpdatedAt: data.TimeToTimestamppb(user.UpdatedAt)}, nil
}

// UpdateUser - обновление данных пользователя
func (s *server) UpdateUser(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update data: %d | %s | %s | %v", req.GetId(), req.GetName(), req.GetEmail(), req.GetRole())
	err := s.models.User.Update(&data.UserEdit{
		ID:        req.GetId(),
		Name:      req.GetName().Value,
		Email:     req.GetEmail(),
		Role:      data.RoleToStr(req.GetRole()),
		UpdatedAt: time.Now()})
	if err != nil {
		fmt.Println("Error updating user:", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// DeleteUser - удаление пользователя
func (s *server) DeleteUser(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Chat to delete id: %d", req.GetId())
	err := s.models.User.Delete(req.GetId())
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
