package app

import (
	"log"
	"os"
	"testing"

	"github.com/agustin-sarasua/rs-model"

	model "github.com/agustin-sarasua/rs-model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestMain(m *testing.M) {
	log.Printf("Running test main")
	// call flag.Parse() here if TestMain uses flags
	db, _ := gorm.Open("sqlite3", "/tmp/gorm.db")
	defer db.Close()
	db.LogMode(true)
	var property model.Property
	var address model.Address
	var propertyState model.PropertyState
	db.DropTableIfExists(&property, &address, &propertyState)
	db.AutoMigrate(&property, &address, &propertyState)

	log.Printf("Running specific test")
	db.Model(&property).Related(&address)
	Db = db

	os.Exit(m.Run())
}

func TestCreateProperty(t *testing.T) {
	log.Println("Running TestCreateProperty")
	var p model.Property
	p = model.Property{Type: "APARTAMENTO", Orientation: "FRENTE", State: 3, Address: &m.Address{}}

	id, errs := CreateProperty(&p)
	if len(errs) > 0 {
		t.Errorf("Property was not saved")
	}

	log.Printf("Loading Property recently saved ID %v\n", p.ID)
	if err := Db.Find(&model.Property{}, id).Error; err != nil {
		t.Errorf("Property was not saved, err= %v", err)
	}
}

func TestValidateProperty(t *testing.T) {
	p := model.Property{Type: "APARTAMENTO", Orientation: "FRENTE", State: 3, Address: &m.Address{}}
	errs := validateProperty(&p)
	if len(errs) > 0 {
		t.Errorf("Error validating property")
	}

	p = model.Property{Type: "FRUTA", Orientation: "FRENTE", State: 3}
	errs = validateProperty(&p)
	log.Printf("Errores: %v", len(errs))
	if len(errs) != 2 {
		t.Errorf("Error validating property")
	}

	p = model.Property{Type: "FRUTA", Orientation: "FRUTA", State: 3, Address: &m.Address{}}
	errs = validateProperty(&p)
	log.Printf("Errores: %v", len(errs))
	if len(errs) != 2 {
		t.Errorf("Error validating property")
	}

	p = model.Property{Type: "FRUTA", Orientation: "FRUTA", State: 23, Address: &m.Address{}}
	errs = validateProperty(&p)
	log.Printf("Errores: %v", len(errs))
	if len(errs) != 3 {
		t.Errorf("Error validating property")
	}
}
