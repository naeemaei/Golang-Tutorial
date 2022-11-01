package management

import (
	"http-mini-project/handlers"
	"net/http"
	"time"
)

func Run() {
	mux := http.NewServeMux()
	mux.Handle("/users/", &handlers.MovieHandler{})
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
