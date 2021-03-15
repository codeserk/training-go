package http

import (
	"api/cat"
	"encoding/json"
	"net/http"
)

type request struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func createCat(service cat.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input request
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Invalid input"))
		}

		cat, err := service.Create(input.Name, input.Color)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("There was a problem while trying to create your cat."))
		}

		json.NewEncoder(w).Encode(cat)
	})
}
