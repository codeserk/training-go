package http

import (
	"api/cat"

	"github.com/gorilla/mux"
)

func MakeCatHandlers(router *mux.Router, service cat.Service) {
	router.Handle("/v1/cat/{id}", findCat(service)).Methods("GET")
	router.Handle("/v1/cat", findAllCats(service)).Methods("GET")

	router.Handle("/v1/cat", createCat(service)).Methods("POST")
}
