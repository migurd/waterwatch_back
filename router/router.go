package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/migurd/waterwatch_back/controllers"
	"github.com/migurd/waterwatch_back/helpers"
)

func Routes() http.Handler {
	router := mux.NewRouter()

	// only a specific group of people or devices sholud be able to do this
	router.HandleFunc("/create_employee", helpers.MakeHTTPHandleFunc(controllers.CreateEmployee)).Methods("POST")
	router.HandleFunc("/complete_appointment", helpers.MakeHTTPHandleFunc(controllers.CompleteAppointment)).Methods("POST")

	// all public
	router.HandleFunc("/create_client_appointment", helpers.MakeHTTPHandleFunc(controllers.CreateClientAppointment)).Methods("POST")
	// router.HandleFunc("/create_client_appointment_with_appointment", helpers.MakeHTTPHandleFunc(controllers.CreateClientAppointment)).Methods("POST")
	return router
}
