package service

import (
	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
)

type createUserResponse struct {
	Code               int                   `json:"code"`
	Message            string                `json:"message"`
	CreateUserResponse pb.CreateUserResponse `json:"createUserResponse"`
}

type getUserResponse struct {
	Code            int                `json:"code"`
	Message         string             `json:"message"`
	GetUserResponse pb.GetUserResponse `json:"getUserResponse"`
}

type getAllUsersResponse struct {
	Code                int                    `json:"code"`
	Message             string                 `json:"message"`
	GetAllUsersResponse pb.GetAllUsersResponse `json:"getAllUsersResponse"`
}

type updateUserResponse struct {
}

type deleteUserResponse struct {
}
