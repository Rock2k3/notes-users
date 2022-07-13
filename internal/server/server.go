package server

import (
	"context"
	NotesGrpcApi "github.com/Rock2k3/notes-grpc-api/generated-sources"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	"notes-users/config"
	"notes-users/internal/adapters"
	"notes-users/internal/domain/users"
)

type server struct {
	config *config.AppConfig
}

type grpcSrv struct {
	NotesGrpcApi.UnimplementedUserServiceServer
	config *config.AppConfig
}

func NewServer(c *config.AppConfig) *server {
	return &server{config: c}
}

func (s *server) Run() error {
	lis, err := net.Listen("tcp", s.config.GrpcPort())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	NotesGrpcApi.RegisterUserServiceServer(srv, &grpcSrv{config: s.config})
	log.Printf("server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (g *grpcSrv) GetUser(ctx context.Context, request *NotesGrpcApi.GetUserRequest) (*NotesGrpcApi.GetUserResponse, error) {

	requestUserId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, err
	}
	usr, err := users.GetUserById(adapters.NewUsersPostgres(g.config), requestUserId)

	if usr == nil || err != nil {
		return &NotesGrpcApi.GetUserResponse{User: nil}, err
	}

	return &NotesGrpcApi.GetUserResponse{User: &NotesGrpcApi.User{UserId: request.GetUserId(), Name: usr.Name}}, nil
}

func (g *grpcSrv) AddUser(ctx context.Context, request *NotesGrpcApi.AddUserRequest) (*NotesGrpcApi.AddUserResponse, error) {

	usr, err := users.AddUser(adapters.NewUsersPostgres(g.config), request.Name)
	if usr == nil || err != nil {
		return &NotesGrpcApi.AddUserResponse{User: nil}, err
	}

	return &NotesGrpcApi.AddUserResponse{User: &NotesGrpcApi.User{UserId: usr.UserId.String(), Name: usr.Name}}, nil
}
