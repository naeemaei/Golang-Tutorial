package handlers

import (
	"encoding/json"
	"fmt"
	"http-request-examples/models"
	"net/http"
)

type UsersHandler struct {
}
 

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
		w.WriteHeader(401)
		fmt.Fprintf(w, "Invalid API Key")
		return
	}

	var user *models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Decode Error: %v", err)
		return
	} 

	w.WriteHeader(200)
	fmt.Fprintf(w, "User created")

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	// get form db
	fmt.Fprintf(w, "get user by id: %s", id)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// get list form db
	fmt.Fprintf(w, "get users")
}
