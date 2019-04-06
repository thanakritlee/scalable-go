package router

import (
	"github.com/gorilla/mux"
	"github.com/thanakritlee/scalable-go/entryserver/controllers"
)

// GetRouter return a router with registered routes.
func GetRouter() *mux.Router {
	router := mux.NewRouter()

	// APIs route
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")

	return router

}
