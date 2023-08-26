package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
	"github.com/dzwiedz90/helpdesk-service-frontend/config"
	"github.com/dzwiedz90/helpdesk-service-frontend/logs"
	"github.com/dzwiedz90/helpdesk-service-frontend/pkg/users"
)

// Repo the Repository used by the handlers
var Repo *Repository

// Repository is the Repository type
type Repository struct {
	Cfg         *config.Config
	UsersClient *users.UsersClient
}

// NewRepo creates a new Repository
func NewRepo(cfg *config.Config, usersClient *users.UsersClient) *Repository {
	return &Repository{
		Cfg:         cfg,
		UsersClient: usersClient,
	}
}

// NewHandlers sets the Repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to read request's body: %v", err))
		http.Error(w, "Failed to read request's body", http.StatusInternalServerError)
		return
	}

	user := &pb.User{}

	err = json.Unmarshal(body, user)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to protobuf decode data into Go structure: %v", err))
		http.Error(w, "Failed to protobuf decode data into Go structure", http.StatusBadRequest)
		return
	}

	req := pb.CreateUserRequest{
		User: user,
	}

	resp, err := m.UsersClient.CreateUser(r.Context(), &req)
	if err != nil {
		jsonResp := createUserResponse{
			Code:    500,
			Message: err.Error(),
		}

		out, _ := json.MarshalIndent(jsonResp, "", "    ")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(out)

		logs.ErrorLogger(fmt.Sprintf("Internal Server Error: %v", err))
		return
	}

	jsonResp := createUserResponse{
		Code:               200,
		Message:            "User created",
		CreateUserResponse: *resp,
	}

	out, _ := json.MarshalIndent(jsonResp, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(out)
	return
}

type UserID struct {
	ID int `json:"id"`
}

func (m *Repository) GetUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to read request's body: %v", err))
		http.Error(w, "Failed to read request's body", http.StatusInternalServerError)
		return
	}

	var userID UserID

	err = json.Unmarshal(body, &userID)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to protobuf decode data into Go structure: %v", err))
		http.Error(w, "Failed to protobuf decode data into Go structure", http.StatusBadRequest)
		return
	}

	req := pb.GetUserRequest{
		Id: int64(userID.ID),
	}

	resp, err := m.UsersClient.GetUser(r.Context(), &req)
	if err != nil {
		jsonResp := getUserResponse{
			Code:    500,
			Message: err.Error(),
		}

		out, _ := json.MarshalIndent(jsonResp, "", "    ")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(out)

		logs.ErrorLogger(fmt.Sprintf("Internal Server Error: %v", err))
		return
	}

	jsonResp := getUserResponse{
		Code:            200,
		GetUserResponse: *resp,
	}

	out, _ := json.MarshalIndent(jsonResp, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
	return
}

func (m *Repository) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	req := pb.GetAllUsersRequest{}

	resp, err := m.UsersClient.GeAlltUsers(r.Context(), &req)
	if err != nil {
		jsonResp := getAllUsersResponse{
			Code:    500,
			Message: err.Error(),
		}

		out, _ := json.MarshalIndent(jsonResp, "", "    ")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(out)

		logs.ErrorLogger(fmt.Sprintf("Internal Server Error: %v", err))
		return
	}

	jsonResp := getAllUsersResponse{
		Code:                200,
		GetAllUsersResponse: *resp,
	}

	out, _ := json.MarshalIndent(jsonResp, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
	return
}

func (m *Repository) UpdateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to read request's body: %v", err))
		http.Error(w, "Failed to read request's body", http.StatusInternalServerError)
		return
	}

	user := &pb.User{}

	err = json.Unmarshal(body, user)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to protobuf decode data into Go structure: %v", err))
		http.Error(w, "Failed to protobuf decode data into Go structure", http.StatusBadRequest)
		return
	}

	req := pb.UpdateUserRequest{
		User: user,
	}

	resp, err := m.UsersClient.UpdateUser(r.Context(), &req)
	if err != nil {
		jsonResp := updateUserResponse{
			Code:    500,
			Message: err.Error(),
		}

		out, _ := json.MarshalIndent(jsonResp, "", "    ")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(out)

		logs.ErrorLogger(fmt.Sprintf("Internal Server Error: %v", err))
		return
	}

	jsonResp := updateUserResponse{
		Code:                200,
		Message:             "User updated",
		UpdateUsersResponse: *resp,
	}

	out, _ := json.MarshalIndent(jsonResp, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(out)
	return
}

func (m *Repository) DeleteUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to read request's body: %v", err))
		http.Error(w, "Failed to read request's body", http.StatusInternalServerError)
		return
	}

	var userID UserID

	err = json.Unmarshal(body, &userID)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to protobuf decode data into Go structure: %v", err))
		http.Error(w, "Failed to protobuf decode data into Go structure", http.StatusBadRequest)
		return
	}

	req := pb.DeleteUserRequest{
		Id: int64(userID.ID),
	}

	resp, err := m.UsersClient.DeleteUser(r.Context(), &req)
	if err != nil {
		jsonResp := deleteUserResponse{
			Code:    500,
			Message: err.Error(),
		}

		out, _ := json.MarshalIndent(jsonResp, "", "    ")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(out)

		logs.ErrorLogger(fmt.Sprintf("Internal Server Error: %v", err))
		return
	}

	jsonResp := deleteUserResponse{
		Code:                200,
		Message:             "User deleted",
		DeleteUsersResponse: *resp,
	}

	out, _ := json.MarshalIndent(jsonResp, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
	return
}
