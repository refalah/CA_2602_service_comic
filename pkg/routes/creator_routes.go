package routes

import (
	"github.com/gorilla/mux"
	"github.com/refalah/go-comics/pkg/controllers"
)

var RegisterCreatorRoutes = func(router *mux.Router) {
	router.HandleFunc("/creator/", controllers.GetCreators).Methods("GET")
	router.HandleFunc("/creator/{creatorId}", controllers.GetCreatorById).Methods("GET")
	router.HandleFunc("/creator/", controllers.CreateCreators).Methods("POST")
	router.HandleFunc("/creator/{creatorId}", controllers.UpdateCreator).Methods("PUT")
	router.HandleFunc("/creator/{creatorId}", controllers.DeleteCreatorById).Methods("DELETE")
}