package examples

import (
	"log"
	"net/http"
	"time"
)

func CreateServer() {
	go func() {
		server1 := &http.Server{
			Addr:         ":8080",
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		}

		err := server1.ListenAndServe()

		if err != nil {
			panic(err)
		}
	}()

	log.Fatal(http.ListenAndServe(":8090", nil))
}
