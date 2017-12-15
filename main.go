package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agustin-sarasua/rs-publication-api/app"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/publication", app.CreatePublicationEndpoint).Methods("POST")
	router.HandleFunc("/publication/{id:[0-9]+}", app.GetPublicationEndpoint).Methods("GET")

	fmt.Println("Hello there")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
