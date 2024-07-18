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
	router.HandleFunc("/client/check-email", helpers.MakeHTTPHandleFunc(controllers.CheckClientEmail)).Methods("POST")
	router.HandleFunc("/client/register", helpers.MakeHTTPHandleFunc(controllers.CreateClient)).Methods("POST")
	router.HandleFunc("/client/login", helpers.MakeHTTPHandleFunc(controllers.ClientLogin)).Methods("POST")

	// TODO
	// already logged in user
	// --> create-installation-appointment
	// --> read-installation-appointment
	// --> update-installation-appoinment
	// --> delete-installation-appoinment
	// == TAKE INTO CONSIDERATION THAT THE TYPE OF APPOINTMENT MUST BE SENT
	router.HandleFunc("/client/create-installation-appointment", helpers.MakeHTTPHandleFunc(controllers.CreateAppointment(1))).Methods("POST")
	router.HandleFunc("/client/get-pending-installation-appointment", helpers.MakeHTTPHandleFunc(controllers.GetPendingAppointment(1))).Methods("GET")
	router.HandleFunc("/client/update-installation-appointment", helpers.MakeHTTPHandleFunc(controllers.UpdateAppointment)).Methods("PATCH")
	router.HandleFunc("/client/delete-installation-appointment", helpers.MakeHTTPHandleFunc(controllers.DeleteAppointment)).Methods("DELETE")

	// --> create-client-address
	// --> read-client-address
	// --> read-client-addresses
	// --> update-client-address
	// --> delete-client-address
	router.HandleFunc("/client/create-address", helpers.MakeHTTPHandleFunc(controllers.CreateClientAddress)).Methods("POST")
	router.HandleFunc("/client/get-address", helpers.MakeHTTPHandleFunc(controllers.GetClientAddress)).Methods("GET")
	router.HandleFunc("/client/get-all-addresses", helpers.MakeHTTPHandleFunc(controllers.GetAllClientAddresses)).Methods("GET")
	router.HandleFunc("/client/update-address", helpers.MakeHTTPHandleFunc(controllers.UpdateClientAddress)).Methods("PATCH")
	router.HandleFunc("/client/delete-address", helpers.MakeHTTPHandleFunc(controllers.DeleteClientAddress)).Methods("DELETE")

	// ================
	// employee register & login process
	// --> register (there's no interface for that, the only person that can post in this router is the sysadmin)
	// --> login
	router.HandleFunc("/employee/register", helpers.MakeHTTPHandleFunc(controllers.CreateEmployee)).Methods("POST")
	router.HandleFunc("/employee/login", helpers.MakeHTTPHandleFunc(controllers.EmployeeLogin)).Methods("POST")

	// already logged in employee
	// --> get appointments that haven't been assigned to an employee and fit to their role (installer, mantainer or both)
	// --> get appointments assigned to them
	// --> update cancel assigned appointment
	// --> based on employee type
	//				--> post complete installation
	//				--> post complete maintenance
	router.HandleFunc("/employee/get-all-appointments-not-assigned", helpers.MakeHTTPHandleFunc(controllers.GetAllUnassignedAppointments)).Methods("GET")
	router.HandleFunc("/employee/get-all-appointments-assigned", helpers.MakeHTTPHandleFunc(controllers.GetAllAppointmentsAssigned)).Methods("GET")
	router.HandleFunc("/employee/accept-appointment", helpers.MakeHTTPHandleFunc(controllers.AcceptAppointment)).Methods("PATCH")
	router.HandleFunc("/employee/cancel-appointment", helpers.MakeHTTPHandleFunc(controllers.CancelAppointmentEmployee)).Methods("PATCH")
	router.HandleFunc("/employee/complete-installation", helpers.MakeHTTPHandleFunc(controllers.CompleteAppointment(1))).Methods("GET")
	router.HandleFunc("/employee/complete-maintenance", helpers.MakeHTTPHandleFunc(controllers.CompleteAppointment(2))).Methods("GET")


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

	// --> create maintenance appointment
	// --> read maintenance appointment
	// --> update maintenance appointment
	// --> delete maintenance appointment
	router.HandleFunc("/employee/create-installation-appointment", helpers.MakeHTTPHandleFunc(controllers.CreateAppointment(2))).Methods("POST")
	router.HandleFunc("/employee/get-pending-installation-appointment", helpers.MakeHTTPHandleFunc(controllers.GetPendingAppointment(2))).Methods("GET")
	router.HandleFunc("/employee/update-installation-appointment", helpers.MakeHTTPHandleFunc(controllers.UpdateAppointment)).Methods("PATCH")
	router.HandleFunc("/employee/delete-installation-appointment", helpers.MakeHTTPHandleFunc(controllers.DeleteAppointment)).Methods("DELETE")

	return router
}
