package http

import (
	"api/cat"
	"encoding/json"
	"net/http"
)

// Finds all the cats
// @Summary Finds all the cats
// @Description finds all the cats
// @ID find-all-cats
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Cat
// @Header 200 {string} Token "qwerty"
// @Router /v1/cat [get]
func findAllCats(service cat.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cats, err := service.FindAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Problem found while trying to get cats"))
		}

		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(cats)
	})
}
