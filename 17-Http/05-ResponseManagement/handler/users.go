package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"http-request-examples/api"
	"http-request-examples/models"
	"net/http"
	"strconv"
)

type UsersHandler struct {
}

var Users = map[string]models.User{}

func (u *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {

	case r.Method == http.MethodGet && len(r.URL.Query().Get("id")) > 0:
		GetUser(w, r)
		return
	case r.Method == http.MethodGet && len(r.URL.Query().Get("id")) == 0:
		GetUsers(w, r)
		return
	case r.Method == http.MethodPost:
		CreateUser(w, r)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-Api-Key")

	for k, v := range r.Header {
		fmt.Println(k, v)
	}

	if apiKey != "9289652658659856" {
		api.SetResult(http.StatusUnauthorized, nil, nil, w)
		return
	}

	var user *models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.SetResult(http.StatusInternalServerError, nil, err, w)
		return
	}

	// add to db
	if _, exists := Users[strconv.Itoa(user.Id)]; exists {
		api.SetResult(http.StatusInternalServerError, nil, errors.New("User is exists"), w)
		return
	}
	Users[strconv.Itoa(user.Id)] = *user

	api.SetResult(http.StatusOK, *user, nil, w)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// get form db
	user, exists := Users[id]

	if exists {
		api.SetResult(http.StatusOK, user, nil, w)
	} else {
		api.SetResult(http.StatusNotFound, nil, nil, w)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	api.SetResult(http.StatusOK, Users, nil, w)
}
