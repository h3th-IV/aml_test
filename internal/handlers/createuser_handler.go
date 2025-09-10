package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/h3th-IV/aml_test/internal/api"
	"github.com/h3th-IV/aml_test/internal/services"
	"go.uber.org/zap"
)

var _ http.Handler = &CreateUserHandler{}

type CreateUserHandler struct {
	logger      *zap.Logger
	userService services.UserService
}

func NewCreateUserHandler(logger *zap.Logger, userService services.UserService) *CreateUserHandler {
	return &CreateUserHandler{
		logger:      logger,
		userService: userService,
	}
}

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (handler *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handler.logger.Sugar().Error("epxecting method: %v, got method: %v", http.MethodPost, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Method not allowed",
			Data:    nil,
		})
		return
	}
	apiClient := api.NewAPIClient("https://randomuser.me/api/")

	user, err := apiClient.FetchUser(r.Context())
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
