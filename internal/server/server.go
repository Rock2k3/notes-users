package server

import (
	"context"
	notesgrpcapi "github.com/Rock2k3/notes-grpc-api/v2/generated-sources"
	"github.com/Rock2k3/notes-users/config"
	"github.com/Rock2k3/notes-users/internal/adapters"
	"github.com/Rock2k3/notes-users/internal/domain/users"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	config *config.AppConfig
	//httpSrv *noteshttpserver.HttpServer
	grpcSrv *grpcServer
}

type grpcServer struct {
	notesgrpcapi.UnimplementedUserServiceServer
	config *config.AppConfig
}

func NewServer(c *config.AppConfig) *server {
	return &server{
		config: c,
		//httpSrv: noteshttpserver.NewHttpServer(),
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

	//go func() {
	//	err := s.httpSrv.Start(s.config.HttpAddress())
	//	if err != nil {
	//		return
	//	}
	//}()

	log.Printf("server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (g *grpcServer) GetUserByUUID(ctx context.Context, request *notesgrpcapi.GetUserByUUIDRequest) (*notesgrpcapi.GetUserResponse, error) {

	userUUID, err := uuid.Parse(request.GetUserUUID())
	if err != nil {
		return nil, err
	}
	usr, err := users.GetUserByUUID(adapters.NewUsersPostgres(g.config), userUUID)

	if usr == nil || err != nil {
		return &notesgrpcapi.GetUserResponse{User: nil}, err
	}

	return &notesgrpcapi.GetUserResponse{
		User: &notesgrpcapi.User{
			UserUUID: request.GetUserUUID(), Name: usr.Name,
		},
	}, nil
}

func (g *grpcServer) GetUserByName(ctx context.Context, request *notesgrpcapi.GetUserByNameRequest) (*notesgrpcapi.GetUserResponse, error) {

	usr, err := users.GetUserByName(adapters.NewUsersPostgres(g.config), request.GetName())

	if usr == nil || err != nil {
		return &notesgrpcapi.GetUserResponse{User: nil}, err
	}

	return &notesgrpcapi.GetUserResponse{
		User: &notesgrpcapi.User{
			UserUUID: usr.UserUUID.String(), Name: usr.Name,
		},
	}, nil
}

func (g *grpcServer) AddUser(ctx context.Context, request *notesgrpcapi.AddUserRequest) (*notesgrpcapi.AddUserResponse, error) {

	usr, err := users.AddUser(adapters.NewUsersPostgres(g.config), request.Name)
	if usr == nil || err != nil {
		return &notesgrpcapi.AddUserResponse{User: nil}, err
	}

	return &notesgrpcapi.AddUserResponse{
		User: &notesgrpcapi.User{
			UserUUID: usr.UserUUID.String(), Name: usr.Name,
		},
	}, nil
}
