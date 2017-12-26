package app

import (
	"log"

	m "github.com/agustin-sarasua/rs-model"
)

func CreateProperty(p *m.Property) (uint64, []error) {
	log.Printf("Creating new Property: %+v\n", p)
	if errs := validateProperty(p); len(errs) > 0 {
		return 0, errs
	}
	Db.Create(p)
	log.Printf("Property ID: %+v\n", p.ID)
	return p.ID, nil
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

func ListProperties(offset, limit int) ([]m.Property, error) {
	log.Println("Loading all Properties")
	var properties []m.Property
	Db.Offset(offset).Limit(limit).Find(&properties)
	return properties, nil
}
