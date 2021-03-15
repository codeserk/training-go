package http

import (
	"api/cat"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func findCat(service cat.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		cat, err := service.Find(id)
		if cat == nil {
			json.NewEncoder(w).Encode("Cat not found!")
			return
		}
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(cat)
	})
}
