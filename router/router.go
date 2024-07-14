package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/migurd/waterwatch_back/controllers"
	"github.com/migurd/waterwatch_back/helpers"
)

func Routes() http.Handler {
	router := mux.NewRouter()

	// test
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi, world!")
	})

	// client sign up process
	router.HandleFunc("/check_client_email", helpers.MakeHTTPHandleFunc(controllers.CheckClientEmail)).Methods("POST")
	router.HandleFunc("/create_client", helpers.MakeHTTPHandleFunc(controllers.CreateClient)).Methods("POST")

	// only a specific group of people or devices sholud be able to do this
	router.HandleFunc("/create_employee", helpers.MakeHTTPHandleFunc(controllers.CreateEmployee)).Methods("POST")
	router.HandleFunc("/complete_appointment", helpers.MakeHTTPHandleFunc(controllers.CompleteAppointment)).Methods("POST")

	// all public
	router.HandleFunc("/create_client_appointment", helpers.MakeHTTPHandleFunc(controllers.CreateClientAppointment)).Methods("POST")
	// router.HandleFunc("/create_client_appointment_with_appointment", helpers.MakeHTTPHandleFunc(controllers.CreateClientAppointment)).Methods("POST")
	return router
}
