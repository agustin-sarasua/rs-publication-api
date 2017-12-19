package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	m "github.com/agustin-sarasua/rs-model"
)

func CreatePublicationEndpoint(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetPublicationEndpoint(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SearchPublicationEndpoint(w http.ResponseWriter, req *http.Request) {
	t := req.FormValue("t")
	if strings.ToLower(t) == "nearby" {
		if ps, err := SearchNearByPublication(); err != nil {
			log.Printf("Error creating property")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(m.BuildErrorResponse([]error{err}))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(SearchResutlDTO{items: ps})
		}
	}
}
