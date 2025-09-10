package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/h3th-IV/aml_test/internal/models"
	"github.com/h3th-IV/aml_test/internal/services"
	"go.uber.org/zap"
)

var _ http.Handler = &CreateUserHandler{}

type RandomUserFetcher interface {
	FetchUser(ctx context.Context) (*models.User, error)
}

type CreateUserHandler struct {
	logger      *zap.Logger
	userService services.UserService
	fetcher     RandomUserFetcher
}

func NewCreateUserHandler(logger *zap.Logger, userService services.UserService, fetcher RandomUserFetcher) *CreateUserHandler {
	return &CreateUserHandler{
		logger:      logger,
		userService: userService,
		fetcher:     fetcher,
	}
}

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (handler *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handler.logger.Sugar().Errorf("expecting method: %s, got: %s", http.MethodPost, r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Method not allowed",
			Data:    nil,
		})
		return
	}

	user, err := handler.fetcher.FetchUser(r.Context())
	if err != nil {
		handler.logger.Sugar().Error("err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "An error occured while fetching user from external API.",
			Data:    nil,
		})
		return
	}

	user_service, err := handler.userService.CreateNewUser(r.Context(), user.Name, user.Email, user.Gender, user.Address, user.Dob)
	if err != nil {
		handler.logger.Sugar().Error("err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "An error occured while creating user.",
			Data:    nil,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "User created successfully",
		Data:    user_service,
	})
}
