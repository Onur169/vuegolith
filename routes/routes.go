package routes

import (
	"github.com/gorilla/mux"
	"onursahin.dev/vuegolith/endpoints"
)

func SetupRoutes(apiPrefix string) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(apiPrefix + "/log", endpoints.HandleLogGet).Methods("GET")
	router.HandleFunc(apiPrefix + "/log", endpoints.HandleLogPost).Methods("POST")

	router.HandleFunc(apiPrefix + "/uploads", endpoints.HandleUpload).Methods("POST")
	router.HandleFunc(apiPrefix + "/uploads", endpoints.HandleListUploads).Methods("GET")
	router.HandleFunc(apiPrefix + "/uploads", endpoints.HandleDeleteUpload).Methods("DELETE")

	return router
}
