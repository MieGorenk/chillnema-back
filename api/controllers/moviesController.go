package controllers

import (
    "encoding/json"
    "io/ioutil"
	"net/http"
	
	"github.com/MieGorenk/data-api/api/models"
	"github.com/MieGorenk/data-api/api/responses"
)

// CreateMovie controller for creating movie in db
func (a *App) CreateMovie(w http.ResponseWriter, r *http.Request) {
    var response = map[string]interface{}{"status": "success", "message": "Registered successfully"}
	movie := &models.Movie{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &movie)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	movieCreated, err := movie.SaveMovie(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	response["movie"] = movieCreated
	responses.JSON(w, http.StatusCreated, response)
	return
}