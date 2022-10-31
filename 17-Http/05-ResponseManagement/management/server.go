package management

import (
	handlers "http-request-examples/handler"
	"net/http"
	"time"
)

func Run() {
	mux := http.NewServeMux()
	mux.Handle("/users/", &handlers.UsersHandler{})
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
