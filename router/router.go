package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/migurd/waterwatch_back/controllers"
	"github.com/migurd/waterwatch_back/helpers"
)

func Routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/client", helpers.MakeHTTPHandleFunc(controllers.GetAllClients))
	return router
}
