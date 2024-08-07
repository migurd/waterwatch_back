package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/migurd/waterwatch_back/controllers"
	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/middleware"
)

func Routes(controllers *controllers.Controllers) *mux.Router {
	router := mux.NewRouter()

	// cross-origin resource sharing
	router.Use(middleware.CORS)

	// protected routers
	clientRoutes := router.PathPrefix("/client").Subrouter()
	clientRoutes.Use(middleware.Authentication)
	clientRoutes.Use(middleware.ClientOnly)

	employeeRoutes := router.PathPrefix("/employee").Subrouter()
	employeeRoutes.Use(middleware.Authentication)
	employeeRoutes.Use(middleware.EmployeeOnly)

	// test
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi, world!")
	})

	// ================
	// client sign up & login process
	router.HandleFunc("/client/check-email", helpers.MakeHTTPHandleFunc(controllers.CheckClientEmail)).Methods("POST")
	router.HandleFunc("/client/register", helpers.MakeHTTPHandleFunc(controllers.CreateClient)).Methods("POST")
	router.HandleFunc("/client/login", helpers.MakeHTTPHandleFunc(controllers.ClientLogin)).Methods("POST")

	// PROTECTED
	// already logged in user
	// --> create-installation-appointment
	// --> read-installation-appointment
	// --> update-installation-appoinment
	// --> delete-installation-appoinment
	// == TAKE INTO CONSIDERATION THAT THE TYPE OF APPOINTMENT MUST BE SENT
	clientRoutes.HandleFunc("/create-installation-appointment", helpers.MakeHTTPHandleFunc(controllers.CreateAppointment(1))).Methods("POST")
	clientRoutes.HandleFunc("/get-pending-installation-appointment", helpers.MakeHTTPHandleFunc(controllers.GetPendingAppointment(1))).Methods("GET")
	clientRoutes.HandleFunc("/update-installation-appointment", helpers.MakeHTTPHandleFunc(controllers.UpdateAppointment(1))).Methods("PATCH")
	clientRoutes.HandleFunc("/delete-installation-appointment", helpers.MakeHTTPHandleFunc(controllers.DeleteAppointment(1))).Methods("DELETE")

	// PROTECTED
	// --> create-client-address
	// --> read-client-address
	// --> read-client-addresses
	// --> update-client-address
	// --> delete-client-address
	clientRoutes.HandleFunc("/create-address", helpers.MakeHTTPHandleFunc(controllers.CreateClientAddress)).Methods("POST")
	clientRoutes.HandleFunc("/get-address", helpers.MakeHTTPHandleFunc(controllers.GetClientAddress)).Methods("GET")
	clientRoutes.HandleFunc("/get-all-addresses", helpers.MakeHTTPHandleFunc(controllers.GetAllClientAddresses)).Methods("GET")
	clientRoutes.HandleFunc("/update-address", helpers.MakeHTTPHandleFunc(controllers.UpdateClientAddress)).Methods("PATCH")
	clientRoutes.HandleFunc("/delete-address", helpers.MakeHTTPHandleFunc(controllers.DeleteClientAddress)).Methods("DELETE")

	// ================
	// employee register & login process
	// --> register (there's no interface for that, the only person that can post in this router is the sysadmin)
	// --> login
	router.HandleFunc("/employee/register", helpers.MakeHTTPHandleFunc(controllers.CreateEmployee)).Methods("POST")
	router.HandleFunc("/employee/login", helpers.MakeHTTPHandleFunc(controllers.EmployeeLogin)).Methods("POST")

	// PROTECTED
	// already logged in employee
	// --> get appointments that haven't been assigned to an employee and fit to their role (installer, mantainer or both)
	// --> get appointments assigned to them
	// --> update cancel assigned appointment
	// --> based on employee type
	//				--> post complete installation
	//				--> post complete maintenance
	employeeRoutes.HandleFunc("/get-all-appointments-not-assigned", helpers.MakeHTTPHandleFunc(controllers.GetAllUnassignedAppointments)).Methods("GET")
	employeeRoutes.HandleFunc("/get-all-appointments-assigned", helpers.MakeHTTPHandleFunc(controllers.GetAllAppointmentsAssigned)).Methods("GET")
	employeeRoutes.HandleFunc("/accept-appointment", helpers.MakeHTTPHandleFunc(controllers.AcceptAppointment)).Methods("PATCH")
	employeeRoutes.HandleFunc("/cancel-appointment", helpers.MakeHTTPHandleFunc(controllers.CancelAppointmentEmployee)).Methods("PATCH")
	employeeRoutes.HandleFunc("/complete-installation", helpers.MakeHTTPHandleFunc(controllers.CompleteAppointment(1))).Methods("POST")
	employeeRoutes.HandleFunc("/complete-maintenance", helpers.MakeHTTPHandleFunc(controllers.CompleteAppointment(2))).Methods("POST")

	// ================
	// IoT Device
	// --> post saa_record using serial_key
	// --> gett all saa_record using serial_key
	router.HandleFunc("/create-saa-record", helpers.MakeHTTPHandleFunc(controllers.CreateSaaRecord)).Methods("POST")
	router.Handle("/get-saa-height", helpers.MakeHTTPHandleFunc(controllers.GetSaaHeight)).Methods("GET")
	clientRoutes.HandleFunc("/get-all-saa-records", helpers.MakeHTTPHandleFunc(controllers.GetSaaRecords)).Methods("GET")

	// =======================================================================
	// PROTECTED
	// mobile home based on the user
	// --> get view home (kind of all client tables)
	// --> get view saa all
	// --> get view saa specific
	// --> patch saa specific name and description
	// --> get view contact
	// --> get view saa_maintenance all
	clientRoutes.HandleFunc("/get-profile", helpers.MakeHTTPHandleFunc(controllers.GetHome)).Methods("GET")
	clientRoutes.HandleFunc("/get-all-active-saa", helpers.MakeHTTPHandleFunc(controllers.GetAllActiveSaaForClient)).Methods("GET")
	// clientRoutes.HandleFunc("/get-active-saa", helpers.MakeHTTPHandleFunc(controllers.GetActiveSaa)).Methods("GET")
	// -->	going to implement later. needs to return smth related to how it good is the water.
	// -->	take into consideration that a new table is needed to compare the water (we need to know how GOOD is GOOD in pH)
	clientRoutes.HandleFunc("/get-saa-description", helpers.MakeHTTPHandleFunc(controllers.GetSaaDescription)).Methods("GET")
	clientRoutes.HandleFunc("/update-saa-description", helpers.MakeHTTPHandleFunc(controllers.UpdateSaaDescription)).Methods("PATCH")
	clientRoutes.HandleFunc("/get-all-contacts", helpers.MakeHTTPHandleFunc(controllers.GetAllContactsInfo)).Methods("GET")
	clientRoutes.HandleFunc("/get-all-installation-appointments-done", helpers.MakeHTTPHandleFunc(controllers.GetAllDoneAppointments(1))).Methods("GET")
	clientRoutes.HandleFunc("/get-all-maintenance-appointments-done", helpers.MakeHTTPHandleFunc(controllers.GetAllDoneAppointments(2))).Methods("GET")

	// PROTECTED
	// --> create maintenance appointment
	// --> read maintenance appointment
	// --> update maintenance appointment
	// --> delete maintenance appointment
	clientRoutes.HandleFunc("/create-maintenance-appointment", helpers.MakeHTTPHandleFunc(controllers.CreateAppointment(2))).Methods("POST")
	clientRoutes.HandleFunc("/get-pending-maintenance-appointment", helpers.MakeHTTPHandleFunc(controllers.GetPendingAppointment(2))).Methods("GET")
	clientRoutes.HandleFunc("/update-maintenance-appointment", helpers.MakeHTTPHandleFunc(controllers.UpdateAppointment(2))).Methods("PATCH")
	clientRoutes.HandleFunc("/delete-maintenance-appointment", helpers.MakeHTTPHandleFunc(controllers.DeleteAppointment(2))).Methods("DELETE")

	return router
}
