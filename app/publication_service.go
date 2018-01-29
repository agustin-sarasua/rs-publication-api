package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"path"

	"cloud.google.com/go/storage"
	m "github.com/agustin-sarasua/rs-model"
	uuid "github.com/satori/go.uuid"
)

func CreatePublication(p *m.Publication) (uint64, []error) {
	log.Printf("Creating new Publication: %+v\n", p)
	if errs := validatePublication(p); len(errs) > 0 {
		return 0, errs
	}
	Db.Create(p)
	log.Printf("Publication ID: %+v\n", p.ID)
	return p.ID, nil
}

func SearchNearByPublication() (*[]m.Publication, error) {

	return nil, nil
}

func PushImageToCloudStorage(f multipart.File, fh *multipart.FileHeader) <-chan string {
	out := make(chan string)
	go func() {
		if StorageBucket == nil {
			return
		}
		b1 := make([]byte, 5)
		n1, err := f.Read(b1)
		fmt.Printf("%d bytes: %s\n", n1, string(b1))

		if err != nil {
			return
		}
		// random filename, retaining existing extension.
		name := uuid.NewV4().String() + path.Ext(fh.Filename)

		ctx := context.Background()
		w := StorageBucket.Object(name).NewWriter(ctx)
		w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
		w.ContentType = fh.Header.Get("Content-Type")

		// Entries are immutable, be aggressive about caching (1 day).
		w.CacheControl = "public, max-age=86400"

		if _, err := io.Copy(w, f); err != nil {
			return
		}
		if err := w.Close(); err != nil {
			return
		}
		out <- name
	}()
	return out
}
