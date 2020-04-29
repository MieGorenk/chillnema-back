package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/MieGorenk/data-api/api/models"
	"github.com/MieGorenk/data-api/api/responses"
	"github.com/gorilla/mux"
)

// CreateMovie controller for creating movie in db
func (a *App) CreateMovie(w http.ResponseWriter, r *http.Request) {
    var response = map[string]interface{}{"status": "success", "message": "Movie uploaded successfully"}
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

//GetMovies to get all movies in the databaase
func (a *App) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := models.GetAllMovies(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, movies)
	return
}

//GetMovieByID to get specific movie by ID
func (a *App) GetMovieByID(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	movie, err := models.GetMovieByID(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, movie)
	return

}