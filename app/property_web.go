package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	c "github.com/agustin-sarasua/rs-common"
	m "github.com/agustin-sarasua/rs-model"
	"github.com/gorilla/mux"
)

func CreatePropertyEndpoint(w http.ResponseWriter, req *http.Request) {
	var msg m.Property
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		c.ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}
	msg.CreatedAt = time.Now()
	if id, errs := CreateProperty(&msg); len(errs) > 0 {
		log.Printf("Error creating property")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(m.BuildErrorResponse(errs))
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "{id: %q}", id)
	}
	w.Header().Set("Content-Type", "application/json")
}

func UpdatePropertyEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.ParseUint(mux.Vars(req)["id"], 10, 64)
	var msg m.Property
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		c.ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}
	msg.CreatedAt = time.Now()
	msg.ID = id
	UpdateProperty(&msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func SavePropertyStateEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.ParseUint(mux.Vars(req)["id"], 10, 64)
	var msg m.PropertyState
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		c.ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}
	msg.CreatedAt = time.Now()
	msg.PropertyID = id
	SavePropertyState(&msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func GetPropertyEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.ParseUint(mux.Vars(req)["id"], 10, 64)

	p := LoadProperty(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func ListPropertiesEndpoint(w http.ResponseWriter, req *http.Request) {
	offset, _ := strconv.Atoi(mux.Vars(req)["offset"])
	if ps, err := ListProperties(offset, 50); err != nil {
		log.Printf("Error loading properties")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(m.BuildErrorResponse([]error{err}))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(SearchResutlDTO{items: ps})
	}
}
