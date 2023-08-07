package service

import (
	"encoding/json"
	"io"
	"net/http"

	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
	"github.com/dzwiedz90/helpdesk-service-frontend/config"
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
		http.Error(w, "Failed to read request's body", http.StatusInternalServerError)
		return
	}

	user := &pb.User{}

	err = json.Unmarshal(body, user)
	if err != nil {
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
