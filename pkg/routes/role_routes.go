package routes

import (
	"github.com/gorilla/mux"
	"github.com/refalah/CA_2602_service_comic/pkg/controllers"
)

var RegisterRoleRoutes = func(router *mux.Router) {
	router.HandleFunc("/role/", controllers.GetRoles).Methods("GET")
	router.HandleFunc("/role/{roleId}", controllers.GetRole).Methods("GET")
	router.HandleFunc("/role/", controllers.CreateRole).Methods("POST")
	router.HandleFunc("/role/{roleId}", controllers.UpdateRole).Methods("PUT")
	router.HandleFunc("/role/{roleId}", controllers.DeleteRole).Methods("DELETE")
}