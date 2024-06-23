package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/migurd/waterwatch_back/controllers"
	"github.com/migurd/waterwatch_back/helpers"
)

func Routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/create_client_appointment", helpers.MakeHTTPHandleFunc(controllers.CreateClientAppointment)).Methods("POST")
	router.HandleFunc("/create_employee", helpers.MakeHTTPHandleFunc(controllers.CreateEmployee)).Methods("POST")
	return router
}
