package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/agustin-sarasua/rs-publication-api/app"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/publication", app.LoadUserPublicationEndpoint).Methods("GET")
	router.HandleFunc("/publication", app.CreatePublicationEndpoint).Methods("POST")
	router.HandleFunc("/publication/{id:[0-9]+}", app.GetPublicationEndpoint).Methods("GET")

	router.HandleFunc("/publication/_search", app.SearchPublicationEndpoint).Methods("GET").Queries("t", "t")

	app.StorageBucketName = "real-estate-project-186513.appspot.com"
	app.StorageBucket, _ = configureStorage(app.StorageBucketName)

	fmt.Println("Hello there")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func configureStorage(bucketID string) (*storage.BucketHandle, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.Bucket(bucketID), nil
}
