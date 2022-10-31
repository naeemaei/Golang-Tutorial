package examples

import (
	"fmt"
	"net/http"
	"time"
)

type TestHandler struct {
}

func CreateServerWithMux() {
	mux := http.NewServeMux()

	mux.Handle("/google", http.RedirectHandler("http://www.google.com", 307))
	mux.Handle("/yahoo", http.RedirectHandler("http://www.yahoo.com", 307))
	mux.HandleFunc("/get-users", GetUsers)
	mux.HandleFunc("/get-status", GetStatus)

	server1 := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      mux,
	}

	err := server1.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func CreateServerWithCustomHandler() {

	server1 := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      TestHandler{},
	}

	err := server1.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func (h TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	fmt.Fprintf(w, "Hello, world!")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	fmt.Fprintf(w, "{'id': '%d', 'name': '%s', 'method': '%s'}", 10, "Peyman", r.Method)
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am OK")
}
