package routes

import (
	"github.com/gorilla/mux"
	"github.com/refalah/CA_2602_service_comic/pkg/controllers"
)

var RegisterComicRoutes = func(router *mux.Router) {
	router.HandleFunc("/comic/", controllers.GetComics).Methods("GET")
	router.HandleFunc("/comic/{comicId}", controllers.GetComicById).Methods("GET")
	router.HandleFunc("/comic/", controllers.CreateComics).Methods("POST")
	router.HandleFunc("/comic/{comicId}", controllers.UpdateComic).Methods("PUT")
	router.HandleFunc("/comic/{comicId}", controllers.DeleteComicById).Methods("DELETE")
}