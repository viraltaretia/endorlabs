package handlers

import (
	"bytes"
	"encoding/json"
	"io"

	"errors"
	"myapp/internal/db"
	"myapp/internal/object"
	"net/http"

	e "myapp/pkg/errors"

	"github.com/gorilla/mux"
)

// GetObjectByIDHandler is the handler for retrieving an object by ID.
func GetObjectByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	dbObj := r.Context().Value("db").(*db.DB)

	obj, err := dbObj.GetObjectByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, e.ErrObjectNotFound) {
			http.NotFound(w, r)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func StoreObjectHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Kind string `json:"kind"`
	}

	data, _ := io.ReadAll(r.Body)

	// For multiple time decoding request body
	r1 := io.NopCloser(bytes.NewBuffer(data))

	err := json.NewDecoder(r1).Decode(&body)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	r2 := io.NopCloser(bytes.NewBuffer(data))

	var obj object.Object
	switch body.Kind {
	case "Person":
		var person object.Person
		err = json.NewDecoder(r2).Decode(&person)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		obj = &person
	case "Animal":
		var animal object.Animal
		err = json.NewDecoder(r2).Decode(&animal)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		obj = &animal
	default:
		http.Error(w, "Unsupported kind", http.StatusBadRequest)
		return
	}

	dbObj := r.Context().Value("db").(*db.DB)

	err = dbObj.Store(r.Context(), obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetObjectByNameHandler is the handler for retrieving an object by name.
func GetObjectByNameHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	dbObj := r.Context().Value("db").(*db.DB)

	obj, err := dbObj.GetObjectByName(r.Context(), name)
	if err != nil {
		if errors.Is(err, e.ErrObjectNotFound) {
			http.NotFound(w, r)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

// ListObjectsHandler is the handler for listing objects of a specific kind.
func ListObjectsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kind := params["kind"]

	dbObj := r.Context().Value("db").(*db.DB)

	objects, err := dbObj.ListObjects(r.Context(), kind)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(objects)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

// DeleteObjectHandler is the handler for deleting an object by ID.
func DeleteObjectHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	dbObj := r.Context().Value("db").(*db.DB)
	err := dbObj.DeleteObject(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
