package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/migurd/waterwatch_back/internal/storage"
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

	router.HandleFunc("/address", makeHTTPHandleFunc(s.handleAddress))
	router.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/phone_number", makeHTTPHandleFunc(s.handlePhoneNumber))
	router.HandleFunc("/microcontroller", makeHTTPHandleFunc(s.handleMicrocontroller))
	router.HandleFunc("/SAA_type", makeHTTPHandleFunc(s.handleSAAType))
	router.HandleFunc("/SAA_record", makeHTTPHandleFunc(s.handleSAARecord))
	router.HandleFunc("/SAA_maintenance", makeHTTPHandleFunc(s.handleSAAMaintenance))
	router.HandleFunc("/SAA", makeHTTPHandleFunc(s.handleSAA))

	log.Println("JSON API server running on:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

// Support fns & types

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
