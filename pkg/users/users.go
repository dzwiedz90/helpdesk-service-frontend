package users

import (
	"context"

	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
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

func (c *UsersClient) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// prepare stuff for client
	// do stuff for client
	// return response
	addr := c.UsersGRPCAddress + ":" + c.UsersGRPCPort

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pb.NewUsersServiceClient(conn)

	res, err := client.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
