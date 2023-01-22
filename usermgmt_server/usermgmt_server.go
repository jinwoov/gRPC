package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/jinwoov/gRPC/usermgmt"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		user_list: &pb.UserList{},
	}
}

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	user_list *pb.UserList
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("Server Listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	createdUser := &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   user_id,
	}
	s.user_list.Users = append(s.user_list.Users, createdUser)
	return createdUser, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUserParams) (*pb.UserList, error) {
	return s.user_list, nil
}

func main() {
	mgmt := NewUserManagementServer()
	if err := mgmt.Run(); err !=nil {
		log.Fatalf("unable to start the server: %v", err)
	}
}
