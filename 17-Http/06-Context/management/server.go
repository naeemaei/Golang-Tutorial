package management

import (
	"fmt"
	"net/http"
	"time"
)

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test/", TestHandler)

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

func TestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("processing request")
		fmt.Fprintf(w, "Response")
	case <-ctx.Done():
		fmt.Println("cancel request")
		return
	}
}
