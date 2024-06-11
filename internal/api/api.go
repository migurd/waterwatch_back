package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/migurd/waterwatch_back/internal/storage"
	"github.com/migurd/waterwatch_back/helpers"
)

type APIServer struct {
	listenAddr string
	store      storage.Storage
}

func NewAPIServer(listenAddr string, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	// General fnctionalities
	// TODO: Shouldn't have access to most of these if they're not logged in
	router.HandleFunc("/address", helpers.MakeHTTPHandleFunc(s.handleAddress))
	router.HandleFunc("/user", helpers.MakeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/account", helpers.MakeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/phone_number", helpers.MakeHTTPHandleFunc(s.handlePhoneNumber))
	router.HandleFunc("/microcontroller", helpers.MakeHTTPHandleFunc(s.handleMicrocontroller))
	router.HandleFunc("/SAA_type", helpers.MakeHTTPHandleFunc(s.handleSAAType))
	router.HandleFunc("/SAA_record", helpers.MakeHTTPHandleFunc(s.handleSAARecord))
	router.HandleFunc("/SAA_maintenance", helpers.MakeHTTPHandleFunc(s.handleSAAMaintenance))
	router.HandleFunc("/SAA", helpers.MakeHTTPHandleFunc(s.handleSAA))

	log.Println("JSON API server running on:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}
