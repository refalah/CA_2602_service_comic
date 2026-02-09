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

func CreateCreators(w http.ResponseWriter, r *http.Request)  {
	newCreator := &models.Creator{}
	utils.ParseBody(r, newCreator)
	b := newCreator.CreateNewCreator()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCreators(w http.ResponseWriter, r *http.Request) {
	newCreators := models.GetAllCreators()
	res, _ := json.Marshal(newCreators)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCreatorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	creatorId := vars["creatorId"]
	ID, err := strconv.ParseInt(creatorId, 0, 0);
	if err != nil {
		fmt.Println("Parsing Error")
	}
	newCreators := models.GetCreatorById(ID)
	res, _ := json.Marshal(newCreators)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCreator(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	creatorId := vars["creatorId"]
	ID, err := strconv.ParseInt(creatorId, 0, 0)
	if err != nil {
		fmt.Println("Error Parsing")
	}
	updatedCreator := &models.Creator{}
	utils.ParseBody(r, updatedCreator)
	b := updatedCreator.UpdateCreatorById(ID)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCreatorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	creatorId := vars["creatorId"]
	creator := &models.Creator{}
	ID, err := strconv.ParseInt(creatorId, 0, 0);
	if err != nil {
		fmt.Println("Parsing Error")
	}
	newCreators := creator.DeleteCreatorById(ID)
	res, _ := json.Marshal(newCreators)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}