package routes

import (
	"github.com/gorilla/mux"
	"onursahin.dev/vuegolith/endpoints"
)

const ApiPrefix = "/api"

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(ApiPrefix + "/log", endpoints.HandleLogGet).Methods("GET")
	router.HandleFunc(ApiPrefix + "/log", endpoints.HandleLogPost).Methods("POST")

	router.HandleFunc(ApiPrefix + "/uploads", endpoints.HandleUpload).Methods("POST")
	router.HandleFunc(ApiPrefix + "/uploads", endpoints.HandleListUploads).Methods("GET")
	router.HandleFunc(ApiPrefix + "/uploads", endpoints.HandleDeleteUpload).Methods("DELETE")

	return router
}
