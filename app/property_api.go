package app

import (
	"github.com/agustin-sarasua/rs-model"
)

type SearchResutlDTO struct {
	Items  []m.Property
	Count  int64
	Offset int64
}
