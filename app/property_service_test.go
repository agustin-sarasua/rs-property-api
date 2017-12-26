package app

import (
	"log"
	"testing"

	"github.com/agustin-sarasua/rs-model"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestValidateProperty(t *testing.T) {
	log.Println("Running TestValidateProperty")
	p := m.Property{Type: "APARTAMENTO", Orientation: "FRENTE", State: 3, Address: &m.Address{}}
	errs := validateProperty(&p)
	if len(errs) > 0 {
		t.Errorf("Error validating property")
	}

	p = m.Property{Type: "FRUTA", Orientation: "FRENTE", State: 3}
	errs = validateProperty(&p)
	log.Printf("Errores: %v", len(errs))
	if len(errs) != 2 {
		t.Errorf("Error validating property")
	}

	p = m.Property{Type: "FRUTA", Orientation: "FRUTA", State: 3, Address: &m.Address{}}
	errs = validateProperty(&p)
	log.Printf("Errores: %v", len(errs))
	if len(errs) != 2 {
		t.Errorf("Error validating property")
	}

	p = m.Property{Type: "FRUTA", Orientation: "FRUTA", State: 23, Address: &m.Address{}}
	errs = validateProperty(&p)
	log.Printf("Errores: %v", len(errs))
	if len(errs) != 3 {
		t.Errorf("Error validating property")
	}
}

func TestListProperties(t *testing.T) {

}
