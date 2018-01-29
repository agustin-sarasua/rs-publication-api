package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	c "github.com/agustin-sarasua/rs-common"
	m "github.com/agustin-sarasua/rs-model"
)

var (
	StorageBucket     *storage.BucketHandle
	StorageBucketName string
)

func LoadUserPublicationEndpoint(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreatePublicationEndpoint(w http.ResponseWriter, req *http.Request) {
	var msg m.Publication
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		c.ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}
	msg.CreatedAt = time.Now()
	if id, errs := CreatePublication(&msg); len(errs) > 0 {
		log.Printf("Error creating publication")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(m.BuildErrorResponse(errs))
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "{id: %q}", id)
	}
	w.Header().Set("Content-Type", "application/json")
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

func UploadHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(rw, "", http.StatusMethodNotAllowed)
		return
	}

	f, fh, err := r.FormFile("file")

	if err == http.ErrMissingFile {
		return
	}
	if err != nil {
		return
	}

	nc := PushImageToCloudStorage(f, fh)
	n := <-nc

	const publicURL = "https://storage.googleapis.com/%s/%s"

	fmt.Fprintf(rw, fmt.Sprintf(publicURL, StorageBucketName, n))
}
