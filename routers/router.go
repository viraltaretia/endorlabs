package routers

import (
	"myapp/handlers"

	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/objects/{id}", handlers.GetObjectByIDHandler).Methods("GET")
	router.HandleFunc("/objects", handlers.StoreObjectHandler).Methods("POST")
	router.HandleFunc("/objects/name/{name}", handlers.GetObjectByNameHandler).Methods("GET")
	router.HandleFunc("/objects/kind/{kind}", handlers.ListObjectsHandler).Methods("GET")
	router.HandleFunc("/objects/{id}", handlers.DeleteObjectHandler).Methods("DELETE")
}
