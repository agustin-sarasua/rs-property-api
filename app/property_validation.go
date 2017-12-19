package app

import (
	"fmt"
	"log"
	"testing"

	c "github.com/agustin-sarasua/rs-common"
	m "github.com/agustin-sarasua/rs-model"
)

func validateProperty(p *m.Property) []error {
	var errs []error

	errs = c.ValidateExistInMap(m.PropertyTypes, p.Type, "Type is incorrect", errs)
	errs = c.ValidateExistInMap(m.Orientation, p.Orientation, "Orientation is incorrect", errs)
	errs = c.ValidateRangeCondition(0, 10, p.State, fmt.Sprintf("State should be between %v and %v", 0, 10), errs)
	errs = c.ValidateCondition(func() bool { return p.Address != nil }, "Address can not be empty", errs)
	return errs
}

func testValidateProperty(t *testing.T) {
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
