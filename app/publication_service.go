package app

import (
	"log"

	m "github.com/agustin-sarasua/rs-model"
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
