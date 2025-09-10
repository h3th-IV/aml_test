package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/h3th-IV/aml_test/internal/api"
	"github.com/h3th-IV/aml_test/internal/config"
	"github.com/h3th-IV/aml_test/internal/handlers"
	"github.com/h3th-IV/aml_test/internal/services"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()

	db, err := config.ConnectDatabase()
	if err != nil {
		logger.Sugar().Error("err creating database connection: ", err)
		return
	}
	defer db.Close()

	mux_router := mux.NewRouter()
	apiClient := api.NewAPIClient("https://randomuser.me/api/")
	mux_router.Handle("/create-user", handlers.NewCreateUserHandler(logger, *services.NewService(db), apiClient)).Methods(http.MethodPost)
	mux_router.Handle("/get-user/{id}", handlers.NewGetUserHandler(logger, *services.NewService(db))).Methods(http.MethodGet)

	server := &http.Server{
		Addr:              ":5000",
		Handler:           mux_router,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	logger.Sugar().Info("listening and serving @ :5000")
	if err := server.ListenAndServe(); err != nil {
		logger.Sugar().Error("error starting server :(")
	}
}
