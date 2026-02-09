package routes

import (
	"github.com/gorilla/mux"
	"github.com/refalah/go-comics/pkg/controllers"
)

var RegisterCreativeRoutes = func(router *mux.Router) {
	router.HandleFunc("/creatives/", controllers.GetCreatives).Methods("GET")
	router.HandleFunc("/creatives/{comicId}", controllers.GetCreativeTeamByComicId).Methods("GET")
	router.HandleFunc("/creatives/", controllers.CreateCreatives).Methods("POST")
}