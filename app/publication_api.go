package app

import (
	"github.com/agustin-sarasua/rs-model"
)

type SearchResutlDTO struct {
	items  *[]m.Publication
	count  int64
	offset int64
}
