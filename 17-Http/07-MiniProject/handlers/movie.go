package handlers

import (
	"encoding/json"
	db "http-mini-project/database"
	"http-mini-project/models"
	"http-mini-project/utils"
	"net/http"
)

type MovieHandler struct {
}

var Users = map[string]models.Movie{}

func (u *MovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {

	case r.Method == http.MethodGet && len(r.URL.Query().Get("id")) > 0:
		GetMovie(w, r)
		return
	case r.Method == http.MethodGet && len(r.URL.Query().Get("id")) == 0:
		GetMovies(w, r)
		return
	case r.Method == http.MethodPost:
		AddMovie(w, r)
		return
	case r.Method == http.MethodDelete:
		DeleteMovie(w, r)
		return
	}
}

func GetMovies(w http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		utils.SetResult(w, http.StatusMethodNotAllowed, nil, utils.ApiError{Id: "httpMethod", Message: "Not allowed"})
		return
	}

	var movies []models.Movie

	for _, movie := range db.MovieDb {
		movies = append(movies, movie)
	}

	// parse the movie data into json format
	// moviesJSON, err := json.Marshal(&movies)
	// if err != nil {
	// 	utils.SetResult(w, http.StatusInternalServerError, nil, utils.ApiError{Id: "jsonMarshal", Message: "Error parsing the movie data" + err.Error()})
	// 	return
	// }

	utils.SetResult(w, http.StatusOK, movies, nil)
}

func GetMovie(w http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		utils.SetResult(w, http.StatusMethodNotAllowed, nil, utils.ApiError{Id: "httpMethod", Message: "Not allowed"})
		return
	}

	if _, ok := req.URL.Query()["id"]; !ok {
		utils.SetResult(w, http.StatusBadRequest, nil, utils.ApiError{Id: "QueryStringMissed", Message: "This method requires the movie id"})
		return
	}

	id := req.URL.Query()["id"][0]

	movie, ok := db.MovieDb[id]
	if !ok {
		utils.SetResult(w, http.StatusNotFound, nil, nil)
		return
	}

	// parse the movie data into json format
	// movieJSON, err := json.Marshal(&movie)
	// if err != nil {
	// 	utils.SetResult(w, http.StatusInternalServerError, nil, utils.ApiError{Id: "jsonMarshal", Message: "Error parsing the movie data" + err.Error()})
	// 	return
	// }

	utils.SetResult(w, http.StatusOK, movie, nil)
}

func AddMovie(w http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		utils.SetResult(w, http.StatusMethodNotAllowed, nil, utils.ApiError{Id: "httpMethod", Message: "Not allowed"})
		return
	}

	var movie models.Movie

	payload := req.Body

	defer req.Body.Close()
	// parse the movie data into json format
	err := json.NewDecoder(payload).Decode(&movie)
	if err != nil {
		utils.SetResult(w, http.StatusInternalServerError, nil, utils.ApiError{Id: "jsonMarshal", Message: "Error parsing the movie data" + err.Error()})
		return
	}

	db.MovieDb[movie.ID] = movie

	utils.SetResult(w, http.StatusCreated, nil, nil)
}

func DeleteMovie(w http.ResponseWriter, req *http.Request) {

	if req.Method != "DELETE" {
		utils.SetResult(w, http.StatusMethodNotAllowed, nil, utils.ApiError{Id: "httpMethod", Message: "Not allowed"})
		return
	}

	if _, ok := req.URL.Query()["id"]; !ok {
		utils.SetResult(w, http.StatusBadRequest, nil, utils.ApiError{Id: "QueryStringMissed", Message: "This method requires the movie id"})
		return
	}

	id := req.URL.Query()["id"][0]
	movie, ok := db.MovieDb[id]
	if !ok {
		utils.SetResult(w, http.StatusNotFound, nil, nil)
		return
	}
	// parse the movie data into json format
	movieJSON, err := json.Marshal(&movie)
	if err != nil {
		utils.SetResult(w, http.StatusInternalServerError, nil, utils.ApiError{Id: "jsonMarshal", Message: "Error parsing the movie data" + err.Error()})
		return
	}

	utils.SetResult(w, http.StatusOK, nil, movieJSON)
}
