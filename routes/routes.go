package routes

import (
	"github.com/gorilla/mux"
	"onursahin.dev/vuegolith/endpoints"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/log", endpoints.HandleLogGet).Methods("GET")
	router.HandleFunc("/api/log", endpoints.HandleLogPost).Methods("POST")

	router.HandleFunc("/api/uploads", endpoints.HandleUpload).Methods("POST")
	router.HandleFunc("/api/uploads", endpoints.HandleListUploads).Methods("GET")
	router.HandleFunc("/api/uploads", endpoints.HandleDeleteUpload).Methods("DELETE")

	return router
}
