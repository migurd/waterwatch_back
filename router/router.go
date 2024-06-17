package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/migurd/waterwatch_back/controllers"
	"github.com/migurd/waterwatch_back/helpers"
)

func Routes() http.Handler {
	router := mux.NewRouter()

	// Create
	router.HandleFunc("/client", helpers.MakeHTTPHandleFunc(controllers.CreateClient)).Methods("POST")

	// Read
	router.HandleFunc("/client", helpers.MakeHTTPHandleFunc(controllers.GetAllClients)).Methods("GET")
	return router
}
