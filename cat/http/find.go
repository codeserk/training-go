package http

import (
	"api/cat"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Finds a cat by its id
// @Summary Finds a cat by its id
// @Description get cat by id
// @ID get-cat-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Cat ID"
// @Success 200 {object} entity.Cat
// @Router /v1/cat/{id} [get]
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

		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(cat)
	})
}
