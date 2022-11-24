package routes

import (
	"github.com/gorilla/mux"

	controllers "go-table/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/get/", controllers.GetTableProperty).Methods("GET")
	router.HandleFunc("/delete/", controllers.DeleteTableProperty).Methods("PUT")
	router.HandleFunc("/update", controllers.UpdateTableProperty).Methods("PUT")
	router.HandleFunc("/single", controllers.GetSingleProperty).Methods("GET")
}
