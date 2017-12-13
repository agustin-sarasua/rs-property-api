package app_test

import (
	"log"
	"os"
	"testing"

	model "github.com/agustin-sarasua/rs-model"
	"github.com/agustin-sarasua/rs-property-api/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestMain(m *testing.M) {
	log.Printf("Running test main")
	// call flag.Parse() here if TestMain uses flags
	db, _ := gorm.Open("sqlite3", "/tmp/gorm.db")
	defer db.Close()

	// Migrate the schema
	var property model.Property
	var address model.Address
	var propertyState model.PropertyState
	db.DropTableIfExists(&property, &address, &propertyState)
	db.AutoMigrate(&property, &address, &propertyState)
	//db.Model(&property).AddForeignKey("address_id", "addresses(id)", "CASCADE", "CASCADE")
	//db.Model(&propertyState).AddIndex("idx_property", "property_id")
	log.Printf("Running specific test")
	db.Model(&property).Related(&address)
	app.Db = db

	os.Exit(m.Run())
}

func TestCreateProperty(t *testing.T) {
	log.Printf("Running TestCreateProperty")
	var p model.Property
	p = model.Property{Type: "RENTAL"}
	id := app.CreateProperty(&p)

	log.Printf("Loading Property recently saved ID %v", p.ID)
	var pSaved model.Property
	app.Db.Find(&pSaved, id)

	if &pSaved == nil {
		t.Errorf("Property was not saved")
	}

}
