package main

import (
	"encoding/json"
	"log"
	"net/http"

	cat "api/cat"
	catHttp "api/cat/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Print("Request at home")
	json.NewEncoder(w).Encode("hello")
}

func main() {
	log.SetPrefix("training> ")

	router := mux.NewRouter()

	cats := cat.NewLocalService()

	catHttp.MakeCatHandlers(router, cats)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
