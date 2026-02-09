package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/refalah/go-comics/pkg/models"
	"github.com/refalah/go-comics/pkg/utils"
)


func CreateRole(w http.ResponseWriter, r *http.Request) {
	newRole := &models.Role{}
	utils.ParseBody(r, newRole)
	b := newRole.CreateNewRole()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRoles(w http.ResponseWriter, r *http.Request) {
	newRoles := models.GetAllRoles()
	res, _ := json.Marshal(newRoles)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRole(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	roleId := vars["roleId"]
	ID, err := strconv.ParseInt(roleId,0,0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	role := models.GetRoleById(ID)
	res, _ := json.Marshal(role)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateRole(w http.ResponseWriter, r *http.Request)  {
	updatedRole := &models.Role{}
	vars := mux.Vars(r)
	roleId := vars["roleId"]
	ID, err := strconv.ParseInt(roleId,0,0)
	if err != nil {
		fmt.Println("Parsing Error")
	}	

	utils.ParseBody(r, updatedRole)
	
	role := updatedRole.UpdateRole(ID)
	res, _ := json.Marshal(role)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteRole(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	roleId := vars["roleId"]
	ID, err := strconv.ParseInt(roleId,0,0)
	if err != nil {
		fmt.Println("Parsing error")
	}
	role := models.DeleteRoleById(ID)
	res, _ := json.Marshal(role)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
