package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/h3th-IV/aml_test/internal/services"
	"go.uber.org/zap"
)

var _ http.Handler = &GetUserHandler{}

type GetUserHandler struct {
	logger      *zap.Logger
	userService services.UserService
}

func NewGetUserHandler(logger *zap.Logger, userService services.UserService) *GetUserHandler {
	return &GetUserHandler{
		logger:      logger,
		userService: userService,
	}
}

func (handler *GetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handler.logger.Sugar().Error("epxecting method: %v, got method: %v", http.MethodPost, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Method not allowed",
			Data:    nil,
		})
		return
	}

	//fetch id from request
	vars := mux.Vars(r)
	wet_id := vars["id"]
	id, err := strconv.Atoi(wet_id)
	if err != nil {
		handler.logger.Sugar().Error("Error contaning id generation")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "An error occured while fetcing user",
			Data:    nil,
		})
		return
	}

	user, err := handler.userService.GetuserById(r.Context(), id)
	if err != nil {
		handler.logger.Sugar().Error("Err fethcing the usrer from database")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "An error occured while fetcing user",
			Data:    nil,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	})
}
