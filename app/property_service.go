package app

import (
	"database/sql"
	"log"
	"strings"
	"time"

	m "github.com/agustin-sarasua/rs-model"
	"github.com/agustin-sarasua/rs-property-api/cons"
	_ "github.com/go-sql-driver/mysql"
)

func CreateProperty(p *m.Property) int64 {
	log.Printf("Creating new Property: %+v\n", p)

	db, err := sql.Open("mysql", cons.ConnectionString)
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
	stmtIns, err := db.Prepare(cons.PropertyInsertQuery) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()

	r, err := stmtIns.Exec(p.Description,
		p.Type,
		p.Orientation,
		p.CourtyardSize,
		p.Bedrooms,
		p.LivingRoomSize,
		p.KitchenSize,
		strings.Join(p.BedroomsSizes, ","),
		p.Bathrooms,
		p.Showers,
		p.Size,
		p.ConstructionYear,
		p.Padron,
		p.BuildingName,
		p.ApartmentsPerFloor,
		p.Floors,
		p.TerraceSize,
		p.BalconySize,
		p.Expenses,
		strings.Join(p.Amenities, ","),
		time.Now(),
		p.Address.Street,
		p.Address.Number,
		p.Address.ApartmentNumber,
		p.Address.Neighborhood,
		p.Address.City,
		p.Address.Country,
		p.Address.PostalCode,
		p.Address.Location.Latitude,
		p.Address.Location.Longitude,
		p.Elevators) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	id, err := r.LastInsertId()
	log.Printf("LAST ID %v", id)
	return id
}
