package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agustin-sarasua/pimbay/app/route"
	"github.com/agustin-sarasua/rs-property-api/app"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/property", app.CreatePropertyEndpoint).Methods("POST")
	router.HandleFunc("/property", app.ListPropertiesEndpoint).Methods("GET")
	router.HandleFunc("/property/{id:[0-9]+}", use(app.GetPropertyEndpoint, route.ValidateToken)).Methods("GET")
	router.HandleFunc("/property/{id:[0-9]+}", app.UpdatePropertyEndpoint).Methods("PUT")
	router.HandleFunc("/property/{id:[0-9]+}/state", app.SavePropertyStateEndpoint).Methods("PUT")

	fmt.Println("Hello there")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
