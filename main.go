package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	httpSwagger "github.com/swaggo/http-swagger"

	cat "api/cat"
	catHttp "api/cat/http"
	"api/config"

	"github.com/gorilla/mux"

	_ "api/docs"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Print("Request at home")
	json.NewEncoder(w).Encode("hello")
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.

// @contact.name API Support
// @license.name Apache 2.0

// @host localhost:8050
// @BasePath /
func main() {
	log.SetPrefix("training> ")

	router := mux.NewRouter()

	conf, err := config.Parse()
	if err != nil {
		log.Fatalf("There was an error found while trying to load the configuration: \n%s \n", err.Error())
	}

	cats := cat.NewLocalService()
	catHttp.MakeCatHandlers(router, cats)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	http.Handle("/", router)

	log.Printf("Server started at http://localhost:%v/swagger/", conf.Api.Port)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(conf.Api.Port), nil))
}
