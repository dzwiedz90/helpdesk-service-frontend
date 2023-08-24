package users

import (
	"context"
	"fmt"

	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
	"github.com/dzwiedz90/helpdesk-service-frontend/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UsersClient struct {
	UsersGRPCPort    string
	UsersGRPCAddress string
}

func NewClient(usersGRPCPort string, usersGRPCAddress string) *UsersClient {
	return &UsersClient{
		UsersGRPCPort:    usersGRPCPort,
		UsersGRPCAddress: usersGRPCAddress,
	}
}

func (c *UsersClient) CreateConn() (*grpc.ClientConn, error) {
	addr := c.UsersGRPCAddress + ":" + c.UsersGRPCPort

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to create grpc connection to service-users: %v", err))
		return nil, err
	}

	return conn, nil
}

func (c *UsersClient) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	conn, err := c.CreateConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewUsersServiceClient(conn)

	res, err := client.CreateUser(ctx, req)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to create user: %v", err))
		return nil, err
	}

	return res, nil
}

func (c *UsersClient) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	conn, err := c.CreateConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewUsersServiceClient(conn)

	res, err := client.GetUser(ctx, req)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to get user: %v", err))
		return nil, err
	}

	return res, nil
}

func (c *UsersClient) GeAlltUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	conn, err := c.CreateConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewUsersServiceClient(conn)

	res, err := client.GetAllUsers(ctx, req)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to get all user: %v", err))
		return nil, err
	}

	return res, nil
}

func (c *UsersClient) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, nil
}

func (c *UsersClient) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}
