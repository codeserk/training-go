package http

import (
	"api/cat"
	"encoding/json"
	"net/http"
)

type request struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}// @name CreateCatRequest

// Creates a cat
// @Summary Creates a cat
// @Description creates a cat
// @ID create-cat
// @Accept  json
// @Produce  json
// @Param body body request request "Cat parameters"
// @Success 200 {object} entity.Cat
// @Router /v1/cat [post]
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

		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(cat)
	})
}
