package main

import (
	"net/http"

	"github.com/gorilla/mux"
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
	mux_router.Handle("/create-user", handlers.NewCreateUserHandler(logger, *services.NewService(db))).Methods(http.MethodPost)
	mux_router.Handle("/get-user/{id}", handlers.NewGetUserHandler(logger, *services.NewService(db))).Methods(http.MethodGet)
	logger.Sugar().Info("listening and serving @ :5000")
	if err := http.ListenAndServe(":5000", mux_router); err != nil {
		logger.Sugar().Error("error starting server :(")
	}
}
