package server

import (
	"context"
	"github.com/Rock2k3/notes-core/pkg/noteshttpserver"
	notesgrpcapi "github.com/Rock2k3/notes-grpc-api/generated-sources"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	"notes-users/config"
	"notes-users/internal/adapters"
	"notes-users/internal/domain/users"
)

type server struct {
	config  *config.AppConfig
	httpSrv *noteshttpserver.HttpServer
	grpcSrv *grpcServer
}

type grpcServer struct {
	notesgrpcapi.UnimplementedUserServiceServer
	config *config.AppConfig
}

func NewServer(c *config.AppConfig) *server {
	return &server{
		config:  c,
		httpSrv: noteshttpserver.NewHttpServer(),
		grpcSrv: &grpcServer{config: c},
	}
}

func (s *server) Run() {
	lis, err := net.Listen("tcp", s.config.GrpcPort())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	notesgrpcapi.RegisterUserServiceServer(srv, s.grpcSrv)

	go func() {
		err := s.httpSrv.Start(s.config.HttpAddress())
		if err != nil {
			return
		}
	}()

	log.Printf("server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (g *grpcServer) GetUser(ctx context.Context, request *notesgrpcapi.GetUserRequest) (*notesgrpcapi.GetUserResponse, error) {

	requestUserId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, err
	}
	usr, err := users.GetUserById(adapters.NewUsersPostgres(g.config), requestUserId)

	if usr == nil || err != nil {
		return &notesgrpcapi.GetUserResponse{User: nil}, err
	}

	return &notesgrpcapi.GetUserResponse{User: &notesgrpcapi.User{UserId: request.GetUserId(), Name: usr.Name}}, nil
}

func (g *grpcServer) AddUser(ctx context.Context, request *notesgrpcapi.AddUserRequest) (*notesgrpcapi.AddUserResponse, error) {

	usr, err := users.AddUser(adapters.NewUsersPostgres(g.config), request.Name)
	if usr == nil || err != nil {
		return &notesgrpcapi.AddUserResponse{User: nil}, err
	}

	return &notesgrpcapi.AddUserResponse{User: &notesgrpcapi.User{UserId: usr.UserId.String(), Name: usr.Name}}, nil
}
