package app_test

import (
	"log"
	"os"
	"testing"

	"github.com/agustin-sarasua/rs-model"
	"github.com/agustin-sarasua/rs-property-api/app"

	model "github.com/agustin-sarasua/rs-model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var address = m.Address{Street: "Av. Sarmiento", Number: "2542", ApartmentNumber: "1202",
	Neighborhood: "PUNTA_CARRETAS", City: "MVD", Country: "UY", PostalCode: "11300",
	Location: m.Location{Latitude: 12.312312, Longitude: -123.123123}}
var aptoOK = model.Property{Type: "APARTAMENTO", Orientation: "FRENTE", State: 3,
	Address: &address, ApartmentsPerFloor: 3, Bedrooms: 2, BedroomsSizes: "54m2,32m2,50m2",
	Kitchens: 1, KitchenSizes: "30m2", Bathrooms: 3, BuildingName: "Edificio Esperalda", Amenities: "GYM,BBQ",
	ConstructionYear: 2009, CourtyardSize: 78, Elevators: 3, BalconySize: "12m2", TerraceSize: "40m2",
	Description: "Hermoso apto", Expenses: 180000, Floors: 1, GarageSize: 25, LivingroomSize: 45,
	Padron: "AA1234"}

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
	app.Db = db

	os.Exit(m.Run())
}

func TestCreateProperty_ERROR(t *testing.T) {
	log.Println("Running TestCreateProperty_OK")
	var aptoERROR model.Property
	aptoERROR = aptoOK
	aptoERROR.Address = nil

	_, errs := app.CreateProperty(&aptoERROR)
	if len(errs) != 1 {
		t.Errorf("Address was null and it passes")
	}

}

func TestCreateProperty_OK(t *testing.T) {
	log.Println("Running TestCreateProperty_Errors")

	id, errs := app.CreateProperty(&aptoOK)
	if len(errs) > 0 {
		t.Errorf("Property was not saved")
	}

	log.Printf("Loading Property recently saved ID %v\n", aptoOK.ID)
	if err := app.Db.Find(&model.Property{}, id).Error; err != nil {
		t.Errorf("Property was not saved, err= %v", err)
	}
}
