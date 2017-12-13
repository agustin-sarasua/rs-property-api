package app

import (
	"errors"
	"fmt"
	"log"

	m "github.com/agustin-sarasua/rs-model"
)

var PropertyTypes = map[string]struct{}{
	"APARTAMENTO":          {},
	"CASA":                 {},
	"PENTHOUSE":            {},
	"CAMPO":                {},
	"LOCAL_COMERCIAL":      {},
	"PROPIEDAD_HORIZONTAL": {},
	"TERRENO":              {},
	"GALPON":               {}}

var Orientation = map[string]struct{}{
	"FRENTE":       {},
	"CONTRAFRENTE": {}}

func validateCondition(fn func(map[string]struct{}, string) bool, m map[string]struct{}, v string, msg string, errs []error) []error {
	if ok := fn(m, v); !ok {
		errs = append(errs, errors.New(msg))
	}
	return errs
}

func validateRangeCondition(fn func(int, int, int) bool, min int, max int, val int, msg string, errs []error) []error {
	if ok := fn(min, max, val); !ok {
		errs = append(errs, errors.New(msg))
	}
	return errs
}

func validateNonEmpty(fn func() bool, msg string, errs []error) []error {
	if ok := fn(); !ok {
		errs = append(errs, errors.New(msg))
	}
	return errs
}

func ValidateProperty(p *m.Property) []error {
	var errs []error
	isValueInMap := func(m map[string]struct{}, val string) bool {
		_, ok := m[val]
		return ok
	}
	isValidRange := func(min int, max int, val int) bool {
		return !(val < min || val > max)
	}

	errs = validateCondition(isValueInMap, PropertyTypes, p.Type, "Type is incorrect", errs)
	errs = validateCondition(isValueInMap, Orientation, p.Orientation, "Orientation is incorrect", errs)
	errs = validateRangeCondition(isValidRange, 0, 10, p.State, fmt.Sprintf("State should be between %v and %v", 0, 10), errs)
	errs = validateNonEmpty(func() bool { return p.Address != nil }, "Address can not be empty", errs)
	return errs
}

func CreateProperty(p *m.Property) uint64 {
	log.Printf("Creating new Property: %+v\n", p)
	Db.Create(p)
	log.Printf("Property ID: %+v\n", p.ID)
	return p.ID
}

func UpdateProperty(p *m.Property) uint64 {
	log.Printf("Updating Property: %+v\n", p.ID)
	Db.Save(p)
	return p.ID
}

func LoadProperty(pid uint64) *m.Property {
	log.Printf("Loading Property: %+v\n", pid)
	var p m.Property
	Db.First(&p, pid)
	log.Printf("Loading Property Address: %+v\n", p.AddressID)
	var a m.Address
	Db.First(&a, p.AddressID)
	p.Address = &a
	return &p
}

func SavePropertyState(s *m.PropertyState) uint64 {
	log.Printf("Creating new PropertyState: %+v\n", s)
	var p m.Property
	Db.Find(&p, s.PropertyID)
	if &p == nil {
		panic("E")
	}
	Db.Create(s)
	log.Printf("PropertyState ID: %+v\n", s.ID)
	return s.ID
}
