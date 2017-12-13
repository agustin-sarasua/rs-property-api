package main

import (
	"fmt"
	"log"
	"net/http"

	m "github.com/agustin-sarasua/rs-model"

	"github.com/agustin-sarasua/rs-property-api/app"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", app.ConnectionString)
	app.Db = db
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	var property m.Property
	var address m.Address
	var propertyState m.PropertyState
	db.DropTableIfExists(&property, &address, &propertyState)
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&property, &address, &propertyState)
	//db.Model(&property).AddForeignKey("address_id", "addresses(id)", "CASCADE", "CASCADE")
	//db.Model(&propertyState).AddIndex("idx_property", "property_id")
	db.Model(&property).Related(&address)

	router := mux.NewRouter()
	router.HandleFunc("/property", app.CreatePropertyEndpoint).Methods("POST")
	router.HandleFunc("/property", app.ListPropertiesEndpoint).Methods("GET")
	router.HandleFunc("/property/{id:[0-9]+}", app.GetPropertyEndpoint).Methods("GET")
	router.HandleFunc("/property/{id:[0-9]+}", app.UpdatePropertyEndpoint).Methods("PUT")
	router.HandleFunc("/property/{id:[0-9]+}/state", app.SavePropertyStateEndpoint).Methods("PUT")

	fmt.Println("Hello Property API")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
