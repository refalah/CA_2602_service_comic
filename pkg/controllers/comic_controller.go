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

func CreateComics(w http.ResponseWriter, r *http.Request)  {
	newComic := &models.Comic{}
	utils.ParseBody(r, newComic)
	b := newComic.CreateNewComic()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetComics(w http.ResponseWriter, r *http.Request) {
	newComics := models.GetAllComics()
	res, _ := json.Marshal(newComics)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetComicById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    comicId := vars["comicId"]  // Changed from creatorId
    ID, err := strconv.ParseInt(comicId, 0, 0)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid comic ID"))
        return
    }
    
    comic, dbErr := models.GetComicById(ID)
    if dbErr != nil {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("Comic not found"))
        return
    }
    
    res, _ := json.Marshal(comic)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func UpdateComic(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	creatorId := vars["creatorId"]
	ID, err := strconv.ParseInt(creatorId, 0, 0)
	if err != nil {
		fmt.Println("Error Parsing")
	}
	updatedComic := &models.Comic{}
	utils.ParseBody(r, updatedComic)
	b := updatedComic.UpdateComicById(ID)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteComicById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	creatorId := vars["creatorId"]
	creator := &models.Comic{}
	ID, err := strconv.ParseInt(creatorId, 0, 0);
	if err != nil {
		fmt.Println("Parsing Error")
	}
	newComics := creator.DeleteComicById(ID)
	res, _ := json.Marshal(newComics)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}