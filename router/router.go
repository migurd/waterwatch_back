package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/migurd/waterwatch_back/controllers"
	"github.com/migurd/waterwatch_back/helpers"
)

func Routes(controllers *controllers.Controllers) *mux.Router {
	router := mux.NewRouter()

	// test
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi, world!")
	})

	// ================
	// client sign up & login process
	router.HandleFunc("/check-client-email", helpers.MakeHTTPHandleFunc(controllers.CheckClientEmail)).Methods("POST")
	router.HandleFunc("/client-register", helpers.MakeHTTPHandleFunc(controllers.CreateClient)).Methods("POST")
	router.HandleFunc("/client-login", helpers.MakeHTTPHandleFunc(controllers.ClientLogin)).Methods("POST")

	// TODO
	// already logged in user
	// --> create-appointment
	// --> update-appoinment
	// --> cancel-appoinment
	// --> create-address
	// --> update-address
	// --> delete-address
	// --> get appointments of user

	// ================
	// employee register & login process
	// --> register (probably no)
	// --> login

	// already logged in client
	// --> get appointments assigned to them
	// --> get appointments that haven't been assigned to an employee and fit to their role
	// --> update cancel assigned appointment
	// --> based on employee type
	//				--> post complete installation
	//				--> post complete maintenance

	// ================
	// IoT Device
	// --> post saa_record using serial_key
	// --> gett all saa_record using serial_key
	router.HandleFunc("/create-saa-record", helpers.MakeHTTPHandleFunc(controllers.CreateSaaRecord)).Methods("POST")
	router.HandleFunc("/get-saa-records", helpers.MakeHTTPHandleFunc(controllers.GetSaaRecords)).Methods("GET")

	// ================
	// mobile home based on the user
	// --> get view home (kind of all client tables)
	// --> get view saa all
	// --> get view saa specific
	// --> update saa specific name and description
	// --> get view contact
	// --> get view saa_maintenance all

	return router
}
