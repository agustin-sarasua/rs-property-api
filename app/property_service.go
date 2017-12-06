package app

import (
	"database/sql"
	"log"
	"strings"

	m "github.com/agustin-sarasua/rs-model"
	_ "github.com/go-sql-driver/mysql"
)

func CreateProperty(p *m.Property) {
	db, err := sql.Open("mysql", "root:root@/gmailml")
	if err != nil {
		log.Fatalf("Unable to open mysql connection: %v", err)
	}
	defer db.Close()
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// Prepare statement for inserting data

	stmtIns, err := db.Prepare(PropertyInsertQuery) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	a := strings.Join(p.Amenities, ",")

	_, err = stmtIns.Exec(p.Description, p.Type, p.Orientation, p.CourtyardSize, p.Bedrooms, p.LivingRoomSize, p.KitchenSize, p.BedroomsSizes, p.Showers, p.Size, p.ConstructionYear, p.Padron, p.BuildingName, p.ApartmentsPerFloor, p.Floors, p.TerraceSize, p.BalconySize, p.Expenses, a, p.CreatedDate, p.Address.Street, p.Address.Number, p.Address.ApartmentNumber, p.Address.Neighborhood, p.Address.City, p.Address.Country, p.Address.PostalCode, p.Address.Location.Latitude, p.Address.Location.Longitude) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

}
