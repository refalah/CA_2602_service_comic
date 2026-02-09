package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/refalah/CA_2602_service_comic/pkg/models"
	"github.com/refalah/CA_2602_service_comic/pkg/utils"
)

func CreateCreatives(w http.ResponseWriter, r *http.Request) {
	newCreatives := &models.CreativeTeam{}
	utils.ParseBody(r, newCreatives)
	b := newCreatives.CreateNewCreativeTeam()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCreatives(w http.ResponseWriter, r *http.Request) {
	newCreativess := models.GetAllCreativeTeams()
	res, _ := json.Marshal(newCreativess)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCreativeTeamByComicId(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    comicId := vars["comicId"]
    ID, err := strconv.ParseInt(comicId, 0, 0)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid comic ID"))
        return
    }
    
    creativeTeam, dbErr := models.GetCreativeTeamByComicId(ID)
    if dbErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Error fetching creative team"))
        return
    }
    
    res, _ := json.Marshal(creativeTeam)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}