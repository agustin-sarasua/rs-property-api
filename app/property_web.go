package app

import (
	"encoding/json"
	"net/http"

	c "github.com/agustin-sarasua/rs-common"
	m "github.com/agustin-sarasua/rs-model"
)

func CreatePropertyEndpoint(w http.ResponseWriter, req *http.Request) {

	var msg m.Property
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		c.ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}
	msg.CreatedDate = c.DateTime{}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func UpdatePropertyEndpoint(w http.ResponseWriter, req *http.Request) {
}

func SavePropertyStateEndpoint(w http.ResponseWriter, req *http.Request) {
}

func GetPropertyEndpoint(w http.ResponseWriter, req *http.Request) {

	var msg m.Property
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		c.ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}
	msg.CreatedDate = c.DateTime{}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func ListPropertiesEndpoint(w http.ResponseWriter, req *http.Request) {

	var msg m.Property
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		c.ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}
	msg.CreatedDate = c.DateTime{}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}
