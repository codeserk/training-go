package http

import (
	"api/cat"
	"encoding/json"
	"net/http"
)

func findAllCats(service cat.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cats, err := service.FindAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Problem found while trying to get cats"))
		}

		json.NewEncoder(w).Encode(cats)
	})
}
