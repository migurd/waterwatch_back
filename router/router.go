package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/migurd/waterwatch_back/helpers"
)

type Message struct {
	Message string `json:"message"`
}

func test(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, http.StatusOK, &Message{Message: "hi"})
}

func Routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", test)
	return router
}