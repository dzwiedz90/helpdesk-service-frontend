package service

import (
	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
)

type createUserResponse struct {
	Code               int                   `json:"code"`
	Message            string                `json:"message"`
	CreateUserResponse pb.CreateUserResponse `json:"createUserResponse"`
}